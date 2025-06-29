package pss

import "github.com/ans-group/sdk-go/pkg/connection"

// GetCaseCategories retrieves a list of case categories
func (s *Service) GetCaseCategories(parameters connection.APIRequestParameters) ([]CaseCategory, error) {
	return connection.InvokeRequestAll(s.GetCaseCategoriesPaginated, parameters)
}

// GetCaseCategoriesPaginated retrieves a paginated list of case categories
func (s *Service) GetCaseCategoriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseCategory], error) {
	body, err := s.getCaseCategoriesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetCaseCategoriesPaginated), err
}

func (s *Service) getCaseCategoriesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CaseCategory], error) {
	return connection.Get[[]CaseCategory](s.connection, "/pss/v2/case-categories", parameters)
}
