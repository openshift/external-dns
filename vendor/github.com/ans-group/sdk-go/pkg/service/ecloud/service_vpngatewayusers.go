package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNGatewayUsers retrieves a list of VPN gateway users
func (s *Service) GetVPNGatewayUsers(parameters connection.APIRequestParameters) ([]VPNGatewayUser, error) {
	return connection.InvokeRequestAll(s.GetVPNGatewayUsersPaginated, parameters)
}

// GetVPNGatewayUsersPaginated retrieves a paginated list of VPN gateway users
func (s *Service) GetVPNGatewayUsersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGatewayUser], error) {
	body, err := s.getVPNGatewayUsersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNGatewayUsersPaginated), err
}

func (s *Service) getVPNGatewayUsersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPNGatewayUser], error) {
	body := &connection.APIResponseBodyData[[]VPNGatewayUser]{}

	response, err := s.connection.Get("/ecloud/v2/vpn-gateway-users", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNGatewayUser retrieves a single VPN gateway user by ID
func (s *Service) GetVPNGatewayUser(userID string) (VPNGatewayUser, error) {
	body, err := s.getVPNGatewayUserResponseBody(userID)

	return body.Data, err
}

func (s *Service) getVPNGatewayUserResponseBody(userID string) (*connection.APIResponseBodyData[VPNGatewayUser], error) {
	body := &connection.APIResponseBodyData[VPNGatewayUser]{}

	if userID == "" {
		return body, fmt.Errorf("invalid vpn gateway user id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-gateway-users/%s", userID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayUserNotFoundError{ID: userID}
		}

		return nil
	})
}

// CreateVPNGatewayUser creates a new VPN gateway user
func (s *Service) CreateVPNGatewayUser(req CreateVPNGatewayUserRequest) (TaskReference, error) {
	body, err := s.createVPNGatewayUserResponseBody(req)

	return body.Data, err
}

func (s *Service) createVPNGatewayUserResponseBody(req CreateVPNGatewayUserRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	response, err := s.connection.Post("/ecloud/v2/vpn-gateway-users", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVPNGatewayUser patches a VPN gateway user
func (s *Service) PatchVPNGatewayUser(userID string, req PatchVPNGatewayUserRequest) (TaskReference, error) {
	body, err := s.patchVPNGatewayUserResponseBody(userID, req)

	return body.Data, err
}

func (s *Service) patchVPNGatewayUserResponseBody(userID string, req PatchVPNGatewayUserRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if userID == "" {
		return body, fmt.Errorf("invalid user id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vpn-gateway-users/%s", userID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayUserNotFoundError{ID: userID}
		}

		return nil
	})
}

// DeleteVPNGatewayUser deletes a VPN gateway user
func (s *Service) DeleteVPNGatewayUser(userID string) (string, error) {
	body, err := s.deleteVPNGatewayUserResponseBody(userID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVPNGatewayUserResponseBody(userID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

	if userID == "" {
		return body, fmt.Errorf("invalid user id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpn-gateway-users/%s", userID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewayUserNotFoundError{ID: userID}
		}

		return nil
	})
}
