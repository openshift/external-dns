package ddosx

import (
	"time"

	"github.com/ans-group/go-durationstring"
	"github.com/ans-group/sdk-go/pkg/connection"
)

type DomainStatus string

const (
	DomainStatusConfigured    DomainStatus = "Configured"
	DomainStatusNotConfigured DomainStatus = "Not Configured"
	DomainStatusPending       DomainStatus = "Pending"
	DomainStatusFailed        DomainStatus = "Failed"
	DomainStatusCancelling    DomainStatus = "Cancelling"
	DomainStatusCancelled     DomainStatus = "Cancelled"
)

type DomainPropertyName string

func (s DomainPropertyName) String() string {
	return string(s)
}

const (
	DomainPropertyNameClientMaxBodySize DomainPropertyName = "client_max_body_size"
	DomainPropertyNameProxyTimeout      DomainPropertyName = "proxy_timeout"
	DomainPropertyNameIPv6Enabled       DomainPropertyName = "ipv6_enabled"
	DomainPropertyNameSecureOrigin      DomainPropertyName = "secure_origin"
)

var DomainPropertyNameEnum connection.Enum[DomainPropertyName] = []DomainPropertyName{
	DomainPropertyNameClientMaxBodySize,
	DomainPropertyNameProxyTimeout,
	DomainPropertyNameIPv6Enabled,
	DomainPropertyNameSecureOrigin,
}

type RecordType string

func (s RecordType) String() string {
	return string(s)
}

const (
	RecordTypeA    RecordType = "A"
	RecordTypeAAAA RecordType = "AAAA"
)

var RecordTypeEnum connection.Enum[RecordType] = []RecordType{
	RecordTypeA,
	RecordTypeAAAA,
}

type WAFMode string

func (s WAFMode) String() string {
	return string(s)
}

const (
	WAFModeOn            WAFMode = "On"
	WAFModeOff           WAFMode = "Off"
	WAFModeDetectionOnly WAFMode = "DetectionOnly"
)

var WAFModeEnum connection.Enum[WAFMode] = []WAFMode{
	WAFModeOn,
	WAFModeOff,
	WAFModeDetectionOnly,
}

type WAFParanoiaLevel string

func (s WAFParanoiaLevel) String() string {
	return string(s)
}

const (
	WAFParanoiaLevelLow     WAFParanoiaLevel = "Low"
	WAFParanoiaLevelMedium  WAFParanoiaLevel = "Medium"
	WAFParanoiaLevelHigh    WAFParanoiaLevel = "High"
	WAFParanoiaLevelHighest WAFParanoiaLevel = "Highest"
)

var WAFParanoiaLevelEnum connection.Enum[WAFParanoiaLevel] = []WAFParanoiaLevel{
	WAFParanoiaLevelLow,
	WAFParanoiaLevelMedium,
	WAFParanoiaLevelHigh,
	WAFParanoiaLevelHighest,
}

type WAFRuleSetName string

const (
	WAFRuleSetNameIPRepution                             WAFRuleSetName = "IP Reputation"
	WAFRuleSetNameMethodEnforcement                      WAFRuleSetName = "Method Enforcement"
	WAFRuleSetNameScannerDetection                       WAFRuleSetName = "Scanner Detection"
	WAFRuleSetNameProtocolEnforcement                    WAFRuleSetName = "Protocol Enforcement"
	WAFRuleSetNameProtocolAttack                         WAFRuleSetName = "Protocol Attack"
	WAFRuleSetNameApplicationAttackLocalFileInclusion    WAFRuleSetName = "Application Attack (Local File Inclusion)"
	WAFRuleSetNameApplicationAttackRemoteFileInclusion   WAFRuleSetName = "Application Attack (Remote File Inclusion)"
	WAFRuleSetNameApplicationAttackRemoteCodeExecution   WAFRuleSetName = "Application Attack (Remote Code Execution)"
	WAFRuleSetNameApplicationAttackPHP                   WAFRuleSetName = "Application Attack PHP"
	WAFRuleSetNameApplicationAttackXSSCrossSiteScripting WAFRuleSetName = "Application Attack XSS (Cross Site Scripting)"
	WAFRuleSetNameApplicationAttackSQLISQLInjection      WAFRuleSetName = "Application Attack SQLI (SQL Injection)"
	WAFRuleSetNameApplicationAttackSessionFixation       WAFRuleSetName = "Application Attack Session Fixation"
	WAFRuleSetNameDataDeakages                           WAFRuleSetName = "Data Leakages"
	WAFRuleSetNameDataLeakageSQL                         WAFRuleSetName = "Data Leakage SQL"
	WAFRuleSetNameDataLeakageJava                        WAFRuleSetName = "Data Leakage Java"
	WAFRuleSetNameDataLeakagePHP                         WAFRuleSetName = "Data Leakage PHP"
	WAFRuleSetNameDataLeakageIIS                         WAFRuleSetName = "Data Leakage IIS"
)

