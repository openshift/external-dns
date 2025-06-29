package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetApplications retrieves a list of applications
func (s *Service) GetApplications(parameters connection.APIRequestParameters) ([]Application, error) {
	return connection.InvokeRequestAll(s.GetApplicationsPaginated, parameters)
}

func (s *Service) GetApplicationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Application], error) {
	body, err := s.getApplicationsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetApplicationsPaginated), err
}

func (s *Service) getApplicationsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Application], error) {
	return connection.Get[[]Application](s.connection, "/account/v1/applications", parameters)
}

// GetApplication retrieves a single application by id
func (s *Service) GetApplication(appID string) (Application, error) {
	body, err := s.getApplicationResponseBody(appID)

	return body.Data, err
}

func (s *Service) getApplicationResponseBody(appID string) (*connection.APIResponseBodyData[Application], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[Application]{}, fmt.Errorf("invalid application id")
	}

	return connection.Get[Application](s.connection, fmt.Sprintf("/account/v1/applications/%s", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

// GetApplications retrieves a list of applications
func (s *Service) GetServices(parameters connection.APIRequestParameters) ([]ApplicationService, error) {
	return connection.InvokeRequestAll(s.GetServicesPaginated, parameters)
}

func (s *Service) GetServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ApplicationService], error) {
	body, err := s.getServicesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetServicesPaginated), err
}

func (s *Service) getServicesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ApplicationService], error) {
	body := &connection.APIResponseBodyData[[]ApplicationService]{}

	response, err := s.connection.Get("/account/v1/services", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

func (s *Service) CreateApplication(req CreateApplicationRequest) (CreateApplicationResponse, error) {
	body, err := s.createApplicationResponseBody(req)

	return body.Data, err
}

func (s *Service) createApplicationResponseBody(req CreateApplicationRequest) (*connection.APIResponseBodyData[CreateApplicationResponse], error) {
	return connection.Post[CreateApplicationResponse](s.connection, "/account/v1/applications", &req)
}

func (s *Service) UpdateApplication(appID string, req UpdateApplicationRequest) error {
	_, err := s.updateApplicationResponseBody(appID, req)

	return err
}

func (s *Service) updateApplicationResponseBody(appID string, req UpdateApplicationRequest) (*connection.APIResponseBodyData[Application], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[Application]{}, fmt.Errorf("invalid application id")
	}

	return connection.Patch[Application](s.connection, fmt.Sprintf("/account/v1/applications/%s", appID), &req, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

// GetApplicationServices retrieves the services and roles of an application by id
func (s *Service) GetApplicationServices(appID string) (ApplicationServiceMapping, error) {
	body, err := s.getApplicationServicesResponseBody(appID)

	return body.Data, err
}

func (s *Service) getApplicationServicesResponseBody(appID string) (*connection.APIResponseBodyData[ApplicationServiceMapping], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[ApplicationServiceMapping]{}, fmt.Errorf("invalid application id")
	}

	return connection.Get[ApplicationServiceMapping](s.connection, fmt.Sprintf("/account/v1/applications/%s/services", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

func (s *Service) SetApplicationServices(appID string, req SetServiceRequest) error {
	_, err := s.setApplicationServicesResponseBody(appID, req)

	return err
}

func (s *Service) setApplicationServicesResponseBody(appID string, req SetServiceRequest) (*connection.APIResponseBodyData[interface{}], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[interface{}]{}, fmt.Errorf("invalid application id")
	}

	return connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/services", appID), &req, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

func (s *Service) DeleteApplicationServices(appID string) error {
	_, err := s.deleteApplicationServicesResponseBody(appID)

	return err
}

func (s *Service) deleteApplicationServicesResponseBody(appID string) (*connection.APIResponseBodyData[interface{}], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[interface{}]{}, fmt.Errorf("invalid application id")
	}

	return connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/services", appID), SetServiceRequest{Scopes: []ApplicationServiceScope{}}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

// DeleteApplication removes an application
func (s *Service) DeleteApplication(appID string) error {
	_, err := s.deleteApplicationResponseBody(appID)

	return err
}

func (s *Service) deleteApplicationResponseBody(appID string) (*connection.APIResponseBodyData[interface{}], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[interface{}]{}, fmt.Errorf("invalid application id")
	}

	return connection.Delete[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

// GetApplicationRestrictions retrieves the IP restrictions of an application by id
func (s *Service) GetApplicationRestrictions(appID string) (ApplicationRestriction, error) {
	body, err := s.getApplicationRestrictionsResponseBody(appID)

	return body.Data, err
}

func (s *Service) getApplicationRestrictionsResponseBody(appID string) (*connection.APIResponseBodyData[ApplicationRestriction], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[ApplicationRestriction]{}, fmt.Errorf("invalid application id")
	}

	return connection.Get[ApplicationRestriction](s.connection, fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

func (s *Service) SetApplicationRestrictions(appID string, req SetRestrictionRequest) error {
	_, err := s.setApplicationRestrictionsResponseBody(appID, req)

	return err
}

func (s *Service) setApplicationRestrictionsResponseBody(appID string, req SetRestrictionRequest) (*connection.APIResponseBodyData[interface{}], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[interface{}]{}, fmt.Errorf("invalid application id")
	}

	return connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), &req, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}

func (s *Service) DeleteApplicationRestrictions(appID string) error {
	_, err := s.deleteApplicationRestrictionsResponseBody(appID)

	return err
}

func (s *Service) deleteApplicationRestrictionsResponseBody(appID string) (*connection.APIResponseBodyData[interface{}], error) {
	if appID == "" {
		return &connection.APIResponseBodyData[interface{}]{}, fmt.Errorf("invalid application id")
	}

	return connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
}
