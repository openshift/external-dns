package pss

import "github.com/ans-group/sdk-go/pkg/connection"

type AuthorType string

func (s AuthorType) String() string {
	return string(s)
}

const (
	AuthorTypeClient  AuthorType = "Client"
	AuthorTypeAuto    AuthorType = "Auto"
	AuthorTypeSupport AuthorType = "Support"
)

var AuthorTypeEnum connection.Enum[AuthorType] = []AuthorType{AuthorTypeClient, AuthorTypeAuto, AuthorTypeSupport}

type RequestPriority string

func (s RequestPriority) String() string {
	return string(s)
}

const (
	RequestPriorityNormal   RequestPriority = "Normal"
	RequestPriorityHigh     RequestPriority = "High"
	RequestPriorityCritical RequestPriority = "Critical"
)

var RequestPriorityEnum connection.Enum[RequestPriority] = []RequestPriority{RequestPriorityNormal, RequestPriorityHigh, RequestPriorityCritical}

type RequestStatus string

func (s RequestStatus) String() string {
	return string(s)
}

const (
	RequestStatusCompleted                RequestStatus = "Completed"
	RequestStatusAwaitingCustomerResponse RequestStatus = "Awaiting Customer Response"
	RequestStatusRepliedAndCompleted      RequestStatus = "Replied and Completed"
	RequestStatusSubmitted                RequestStatus = "Submitted"
)

var RequestStatusEnum connection.Enum[RequestStatus] = []RequestStatus{
	RequestStatusCompleted,
	RequestStatusAwaitingCustomerResponse,
	RequestStatusRepliedAndCompleted,
	RequestStatusSubmitted,
}

// Request represents a PSS request
type Request struct {
	ID                int                 `json:"id"`
	Author            Author              `json:"author"`
	Type              string              `json:"type"`
	Secure            bool                `json:"secure"`
	Subject           string              `json:"subject"`
	CreatedAt         connection.DateTime `json:"created_at"`
	Priority          RequestPriority     `json:"priority"`
	Archived          bool                `json:"archived"`
	Status            RequestStatus       `json:"status"`
	RequestSMS        bool                `json:"request_sms"`
	Version           int                 `json:"version"`
	CustomerReference string              `json:"customer_reference"`
	Product           Product             `json:"product"`
	LastRepliedAt     connection.DateTime `json:"last_replied_at"`
	CC                []string            `json:"cc"`
	UnreadReplies     int                 `json:"unread_replies"`
	ContactMethod     string              `json:"contact_method"`
}

// Author represents a PSS request author
type Author struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Type AuthorType `json:"type"`
}

// Reply represents a PSS reply
type Reply struct {
	ID          string              `json:"id"`
	RequestID   int                 `json:"request_id"`
	Author      Author              `json:"author"`
	Description string              `json:"description"`
	Attachments []Attachment        `json:"attachments"`
	Read        bool                `json:"read"`
	CreatedAt   connection.DateTime `json:"created_at"`
}

// Attachment represents a PSS attachment
type Attachment struct {
	Name string `json:"name"`
}

// Product represents a product to which the request applies to
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Feedback represents PSS feedback
type Feedback struct {
	ID               int                 `json:"id"`
	ContactID        int                 `json:"contact_id"`
	Score            int                 `json:"score"`
	Comment          string              `json:"comment"`
	SpeedResolved    int                 `json:"speed_resolved"`
	Quality          int                 `json:"quality"`
	NPSScore         int                 `json:"nps_score"`
	ThirdPartConsent bool                `json:"thirdparty_consent"`
	CreatedAt        connection.DateTime `json:"created_at"`
}

// CaseCategory represents a PSS case category
type CaseCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CaseOption represents a PSS case option
type CaseOption struct {
	Option string `json:"option"`
}

type CaseType string

func (s CaseType) String() string {
	return string(s)
}

const (
	CaseTypeChange   CaseType = "Change"
	CaseTypeIncident CaseType = "Incident"
	CaseTypeProblem  CaseType = "Problem"
)

var CaseTypeEnum connection.Enum[CaseType] = []CaseType{
	CaseTypeChange,
	CaseTypeIncident,
	CaseTypeProblem,
}

type CaseStatus string

