package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBackupGateways retrieves a list of backup gateways
func (s *Service) GetBackupGateways(parameters connection.APIRequestParameters) ([]BackupGateway, error) {
	return connection.InvokeRequestAll(s.GetBackupGatewaysPaginated, parameters)
}

// GetBackupGatewaysPaginated retrieves a paginated list of backup gateways
func (s *Service) GetBackupGatewaysPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BackupGateway], error) {
	body, err := s.getBackupGatewaysPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetBackupGatewaysPaginated), err
}

func (s *Service) getBackupGatewaysPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]BackupGateway], error) {
	return connection.Get[[]BackupGateway](s.connection, "/ecloud/v2/backup-gateways", parameters)
}

// GetBackupGateway retrieves a single backup gateway by ID
func (s *Service) GetBackupGateway(gatewayID string) (BackupGateway, error) {
	body, err := s.getBackupGatewayResponseBody(gatewayID)

	return body.Data, err
}

func (s *Service) getBackupGatewayResponseBody(gatewayID string) (*connection.APIResponseBodyData[BackupGateway], error) {
	if gatewayID == "" {
		return &connection.APIResponseBodyData[BackupGateway]{}, fmt.Errorf("invalid backup gateway id")
	}

	return connection.Get[BackupGateway](s.connection, fmt.Sprintf("/ecloud/v2/backup-gateways/%s", gatewayID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&BackupGatewayNotFoundError{ID: gatewayID}))
}

// CreateBackupGateway creates a new backup gateway
func (s *Service) CreateBackupGateway(req CreateBackupGatewayRequest) (TaskReference, error) {
	body, err := s.createBackupGatewayResponseBody(req)

	return body.Data, err
}

func (s *Service) createBackupGatewayResponseBody(req CreateBackupGatewayRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	return connection.Post[TaskReference](s.connection, "/ecloud/v2/backup-gateways", &req)
}

// PatchBackupGateway patches a backup gateway
func (s *Service) PatchBackupGateway(gatewayID string, req PatchBackupGatewayRequest) (TaskReference, error) {
	body, err := s.patchBackupGatewayResponseBody(gatewayID, req)

	return body.Data, err
}

func (s *Service) patchBackupGatewayResponseBody(gatewayID string, req PatchBackupGatewayRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	if gatewayID == "" {
		return &connection.APIResponseBodyData[TaskReference]{}, fmt.Errorf("invalid gateway id")
	}

	return connection.Patch[TaskReference](
		s.connection,
		fmt.Sprintf("/ecloud/v2/backup-gateways/%s", gatewayID),
		&req,
		connection.NotFoundResponseHandler(&BackupGatewayNotFoundError{ID: gatewayID}),
	)
}

// DeleteBackupGateway deletes a backup gateway
func (s *Service) DeleteBackupGateway(gatewayID string) (string, error) {
	body, err := s.deleteBackupGatewayResponseBody(gatewayID)

	return body.Data.TaskID, err
}

func (s *Service) deleteBackupGatewayResponseBody(gatewayID string) (*connection.APIResponseBodyData[TaskReference], error) {
	if gatewayID == "" {
		return &connection.APIResponseBodyData[TaskReference]{}, fmt.Errorf("invalid gateway id")
	}

	return connection.Delete[TaskReference](
		s.connection,
		fmt.Sprintf("/ecloud/v2/backup-gateways/%s", gatewayID),
		nil,
		connection.NotFoundResponseHandler(&BackupGatewayNotFoundError{ID: gatewayID}),
	)
}
