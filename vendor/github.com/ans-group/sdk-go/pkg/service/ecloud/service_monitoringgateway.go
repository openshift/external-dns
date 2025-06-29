package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetMonitoringGateways retrieves a list of monitoring gateways
func (s *Service) GetMonitoringGateways(parameters connection.APIRequestParameters) ([]MonitoringGateway, error) {
	return connection.InvokeRequestAll(s.GetMonitoringGatewaysPaginated, parameters)
}

// GetMonitoringGatewaysPaginated retrieves a paginated list of monitoring gateways
func (s *Service) GetMonitoringGatewaysPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[MonitoringGateway], error) {
	body, err := s.GetMonitoringGatewaysPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetMonitoringGatewaysPaginated), err
}

func (s *Service) GetMonitoringGatewaysPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]MonitoringGateway], error) {
	body := &connection.APIResponseBodyData[[]MonitoringGateway]{}

	response, err := s.connection.Get("/ecloud/v2/monitoring-gateways", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetMonitoringGateway retrieves a single monitoring gateway by ID
func (s *Service) GetMonitoringGateway(gatewayID string) (MonitoringGateway, error) {
	body, err := s.getMonitoringGatewayResponseBody(gatewayID)

	return body.Data, err
}

func (s *Service) getMonitoringGatewayResponseBody(gatewayID string) (*connection.APIResponseBodyData[MonitoringGateway], error) {
	if gatewayID == "" {
		return &connection.APIResponseBodyData[MonitoringGateway]{}, fmt.Errorf("invalid monitoring gateway id")
	}

	return connection.Get[MonitoringGateway](s.connection, fmt.Sprintf("/ecloud/v2/monitoring-gateways/%s", gatewayID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&MonitoringGatewayNotFoundError{ID: gatewayID}))
}
