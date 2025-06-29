package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNGateways retrieves a list of VPN gateways
func (s *Service) GetVPNGateways(parameters connection.APIRequestParameters) ([]VPNGateway, error) {
	return connection.InvokeRequestAll(s.GetVPNGatewaysPaginated, parameters)
}

// GetVPNGatewaysPaginated retrieves a paginated list of VPN gateways
func (s *Service) GetVPNGatewaysPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGateway], error) {
	body, err := s.getVPNGatewaysPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNGatewaysPaginated), err
}

func (s *Service) getVPNGatewaysPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPNGateway], error) {
	body := &connection.APIResponseBodyData[[]VPNGateway]{}

	response, err := s.connection.Get("/ecloud/v2/vpn-gateways", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNGateway retrieves a single VPN gateway by ID
func (s *Service) GetVPNGateway(gatewayID string) (VPNGateway, error) {
	body, err := s.getVPNGatewayResponseBody(gatewayID)

	return body.Data, err
}

func (s *Service) getVPNGatewayResponseBody(gatewayID string) (*connection.APIResponseBodyData[VPNGateway], error) {
	body := &connection.APIResponseBodyData[VPNGateway]{}

	if gatewayID == "" {
		return body, fmt.Errorf("invalid vpn gateway id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-gateways/%s", gatewayID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayNotFoundError{ID: gatewayID}
		}

		return nil
	})
}

// CreateVPNGateway creates a new VPN gateway
func (s *Service) CreateVPNGateway(req CreateVPNGatewayRequest) (TaskReference, error) {
	body, err := s.createVPNGatewayResponseBody(req)

	return body.Data, err
}

func (s *Service) createVPNGatewayResponseBody(req CreateVPNGatewayRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/vpn-gateways", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVPNGateway patches a VPN gateway
func (s *Service) PatchVPNGateway(gatewayID string, req PatchVPNGatewayRequest) (TaskReference, error) {
	body, err := s.patchVPNGatewayResponseBody(gatewayID, req)

	return body.Data, err
}

func (s *Service) patchVPNGatewayResponseBody(gatewayID string, req PatchVPNGatewayRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if gatewayID == "" {
		return body, fmt.Errorf("invalid gateway id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vpn-gateways/%s", gatewayID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayNotFoundError{ID: gatewayID}
		}

		return nil
	})
}

// DeleteVPNGateway deletes a VPN gateway
func (s *Service) DeleteVPNGateway(gatewayID string) (string, error) {
	body, err := s.deleteVPNGatewayResponseBody(gatewayID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVPNGatewayResponseBody(gatewayID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if gatewayID == "" {
		return body, fmt.Errorf("invalid gateway id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpn-gateways/%s", gatewayID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayNotFoundError{ID: gatewayID}
		}

		return nil
	})
}

// GetVPNGatewayTasks retrieves a list of VPN gateway tasks
func (s *Service) GetVPNGatewayTasks(gatewayID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNGatewayTasksPaginated(gatewayID, p)
	}, parameters)
}

// GetVPNGatewayTasksPaginated retrieves a paginated list of VPN gateway tasks
func (s *Service) GetVPNGatewayTasksPaginated(gatewayID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getVPNGatewayTasksPaginatedResponseBody(gatewayID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNGatewayTasksPaginated(gatewayID, p)
	}), err
}

func (s *Service) getVPNGatewayTasksPaginatedResponseBody(gatewayID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

	if gatewayID == "" {
		return body, fmt.Errorf("invalid vpn gateway id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-gateways/%s/tasks", gatewayID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayNotFoundError{ID: gatewayID}
		}

		return nil
	})
}
