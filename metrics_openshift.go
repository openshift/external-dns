/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	authenticationv1 "k8s.io/api/authentication/v1"
	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// serveSecureMetrics starts an HTTPS metrics server with Kubernetes TokenReview/SAR authentication.
func serveSecureMetrics(cfg *externaldns.Config) {
	restCfg, err := clientcmd.BuildConfigFromFlags(cfg.APIServerURL, cfg.KubeConfig)
	if err != nil {
		restCfg, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("metrics auth: failed to build rest config: %v", err)
		}
	}
	httpClient, err := rest.HTTPClientFor(restCfg)
	if err != nil {
		log.Fatalf("metrics auth: failed to build http client: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfigAndClient(restCfg, httpClient)
	if err != nil {
		log.Fatalf("metrics auth: failed to build kube client: %v", err)
	}

	server := &http.Server{Addr: cfg.MetricsAddress, Handler: newMetricsMux(kubeClient, cfg.MetricsTLSCertDir)}
	log.Fatal(server.ListenAndServeTLS(
		filepath.Join(cfg.MetricsTLSCertDir, "tls.crt"),
		filepath.Join(cfg.MetricsTLSCertDir, "tls.key"),
	))
}

// newMetricsMux builds the ServeMux for the secure metrics server.
// /metrics is wrapped with withMetricsAuth; /healthz is always unauthenticated.
func newMetricsMux(kubeClient kubernetes.Interface, tlsCertDir string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.Handle("/metrics", withMetricsAuth(kubeClient, promhttp.Handler()))
	return mux
}

// withMetricsAuth authenticates requests via TokenReview and authorizes via SubjectAccessReview.
func withMetricsAuth(kubeClient kubernetes.Interface, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		tr, err := kubeClient.AuthenticationV1().TokenReviews().Create(
			r.Context(),
			&authenticationv1.TokenReview{
				Spec: authenticationv1.TokenReviewSpec{Token: token},
			},
			metav1.CreateOptions{},
		)
		if err != nil || !tr.Status.Authenticated {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sar, err := kubeClient.AuthorizationV1().SubjectAccessReviews().Create(
			r.Context(),
			&authorizationv1.SubjectAccessReview{
				Spec: authorizationv1.SubjectAccessReviewSpec{
					User:   tr.Status.User.Username,
					Groups: tr.Status.User.Groups,
					NonResourceAttributes: &authorizationv1.NonResourceAttributes{
						Path: "/metrics",
						Verb: "get",
					},
				},
			},
			metav1.CreateOptions{},
		)
		if err != nil || !sar.Status.Allowed {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
