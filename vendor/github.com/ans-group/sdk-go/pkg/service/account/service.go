package account

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// AccountService is an interface for managing account
type AccountService interface {
	GetClients(parameters connection.APIRequestParameters) ([]Client, error)
	GetClientsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Client], error)
	GetClient(clientID int) (Client, error)
	CreateClient(req CreateClientRequest) (int, error)
	PatchClient(clientID int, patch PatchClientRequest) error
	DeleteClient(clientID int) error

	GetContacts(parameters connection.APIRequestParameters) ([]Contact, error)
	GetContactsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Contact], error)
	GetContact(contactID int) (Contact, error)

	GetDetails() (Details, error)

	GetCredits(parameters connection.APIRequestParameters) ([]Credit, error)

	GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error)
	GetInvoicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Invoice], error)
	GetInvoice(invoiceID int) (Invoice, error)

	GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error)
	GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[InvoiceQuery], error)
	GetInvoiceQuery(invoiceQueryID int) (InvoiceQuery, error)
	CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error)

	GetApplications(parameters connection.APIRequestParameters) ([]Application, error)
	GetApplicationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Application], error)
	GetServices(parameters connection.APIRequestParameters) ([]ApplicationService, error)
	GetServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ApplicationService], error)
	GetApplication(appID string) (Application, error)
	CreateApplication(req CreateApplicationRequest) (CreateApplicationResponse, error)
	UpdateApplication(appID string, req UpdateApplicationRequest) error
	GetApplicationServices(appID string) (ApplicationServiceMapping, error)
	SetApplicationServices(appID string, req SetServiceRequest) error
	DeleteApplicationServices(appID string) error
	GetApplicationRestrictions(appID string) (ApplicationRestriction, error)
	SetApplicationRestrictions(appID string, req SetRestrictionRequest) error
	DeleteApplicationRestrictions(appID string) error
	DeleteApplication(appID string) error
}

// Service implements AccountService for managing
// Account certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of AccountService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
