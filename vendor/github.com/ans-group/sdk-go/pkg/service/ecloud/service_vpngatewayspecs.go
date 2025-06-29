package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNGatewaySpecifications retrieves a list of VPN gateway specifications
func (s *Service) GetVPNGatewaySpecifications(parameters connection.APIRequestParameters) ([]VPNGatewaySpecification, error) {
	return connection.InvokeRequestAll(s.GetVPNGatewaySpecificationsPaginated, parameters)
}

// GetVPNGatewaySpecificationsPaginated retrieves a paginated list of VPN gateway specifications
func (s *Service) GetVPNGatewaySpecificationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNGatewaySpecification], error) {
	body, err := s.getVPNGatewaySpecificationsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNGatewaySpecificationsPaginated), err
}

func (s *Service) getVPNGatewaySpecificationsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPNGatewaySpecification], error) {
	body := &connection.APIResponseBodyData[[]VPNGatewaySpecification]{}

	response, err := s.connection.Get("/ecloud/v2/vpn-gateway-specifications", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNGatewaySpecification retrieves a single VPN gateway specification by ID
func (s *Service) GetVPNGatewaySpecification(specificationID string) (VPNGatewaySpecification, error) {
	body, err := s.getVPNGatewaySpecificationResponseBody(specificationID)

	return body.Data, err
}

func (s *Service) getVPNGatewaySpecificationResponseBody(specificationID string) (*connection.APIResponseBodyData[VPNGatewaySpecification], error) {
	body := &connection.APIResponseBodyData[VPNGatewaySpecification]{}

	if specificationID == "" {
		return body, fmt.Errorf("invalid vpn gateway specification id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-gateway-specifications/%s", specificationID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewaySpecificationNotFoundError{ID: specificationID}
		}

		return nil
	})
}

// GetVPNGatewaySpecificationAvailabilityZones retrieves a list of availability zones for a vpn gateway specification
func (s *Service) GetVPNGatewaySpecificationAvailabilityZones(specificationID string, parameters connection.APIRequestParameters) ([]AvailabilityZone, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
		return s.GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID, p)
	}, parameters)
}

// GetVPNGatewaySpecificationAvailabilityZonesPaginated retrieves a paginated list of availability zones for a vpn gateway specification
func (s *Service) GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID string, parameters connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
	body, err := s.getVPNGatewaySpecificationAvailabilityZonesPaginatedResponseBody(specificationID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[AvailabilityZone], error) {
		return s.GetVPNGatewaySpecificationAvailabilityZonesPaginated(specificationID, p)
	}), err
}

func (s *Service) getVPNGatewaySpecificationAvailabilityZonesPaginatedResponseBody(specificationID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]AvailabilityZone], error) {
	body := &connection.APIResponseBodyData[[]AvailabilityZone]{}

	if specificationID == "" {
		return body, fmt.Errorf("invalid vpn gateway specification id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-gateway-specifications/%s/availability-zones", specificationID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNGatewaySpecificationNotFoundError{ID: specificationID}
		}

		return nil
	})
}