func (s CaseStatus) String() string {
	return string(s)
}

const (
	CaseStatusInProgress                   CaseStatus = "In Progress"
	CaseStatusOnHold                       CaseStatus = "On Hold"
	CaseStatusOpen                         CaseStatus = "Open"
	CaseStatusCaseCreated                  CaseStatus = "Case Created"
	CaseStatusResolved                     CaseStatus = "Resolved"
	CaseStatusPendingAssessment            CaseStatus = "Pending Assessment"
	CaseStatusPendingSME                   CaseStatus = "Pending SME"
	CaseStatusProblemAssessment            CaseStatus = "Problem Assessment"
	CaseStatusClosed                       CaseStatus = "Closed"
	CaseStatusPendingCustomersApproval     CaseStatus = "Pending Customers Approval"
	CaseStatusPendingANSCAB                CaseStatus = "Pending ANS CAB"
	CaseStatusPendingImplementation        CaseStatus = "Pending Implementation"
	CaseStatusImplementationStarted        CaseStatus = "Implementation Started"
	CaseStatusDelivered                    CaseStatus = "Delivered"
	CaseStatusRejected                     CaseStatus = "Rejected"
	CaseStatusScheduleChange               CaseStatus = "Schedule Change"
	CaseStatusCustomerClosure              CaseStatus = "Customer Closure"
	CaseStatusProblemSolved                CaseStatus = "Problem Solved"
	CaseStatusInformationProvided          CaseStatus = "Information Provided"
	CaseStatusClosedWithCustomerConsent    CaseStatus = "Closed with Customer Consent"
	CaseStatusClosedCustomerNotContactable CaseStatus = "Closed - Customer Not Contactable"
	CaseStatusClosedWithTeamLeaderApproval CaseStatus = "Closed with Team Leader Approval"
	CaseStatusCaseClosed                   CaseStatus = "Case Closed"
	CaseStatusANSOperatorError             CaseStatus = "ANS Operator Error"
	CaseStatusCustomerError                CaseStatus = "Customer Error"
	CaseStatusOutofScopeofServiceContract  CaseStatus = "Out of Scope of Service Contract"
)

var CaseStatusEnum connection.Enum[CaseStatus] = []CaseStatus{
	CaseStatusInProgress,
	CaseStatusOnHold,
	CaseStatusOpen,
	CaseStatusCaseCreated,
	CaseStatusResolved,
	CaseStatusPendingAssessment,
	CaseStatusPendingSME,
	CaseStatusProblemAssessment,
	CaseStatusClosed,
	CaseStatusPendingCustomersApproval,
	CaseStatusPendingANSCAB,
	CaseStatusPendingImplementation,
	CaseStatusImplementationStarted,
	CaseStatusDelivered,
	CaseStatusRejected,
	CaseStatusScheduleChange,
	CaseStatusCustomerClosure,
	CaseStatusProblemSolved,
	CaseStatusInformationProvided,
	CaseStatusClosedWithCustomerConsent,
	CaseStatusClosedCustomerNotContactable,
	CaseStatusClosedWithTeamLeaderApproval,
	CaseStatusCaseClosed,
	CaseStatusANSOperatorError,
	CaseStatusCustomerError,
	CaseStatusOutofScopeofServiceContract,
}

type ChangeCaseType string

func (s ChangeCaseType) String() string {
	return string(s)
}

const (
	ChangeCaseTypeNormal             ChangeCaseType = "Normal"
	ChangeCaseTypeStandard           ChangeCaseType = "Standard"
	ChangeCaseTypeEmergency          ChangeCaseType = "Emergency"
	ChangeCaseTypeProject            ChangeCaseType = "Project"
	ChangeCaseTypeCentreOfExcellence ChangeCaseType = "Centre of Excellence"
)

var ChangeCaseTypeEnum connection.Enum[ChangeCaseType] = []ChangeCaseType{
	ChangeCaseTypeNormal,
	ChangeCaseTypeStandard,
	ChangeCaseTypeEmergency,
	ChangeCaseTypeProject,
	ChangeCaseTypeCentreOfExcellence,
}

type ChangeCasePriority string

func (s ChangeCasePriority) String() string {
	return string(s)
}

