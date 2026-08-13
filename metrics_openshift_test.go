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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	authenticationv1 "k8s.io/api/authentication/v1"
	authorizationv1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestWithMetricsAuth(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title          string
		authHeader     string
		tokenReview    *authenticationv1.TokenReview
		tokenReviewErr error
		sar            *authorizationv1.SubjectAccessReview
		sarErr         error
		wantStatus     int
		wantNextCalled bool
	}{
		{
			title:          "no authorization header yields 401",
			authHeader:     "",
			wantStatus:     http.StatusUnauthorized,
			wantNextCalled: false,
		},
		{
			title:          "non-bearer scheme yields 401",
			authHeader:     "Basic dXNlcjpwYXNz",
			wantStatus:     http.StatusUnauthorized,
			wantNextCalled: false,
		},
		{
			title:          "token review API error yields 401",
			authHeader:     "Bearer some-token",
			tokenReviewErr: fmt.Errorf("api server unavailable"),
			wantStatus:     http.StatusUnauthorized,
			wantNextCalled: false,
		},
		{
			title:      "unauthenticated token review yields 401",
			authHeader: "Bearer bad-token",
			tokenReview: &authenticationv1.TokenReview{
				Status: authenticationv1.TokenReviewStatus{Authenticated: false},
			},
			wantStatus:     http.StatusUnauthorized,
			wantNextCalled: false,
		},
		{
			title:      "subject access review API error yields 403",
			authHeader: "Bearer valid-token",
			tokenReview: &authenticationv1.TokenReview{
				Status: authenticationv1.TokenReviewStatus{
					Authenticated: true,
					User:          authenticationv1.UserInfo{Username: "prometheus"},
				},
			},
			sarErr:         fmt.Errorf("api server unavailable"),
			wantStatus:     http.StatusForbidden,
			wantNextCalled: false,
		},
		{
			title:      "subject access review not allowed yields 403",
			authHeader: "Bearer valid-token",
			tokenReview: &authenticationv1.TokenReview{
				Status: authenticationv1.TokenReviewStatus{
					Authenticated: true,
					User:          authenticationv1.UserInfo{Username: "prometheus"},
				},
			},
			sar: &authorizationv1.SubjectAccessReview{
				Status: authorizationv1.SubjectAccessReviewStatus{Allowed: false},
			},
			wantStatus:     http.StatusForbidden,
			wantNextCalled: false,
		},
		{
			title:      "authenticated and authorized passes through to handler",
			authHeader: "Bearer valid-token",
			tokenReview: &authenticationv1.TokenReview{
				Status: authenticationv1.TokenReviewStatus{
					Authenticated: true,
					User: authenticationv1.UserInfo{
						Username: "system:serviceaccount:monitoring:prometheus-k8s",
						Groups:   []string{"system:serviceaccounts", "system:serviceaccounts:monitoring"},
					},
				},
			},
			sar: &authorizationv1.SubjectAccessReview{
				Status: authorizationv1.SubjectAccessReviewStatus{Allowed: true},
			},
			wantStatus:     http.StatusOK,
			wantNextCalled: true,
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			fakeClient := fake.NewSimpleClientset()

			if tc.tokenReviewErr != nil {
				fakeClient.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
					return true, nil, tc.tokenReviewErr
				})
			} else if tc.tokenReview != nil {
				fakeClient.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
					return true, tc.tokenReview, nil
				})
			}

			if tc.sarErr != nil {
				fakeClient.PrependReactor("create", "subjectaccessreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
					return true, nil, tc.sarErr
				})
			} else if tc.sar != nil {
				fakeClient.PrependReactor("create", "subjectaccessreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
					return true, tc.sar, nil
				})
			}

			nextCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
				w.WriteHeader(http.StatusOK)
			})

			req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
			if tc.authHeader != "" {
				req.Header.Set("Authorization", tc.authHeader)
			}
			rec := httptest.NewRecorder()

			withMetricsAuth(fakeClient, next).ServeHTTP(rec, req)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, tc.wantNextCalled, nextCalled)
		})
	}
}

// TestWithMetricsAuthSARPayload verifies that the user identity and NonResourceAttributes
// from the TokenReview response are forwarded correctly to the SubjectAccessReview.
func TestWithMetricsAuthSARPayload(t *testing.T) {
	t.Parallel()

	expectedUser := "system:serviceaccount:monitoring:prometheus-k8s"
	expectedGroups := []string{"system:serviceaccounts", "system:serviceaccounts:monitoring"}

	fakeClient := fake.NewSimpleClientset()
	fakeClient.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
		return true, &authenticationv1.TokenReview{
			Status: authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: expectedUser,
					Groups:   expectedGroups,
				},
			},
		}, nil
	})
	fakeClient.PrependReactor("create", "subjectaccessreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
		sar := action.(k8stesting.CreateAction).GetObject().(*authorizationv1.SubjectAccessReview)
		assert.Equal(t, expectedUser, sar.Spec.User)
		assert.Equal(t, expectedGroups, sar.Spec.Groups)
		assert.NotNil(t, sar.Spec.NonResourceAttributes)
		assert.Equal(t, "/metrics", sar.Spec.NonResourceAttributes.Path)
		assert.Equal(t, "get", sar.Spec.NonResourceAttributes.Verb)
		return true, &authorizationv1.SubjectAccessReview{
			Status: authorizationv1.SubjectAccessReviewStatus{Allowed: true},
		}, nil
	})

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	req.Header.Set("Authorization", "Bearer some-token")
	rec := httptest.NewRecorder()

	withMetricsAuth(fakeClient, next).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestNewMetricsMux(t *testing.T) {
	t.Parallel()

	fakeClient := fake.NewSimpleClientset()
	mux := newMetricsMux(fakeClient, "/some/cert/dir")

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}