type WAFAdvancedRuleSection string

func (s WAFAdvancedRuleSection) String() string {
	return string(s)
}

const (
	WAFAdvancedRuleSectionArgs           WAFAdvancedRuleSection = "ARGS"
	WAFAdvancedRuleSectionMatchedVars    WAFAdvancedRuleSection = "MATCHED_VARS"
	WAFAdvancedRuleSectionRemoteHost     WAFAdvancedRuleSection = "REMOTE_HOST"
	WAFAdvancedRuleSectionRequestBody    WAFAdvancedRuleSection = "REQUEST_BODY"
	WAFAdvancedRuleSectionRequestCookies WAFAdvancedRuleSection = "REQUEST_COOKIES"
	WAFAdvancedRuleSectionRequestHeaders WAFAdvancedRuleSection = "REQUEST_HEADERS"
	WAFAdvancedRuleSectionRequestURI     WAFAdvancedRuleSection = "REQUEST_URI"
)

var WAFAdvancedRuleSectionEnum connection.Enum[WAFAdvancedRuleSection] = []WAFAdvancedRuleSection{
	WAFAdvancedRuleSectionArgs,
	WAFAdvancedRuleSectionMatchedVars,
	WAFAdvancedRuleSectionRemoteHost,
	WAFAdvancedRuleSectionRequestBody,
	WAFAdvancedRuleSectionRequestCookies,
	WAFAdvancedRuleSectionRequestHeaders,
	WAFAdvancedRuleSectionRequestURI,
}

type WAFAdvancedRuleModifier string

func (s WAFAdvancedRuleModifier) String() string {
	return string(s)
}

const (
	WAFAdvancedRuleModifierBeginsWith   WAFAdvancedRuleModifier = "beginsWith"
	WAFAdvancedRuleModifierEndsWith     WAFAdvancedRuleModifier = "endsWith"
	WAFAdvancedRuleModifierContains     WAFAdvancedRuleModifier = "contains"
	WAFAdvancedRuleModifierContainsWord WAFAdvancedRuleModifier = "containsWord"
)

var WAFAdvancedRuleModifierEnum connection.Enum[WAFAdvancedRuleModifier] = []WAFAdvancedRuleModifier{
	WAFAdvancedRuleModifierBeginsWith,
	WAFAdvancedRuleModifierEndsWith,
	WAFAdvancedRuleModifierContains,
	WAFAdvancedRuleModifierContainsWord,
}

type ACLIPMode string

func (s ACLIPMode) String() string {
	return string(s)
}

const (
	ACLIPModeAllow ACLIPMode = "Allow"
	ACLIPModeDeny  ACLIPMode = "Deny"
)

var ACLIPModeEnum connection.Enum[ACLIPMode] = []ACLIPMode{
	ACLIPModeAllow,
	ACLIPModeDeny,
}

type ACLGeoIPRulesMode string

func (s ACLGeoIPRulesMode) String() string {
	return string(s)
}

const (
	ACLGeoIPRulesModeWhitelist ACLGeoIPRulesMode = "Whitelist"
	ACLGeoIPRulesModeBlacklist ACLGeoIPRulesMode = "Blacklist"
)