const (
	ChangeCasePriorityCR1 ChangeCasePriority = "CR1"
	ChangeCasePriorityCR2 ChangeCasePriority = "CR2"
	ChangeCasePriorityCR3 ChangeCasePriority = "CR3"
	ChangeCasePriorityCR4 ChangeCasePriority = "CR4"
	ChangeCasePriorityCR5 ChangeCasePriority = "CR5"
	ChangeCasePriorityCRS ChangeCasePriority = "CR-S"
	ChangeCasePriorityCRE ChangeCasePriority = "CR-E"
	ChangeCasePriorityCRP ChangeCasePriority = "CR-P"
)

var ChangeCasePriorityEnum connection.Enum[ChangeCasePriority] = []ChangeCasePriority{
	ChangeCasePriorityCR1,
	ChangeCasePriorityCR2,
	ChangeCasePriorityCR3,
	ChangeCasePriorityCR4,
	ChangeCasePriorityCR5,
	ChangeCasePriorityCRS,
	ChangeCasePriorityCRE,
	ChangeCasePriorityCRP,
}

type ChangeCaseStage string

func (s ChangeCaseStage) String() string {
	return string(s)
}

const (
	ChangeCaseStageCaseType              ChangeCaseStage = "Case Type"
	ChangeCaseStageIdentify              ChangeCaseStage = "Identify"
	ChangeCaseStagePendingAssessment     ChangeCaseStage = "Pending Assessment"
	ChangeCaseStageApprovals             ChangeCaseStage = "Approvals"
	ChangeCaseStagePendingImplementation ChangeCaseStage = "Pending Implementation"
	ChangeCaseStageImplementationStarted ChangeCaseStage = "Implementation Started"
	ChangeCaseStageDelivered             ChangeCaseStage = "Delivered"
)

var ChangeCaseStageEnum connection.Enum[ChangeCaseStage] = []ChangeCaseStage{
	ChangeCaseStageCaseType,
	ChangeCaseStageIdentify,
	ChangeCaseStagePendingAssessment,
	ChangeCaseStageApprovals,
	ChangeCaseStagePendingImplementation,
	ChangeCaseStageImplementationStarted,
	ChangeCaseStageDelivered,
}

type ChangeCaseImpact string

func (s ChangeCaseImpact) String() string {
	return string(s)
}

const (
	ChangeCaseImpactLow    ChangeCaseImpact = "Low"
	ChangeCaseImpactMedium ChangeCaseImpact = "Medium"
	ChangeCaseImpactHigh   ChangeCaseImpact = "High"
)

var ChangeCaseImpactEnum connection.Enum[ChangeCaseImpact] = []ChangeCaseImpact{
	ChangeCaseImpactLow,
	ChangeCaseImpactMedium,
	ChangeCaseImpactHigh,
}

type ChangeCaseRisk string

func (s ChangeCaseRisk) String() string {
	return string(s)
}

const (
	ChangeCaseRiskLow    ChangeCaseRisk = "Low"
	ChangeCaseRiskMedium ChangeCaseRisk = "Medium"
	ChangeCaseRiskHigh   ChangeCaseRisk = "High"
)

var ChangeCaseRiskEnum connection.Enum[ChangeCaseRisk] = []ChangeCaseRisk{
	ChangeCaseRiskLow,
	ChangeCaseRiskMedium,
	ChangeCaseRiskHigh,
}

type IncidentCaseType string

func (s IncidentCaseType) String() string {
	return string(s)
}

const (
	IncidentCaseTypeFault                IncidentCaseType = "Fault"
	IncidentCaseTypeServiceRequest       IncidentCaseType = "Service Request"
	IncidentCaseTypeChangeAdvisory       IncidentCaseType = "Change Advisory"
	IncidentCaseTypeProjectIncident      IncidentCaseType = "Project Incident"
	IncidentCaseTypeAccessRequest        IncidentCaseType = "Access Request"
	IncidentCaseTypeScheduledMaintenance IncidentCaseType = "Scheduled Maintenance"
	IncidentCaseTypeArchitecturalAdvice  IncidentCaseType = "Architectural Advice"
	IncidentCaseTypeSecurityEvent        IncidentCaseType = "Security Event"
)

