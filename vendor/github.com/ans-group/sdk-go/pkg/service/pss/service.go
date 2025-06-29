package pss

import (
	"io"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// PSSService is an interface for managing PSS
type PSSService interface {
	CreateRequest(req CreateRequestRequest) (int, error)
	GetRequests(parameters connection.APIRequestParameters) ([]Request, error)
	GetRequestsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Request], error)
	GetRequest(requestID int) (Request, error)
	PatchRequest(requestID int, req PatchRequestRequest) error

	GetRequestFeedback(requestID int) (Feedback, error)
	CreateRequestFeedback(requestID int, req CreateFeedbackRequest) (int, error)

	CreateRequestReply(requestID int, req CreateReplyRequest) (string, error)
	GetRequestReplies(solutionID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestRepliesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error)
	GetRequestConversation(requestID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestConversationPaginated(requestID int, parameters connection.APIRequestParameters) (*connection.Paginated[Reply], error)

	GetReply(replyID string) (Reply, error)

	DownloadReplyAttachmentStream(replyID string, attachmentName string) (contentStream io.ReadCloser, err error)
	UploadReplyAttachmentStream(replyID string, attachmentName string, fileStream io.Reader) (err error)
	DeleteReplyAttachment(replyID string, attachmentName string) error

	GetCaseCategories(parameters connection.APIRequestParameters) ([]CaseCategory, error)
	GetCaseCategoriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseCategory], error)

	GetChangeImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetChangeImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)
	GetChangeRiskCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetChangeRiskCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)
	GetIncidentImpactCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetIncidentImpactCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)
	GetIncidentTypeCaseOptions(parameters connection.APIRequestParameters) ([]CaseOption, error)
	GetIncidentTypeCaseOptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[CaseOption], error)

	GetSupportedServices(parameters connection.APIRequestParameters) ([]SupportedService, error)
	GetSupportedServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SupportedService], error)

	GetCases(parameters connection.APIRequestParameters) ([]Case, error)
	GetCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Case], error)
	GetCase(caseID string) (Case, error)

	CreateIncidentCase(req CreateIncidentCaseRequest) (string, error)
	GetIncidentCases(parameters connection.APIRequestParameters) ([]IncidentCase, error)
	GetIncidentCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IncidentCase], error)
	GetIncidentCase(incidentID string) (IncidentCase, error)
	CloseIncidentCase(incidentID string, req CloseIncidentCaseRequest) (string, error)

	CreateChangeCase(req CreateChangeCaseRequest) (string, error)
	GetChangeCases(parameters connection.APIRequestParameters) ([]ChangeCase, error)
	GetChangeCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ChangeCase], error)
	GetChangeCase(changeID string) (ChangeCase, error)
	ApproveChangeCase(changeID string, req ApproveChangeCaseRequest) (string, error)

	GetProblemCases(parameters connection.APIRequestParameters) ([]ProblemCase, error)
	GetProblemCasesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ProblemCase], error)
	GetProblemCase(problemID string) (ProblemCase, error)

	GetCaseUpdates(caseID string, parameters connection.APIRequestParameters) ([]CaseUpdate, error)
	GetCaseUpdatesPaginated(caseID string, parameters connection.APIRequestParameters) (*connection.Paginated[CaseUpdate], error)
	GetCaseUpdate(caseID string, updateID string) (CaseUpdate, error)
	CreateCaseUpdate(caseID string, req CreateCaseUpdateRequest) (string, error)
}

// Service implements PSSService for managing
// PSS via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of PSSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