var ACLGeoIPRulesModeEnum connection.Enum[ACLGeoIPRulesMode] = []ACLGeoIPRulesMode{
	ACLGeoIPRulesModeWhitelist,
	ACLGeoIPRulesModeBlacklist,
}

type CDNRuleCacheControl string

func (s CDNRuleCacheControl) String() string {
	return string(s)
}

const (
	CDNRuleCacheControlCustom CDNRuleCacheControl = "Custom"
	CDNRuleCacheControlOrigin CDNRuleCacheControl = "Origin"
)

var CDNRuleCacheControlEnum connection.Enum[CDNRuleCacheControl] = []CDNRuleCacheControl{
	CDNRuleCacheControlCustom,
	CDNRuleCacheControlOrigin,
}

type CDNRuleType string

func (s CDNRuleType) String() string {
	return string(s)
}

const (
	CDNRuleTypeGlobal CDNRuleType = "global"
	CDNRuleTypePerURI CDNRuleType = "per-uri"
)

var CDNRuleTypeEnum connection.Enum[CDNRuleType] = []CDNRuleType{CDNRuleTypeGlobal, CDNRuleTypePerURI}

type HSTSRuleType string

func (s HSTSRuleType) String() string {
	return string(s)
}

const (
	HSTSRuleTypeDomain HSTSRuleType = "domain"
	HSTSRuleTypeRecord HSTSRuleType = "record"
)

var HSTSRuleTypeEnum connection.Enum[HSTSRuleType] = []HSTSRuleType{HSTSRuleTypeDomain, HSTSRuleTypeRecord}

// Domain represents a DDoSX domain
type Domain struct {
	SafeDNSZoneID *int               `json:"safedns_zone_id"`
	Name          string             `json:"name"`
	Status        DomainStatus       `json:"status"`
	DNSActive     bool               `json:"dns_active"`
	CDNActive     bool               `json:"cdn_active"`
	WAFActive     bool               `json:"waf_active"`
	ExternalDNS   *DomainExternalDNS `json:"external_dns"`
}

// DomainExternalDNS represents a DDoSX domain external DNS configuration
type DomainExternalDNS struct {
	Verified           bool   `json:"verified"`
	VerificationString string `json:"verification_string"`
	Target             string `json:"target"`
}

// DomainProperty represents a DDoSX domain property
type DomainProperty struct {
	ID    string             `json:"id"`
	Name  DomainPropertyName `json:"name"`
	Value interface{}        `json:"value"`
}

// Record represents a DDoSX record
type Record struct {
	ID              string     `json:"id"`
	DomainName      string     `json:"domain_name"`
	SafeDNSRecordID *int       `json:"safedns_record_id"`
	SSLID           *string    `json:"ssl_id"`
	Name            string     `json:"name"`
	Type            RecordType `json:"type"`
	Content         string     `json:"content"`
}

// WAF represents a DDoSX WAF configuration
type WAF struct {
	Mode          WAFMode          `json:"mode"`
	ParanoiaLevel WAFParanoiaLevel `json:"paranoia_level"`
}

// WAFRuleSet represents a DDoSX WAF rule set
type WAFRuleSet struct {
	ID     string         `json:"id"`
	Name   WAFRuleSetName `json:"name"`
	Active bool           `json:"active"`
}

// WAFRule represents a DDoSX WAF rule
type WAFRule struct {
	ID  string               `json:"id"`
	URI string               `json:"uri"`
	IP  connection.IPAddress `json:"ip"`
}

// WAFAdvancedRule represents a DDoSX WAF advanced rule
type WAFAdvancedRule struct {
	ID       string                  `json:"id"`
	Section  WAFAdvancedRuleSection  `json:"section"`
	Modifier WAFAdvancedRuleModifier `json:"modifier"`
	Phrase   string                  `json:"phrase"`
	IP       connection.IPAddress    `json:"ip"`
}

// SSL represents a DDoSX SSL
type SSL struct {
	ID           string   `json:"id"`
	UKFastSSLID  *int     `json:"ukfast_ssl_id"`
	Domains      []string `json:"domains"`
	FriendlyName string   `json:"friendly_name"`
}