var IncidentCaseTypeEnum connection.Enum[IncidentCaseType] = []IncidentCaseType{
	IncidentCaseTypeFault,
	IncidentCaseTypeServiceRequest,
	IncidentCaseTypeChangeAdvisory,
	IncidentCaseTypeProjectIncident,
	IncidentCaseTypeAccessRequest,
	IncidentCaseTypeScheduledMaintenance,
	IncidentCaseTypeArchitecturalAdvice,
	IncidentCaseTypeSecurityEvent,
}

type IncidentCasePriority string

func (s IncidentCasePriority) String() string {
	return string(s)
}

const (
	IncidentCasePriorityP1 IncidentCasePriority = "P1"
	IncidentCasePriorityP2 IncidentCasePriority = "P2"
	IncidentCasePriorityP3 IncidentCasePriority = "P3"
	IncidentCasePriorityP4 IncidentCasePriority = "P4"
	IncidentCasePriorityP5 IncidentCasePriority = "P5"
)

var IncidentCasePriorityEnum connection.Enum[IncidentCasePriority] = []IncidentCasePriority{
	IncidentCasePriorityP1,
	IncidentCasePriorityP2,
	IncidentCasePriorityP3,
	IncidentCasePriorityP4,
	IncidentCasePriorityP5,
}

type IncidentCaseImpact string

func (s IncidentCaseImpact) String() string {
	return string(s)
}

const (
	IncidentCaseImpactMajor    IncidentCaseImpact = "Major"
	IncidentCaseImpactModerate IncidentCaseImpact = "Moderate"
	IncidentCaseImpactMinor    IncidentCaseImpact = "Minor"
)

var IncidentCaseImpactEnum connection.Enum[IncidentCaseImpact] = []IncidentCaseImpact{
	IncidentCaseImpactMajor,
	IncidentCaseImpactModerate,
	IncidentCaseImpactMinor,
}

type ProblemCaseType string

func (s ProblemCaseType) String() string {
	return string(s)
}

const (
	ProblemCaseTypeRCA              ProblemCaseType = "RCA"
	ProblemCaseTypeKnownError       ProblemCaseType = "Known Error"
	ProblemCaseTypeVulnerability    ProblemCaseType = "Vulnerability"
	ProblemCaseTypeIssueReOccurence ProblemCaseType = "Issue Re-Occurence"
	ProblemCaseTypeNonCompliance    ProblemCaseType = "Non-Compliance"
	ProblemCaseTypeBugFix           ProblemCaseType = "Bug Fix"
)

var ProblemCaseTypeEnum connection.Enum[ProblemCaseType] = []ProblemCaseType{
	ProblemCaseTypeRCA,
	ProblemCaseTypeKnownError,
	ProblemCaseTypeVulnerability,
	ProblemCaseTypeIssueReOccurence,
	ProblemCaseTypeNonCompliance,
	ProblemCaseTypeBugFix,
}

type ProblemCasePriority string

func (s ProblemCasePriority) String() string {
	return string(s)
}

const (
	ProblemCasePriorityPRB1   ProblemCasePriority = "PRB1"
	ProblemCasePriorityPRB2   ProblemCasePriority = "PRB2"
	ProblemCasePriorityPRB3   ProblemCasePriority = "PRB3"
	ProblemCasePriorityPRB4   ProblemCasePriority = "PRB4"
	ProblemCasePriorityPRB5   ProblemCasePriority = "PRB5"
	ProblemCasePriorityPRBRCA ProblemCasePriority = "PRB-RCA"
	ProblemCasePriorityPRBMI  ProblemCasePriority = "PRB-MI"
)

var ProblemCasePriorityEnum connection.Enum[ProblemCasePriority] = []ProblemCasePriority{
	ProblemCasePriorityPRB1,
	ProblemCasePriorityPRB2,
	ProblemCasePriorityPRB3,
	ProblemCasePriorityPRB4,
	ProblemCasePriorityPRB5,
	ProblemCasePriorityPRBRCA,
	ProblemCasePriorityPRBMI,
}

type ProblemCaseUrgency string

func (s ProblemCaseUrgency) String() string {
	return string(s)
}

