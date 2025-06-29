package pss

import "github.com/ans-group/sdk-go/pkg/connection"

// GetCaseOptions retrieves a list of change impact case options
func (s *Service) GetChangeImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return connection.InvokeRequestAll(s.GetChangeImpactCaseOptionsPaginated, parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of change impact case options
func (s *Service) GetChangeImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	body, err := s.getChangeImpactCaseOptionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetChangeImpactCaseOptionsPaginated), err
}

func (s *Service) getChangeImpactCaseOptionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CaseOption], error) {
	return connection.Get[[]CaseOption](s.connection, "/pss/v2/case-options/change-impacts", parameters)
}

// GetCaseOptions retrieves a list of change risk case options
func (s *Service) GetChangeRiskCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return connection.InvokeRequestAll(s.GetChangeRiskCaseOptionsPaginated, parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of change risk case options
func (s *Service) GetChangeRiskCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	body, err := s.getChangeRiskCaseOptionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetChangeRiskCaseOptionsPaginated), err
}

func (s *Service) getChangeRiskCaseOptionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CaseOption], error) {
	return connection.Get[[]CaseOption](s.connection, "/pss/v2/case-options/change-risks", parameters)
}

// GetCaseOptions retrieves a list of incident impact case options
func (s *Service) GetIncidentImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return connection.InvokeRequestAll(s.GetIncidentImpactCaseOptionsPaginated, parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of incident impact case options
func (s *Service) GetIncidentImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	body, err := s.getIncidentImpactCaseOptionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetIncidentImpactCaseOptionsPaginated), err
}

func (s *Service) getIncidentImpactCaseOptionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CaseOption], error) {
	return connection.Get[[]CaseOption](s.connection, "/pss/v2/case-options/incident-impacts", parameters)
}

// GetCaseOptions retrieves a list of incident type case options
func (s *Service) GetIncidentTypeCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error) {
	return connection.InvokeRequestAll(s.GetIncidentTypeCaseOptionsPaginated, parameters)
}

// GetCaseOptionsPaginated retrieves a paginated list of incident type case options
func (s *Service) GetIncidentTypeCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error) {
	body, err := s.getIncidentTypeCaseOptionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetIncidentTypeCaseOptionsPaginated), err
}

func (s *Service) getIncidentTypeCaseOptionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CaseOption], error) {
	return connection.Get[[]CaseOption](s.connection, "/pss/v2/case-options/incident-types", parameters)
}