// SSLContent represents a DDoSX SSL content
type SSLContent struct {
	Certificate string `json:"certificate"`
	CABundle    string `json:"ca_bundle"`
}

// SSLPrivateKey represents a DDoSX SSL private key
type SSLPrivateKey struct {
	Key string `json:"key"`
}

// ACLGeoIPRule represents a DDoSX ACL GeoIP rule
type ACLGeoIPRule struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ACLIPRule represents a DDoSX ACL IP rule
type ACLIPRule struct {
	ID   string               `json:"id"`
	IP   connection.IPAddress `json:"ip"`
	URI  string               `json:"uri"`
	Mode ACLIPMode            `json:"mode"`
}

// CDNRule represents a DDoSX CDN rule
type CDNRule struct {
	ID           string              `json:"id"`
	URI          string              `json:"uri"`
	CacheControl CDNRuleCacheControl `json:"cache_control"`
	// CacheControlDuration specifies the cache control duration. May be nil if duration not applicable
	CacheControlDuration CDNRuleCacheControlDuration `json:"cache_control_duration"`
	MimeTypes            []string                    `json:"mime_types"`
	Type                 CDNRuleType                 `json:"type"`
}

// CDNRuleCacheControlDuration represents a DDoSX CDN rule duration
type CDNRuleCacheControlDuration struct {
	Years   int `json:"years"`
	Months  int `json:"months"`
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

// Duration returns the cache control duration as time.Duration
func (d *CDNRuleCacheControlDuration) Duration() time.Duration {
	day := time.Hour * 24
	td := time.Duration(d.Years) * day * 365
	td = td + time.Duration(d.Months)*day*30
	td = td + time.Duration(d.Days)*day
	td = td + time.Duration(d.Hours)*time.Hour
	return td + time.Duration(d.Minutes)*time.Minute
}

func (d *CDNRuleCacheControlDuration) String() string {
	return durationstring.NewDuration(d.Years, d.Months, d.Days, d.Hours, d.Minutes, 0, 0, 0, 0).String()
}

// ParseCDNRuleCacheControlDuration parses string s and returns a pointer to an
// initialised CDNRuleCacheControlDuration
func ParseCDNRuleCacheControlDuration(s string) (*CDNRuleCacheControlDuration, error) {
	d, err := durationstring.Parse(s)
	if err != nil {
		return nil, err
	}

	return &CDNRuleCacheControlDuration{
		Years:   d.Years,
		Months:  d.Months,
		Days:    d.Days,
		Hours:   d.Hours,
		Minutes: d.Minutes,
	}, nil
}

// HSTSConfiguration represents HSTS configuration for a DDoSX domain
type HSTSConfiguration struct {
	Enabled bool `json:"enabled"`
}

// HSTSRule represents HSTS rule for a DDoSX domain
type HSTSRule struct {
	ID                string       `json:"id"`
	MaxAge            int          `json:"max_age"`
	Preload           bool         `json:"preload"`
	IncludeSubdomains bool         `json:"include_subdomains"`
	Type              HSTSRuleType `json:"type"`
	RecordName        *string      `json:"record_name"`
}

// WAFLog represents a WAF log entry
type WAFLog struct {
	ID        string               `json:"id"`
	Host      string               `json:"host"`
	ClientIP  connection.IPAddress `json:"client_ip"`
	Request   string               `json:"request"`
	CreatedAt connection.DateTime  `json:"created_at"`
}

// WAFLogMatch represents a WAF log match
type WAFLogMatch struct {
	ID          string               `json:"id"`
	LogID       string               `json:"log_id"`
	Host        string               `json:"host"`
	ClientIP    connection.IPAddress `json:"client_ip"`
	RequestURI  string               `json:"request_uri"`
	CreatedAt   connection.DateTime  `json:"created_at"`
	CountryCode string               `json:"country_code"`
	Method      string               `json:"method"`
	Content     string               `json:"content"`
	Message     string               `json:"message"`
	MatchData   string               `json:"match_data"`
	URIPart     string               `json:"uri_part"`
	Value       string               `json:"value"`
}