const (
	ProblemCaseUrgencySystemServiceDown ProblemCaseUrgency = "System / Service Down"
	ProblemCaseUrgencySystemAffected    ProblemCaseUrgency = "System / Service Affected"
	ProblemCaseUrgencyUserAffected      ProblemCaseUrgency = "User Down / Affected"
)

var ProblemCaseUrgencyEnum connection.Enum[ProblemCaseUrgency] = []ProblemCaseUrgency{
	ProblemCaseUrgencySystemServiceDown,
	ProblemCaseUrgencySystemAffected,
	ProblemCaseUrgencyUserAffected,
}

type ProblemCaseDetailedImpact string

func (s ProblemCaseDetailedImpact) String() string {
	return string(s)
}

const (
	ProblemCaseDetailedImpactMajor    ProblemCaseDetailedImpact = "Major"
	ProblemCaseDetailedImpactModerate ProblemCaseDetailedImpact = "Moderate"
	ProblemCaseDetailedImpactMinor    ProblemCaseDetailedImpact = "Minor"
)

var ProblemCaseDetailedImpactEnum connection.Enum[ProblemCaseDetailedImpact] = []ProblemCaseDetailedImpact{
	ProblemCaseDetailedImpactMajor,
	ProblemCaseDetailedImpactModerate,
	ProblemCaseDetailedImpactMinor,
}

type ProblemCaseKnownWorkaround string

func (s ProblemCaseKnownWorkaround) String() string {
	return string(s)
}

const (
	ProblemCaseKnownWorkaroundNotCurrentlyKnown                  ProblemCaseKnownWorkaround = "Not Currently Known"
	ProblemCaseKnownWorkaroundTemporaryWorkaround                ProblemCaseKnownWorkaround = "Temporary Workaround"
	ProblemCaseKnownWorkaroundPermanentWorkaround                ProblemCaseKnownWorkaround = "Permanent Workaround"
	ProblemCaseKnownWorkaroundCurrentlyUnavailableAwaitingVendor ProblemCaseKnownWorkaround = "Currently Unavailable / Awaiting Vendor"
)

var ProblemCaseKnownWorkaroundEnum connection.Enum[ProblemCaseKnownWorkaround] = []ProblemCaseKnownWorkaround{
	ProblemCaseKnownWorkaroundNotCurrentlyKnown,
	ProblemCaseKnownWorkaroundTemporaryWorkaround,
	ProblemCaseKnownWorkaroundPermanentWorkaround,
	ProblemCaseKnownWorkaroundCurrentlyUnavailableAwaitingVendor,
}

type ProblemCaseKnownCause string

func (s ProblemCaseKnownCause) String() string {
	return string(s)
}

const (
	ProblemCaseKnownCauseNotCurrentlyKnown     ProblemCaseKnownCause = "Not Currently Known"
	ProblemCaseKnownCauseSoftwareDefect        ProblemCaseKnownCause = "Software Defect"
	ProblemCaseKnownCauseHardwareFailure       ProblemCaseKnownCause = "Hardware Failure"
	ProblemCaseKnownCauseCommunicationsFailure ProblemCaseKnownCause = "Communications Failure"
	ProblemCaseKnownCauseHardwareLimitation    ProblemCaseKnownCause = "Hardware Limitation"
	ProblemCaseKnownCauseBusinessProcessIssue  ProblemCaseKnownCause = "Business Process Issue"
	ProblemCaseKnownCauseUser                  ProblemCaseKnownCause = "User"
)

var ProblemCaseKnownCauseEnum connection.Enum[ProblemCaseKnownCause] = []ProblemCaseKnownCause{
	ProblemCaseKnownCauseNotCurrentlyKnown,
	ProblemCaseKnownCauseSoftwareDefect,
	ProblemCaseKnownCauseHardwareFailure,
	ProblemCaseKnownCauseCommunicationsFailure,
	ProblemCaseKnownCauseHardwareLimitation,
	ProblemCaseKnownCauseBusinessProcessIssue,
	ProblemCaseKnownCauseUser,
}

type Case struct {
	ID                 string     `json:"id"`
	CaseType           CaseType   `json:"case_type"`
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	IsSecurity         bool       `json:"is_security"`
	Squad              string     `json:"squad"`
	Status             CaseStatus `json:"status"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	CategoryID         string     `json:"category_id"`
	SupportedServiceID string     `json:"supported_service_id"`
}

// IncidentCase represents a PSS incident-type case
type IncidentCase struct {
	ID                    string               `json:"id"`
	CaseType              CaseType             `json:"case_type"`
	Title                 string               `json:"title"`
	Description           string               `json:"description"`
	IsSecurity            bool                 `json:"is_security"`
	Type                  IncidentCaseType     `json:"type"`
	Priority              IncidentCasePriority `json:"priority"`
	Squad                 string               `json:"squad"`
	Status                CaseStatus           `json:"status"`
	CustomerReference     string               `json:"customer_reference"`
	Source                string               `json:"source"`
	ContactID             int                  `json:"contact_id"`
	Contact               string               `json:"contact"`
	CriteriaForResolution string               `json:"criteria_for_resolution"`
	ServiceName           string               `json:"service_name"`
	Owner                 string               `json:"owner"`
	ResponseTarget        string               `json:"response_target"`
	ResolutionTarget      string               `json:"resolution_target"`
	CreatedAt             string               `json:"created_at"`
	UpdatedAt             string               `json:"updated_at"`
	CategoryID            string               `json:"category_id"`
	SupportedServiceID    string               `json:"supported_service_id"`
	Impact                IncidentCaseImpact   `json:"impact"`
}

// IncidentCase represents a PSS change-type case
type ChangeCase struct {
	ID                         string             `json:"id"`
	CaseType                   CaseType           `json:"case_type"`
	Title                      string             `json:"title"`
	Description                string             `json:"description"`
	IsSecurity                 bool               `json:"is_security"`
	Type                       ChangeCaseType     `json:"type"`
	Priority                   ChangeCasePriority `json:"priority"`
	Squad                      string             `json:"squad"`
	Status                     CaseStatus         `json:"status"`
	Contact                    string             `json:"contact"`
	ScheduledDate              string             `json:"scheduled_date"`
	CustomerReference          string             `json:"customer_reference"`
	Owner                      string             `json:"owner"`
	Stage                      ChangeCaseStage    `json:"stage"`
	Reason                     string             `json:"reason"`
	Impact                     ChangeCaseImpact   `json:"impact"`
	Risk                       ChangeCaseRisk     `json:"risk"`
	ServiceName                string             `json:"service_name"`
	EstimatedTimeForCompletion int                `json:"estimated_time_for_completion"`
	EstimatedDowntime          int                `json:"estimated_downtime"`
	ImpactAndPotentialRisks    string             `json:"impact_and_potential_risks"`
	BackoutPlan                string             `json:"backout_plan"`
	ImplementationSteps        string             `json:"implementation_steps"`
	TestSteps                  string             `json:"test_steps"`
	CreatedAt                  string             `json:"created_at"`
	UpdatedAt                  string             `json:"updated_at"`
	CategoryID                 string             `json:"category_id"`
	SupportedServiceID         string             `json:"supported_service_id"`
}

// IncidentCase represents a PSS problem-type case
type ProblemCase struct {
	ID                     string                     `json:"id"`
	CaseType               CaseType                   `json:"case_type"`
	Title                  string                     `json:"title"`
	Description            string                     `json:"description"`
	IsSecurity             bool                       `json:"is_security"`
	Type                   ProblemCaseType            `json:"type"`
	Priority               ProblemCasePriority        `json:"priority"`
	Squad                  string                     `json:"squad"`
	Status                 CaseStatus                 `json:"status"`
	CurrentAssignee        string                     `json:"current_assignee"`
	Urgency                ProblemCaseUrgency         `json:"urgency"`
	DetailedImpact         ProblemCaseDetailedImpact  `json:"detailed_impact"`
	KnownWorkaround        ProblemCaseKnownWorkaround `json:"known_workaround"`
	KnownWorkaroundDetails string                     `json:"known_workaround_details"`
	KnownCause             ProblemCaseKnownCause      `json:"known_cause"`
	CreatedAt              string                     `json:"created_at"`
	UpdatedAt              string                     `json:"updated_at"`
}

// CaseUpdate represents a PSS case update
type CaseUpdate struct {
	ID          string `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CaseUpdate represents a PSS supported service
type SupportedService struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
