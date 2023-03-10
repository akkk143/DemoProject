package model

import (
	"time"
)

type OriginalEvent struct {
	Outcome string `json:"outcome,omitempty"`
}

type Rule struct {
	Author         []interface{} `json:"author,omitempty"`
	CreatedAt      time.Time     `json:"created_at,omitempty"`
	CreatedBy      string        `json:"created_by,omitempty"`
	Description    string        `json:"description,omitempty"`
	Enabled        bool          `json:"enabled,omitempty"`
	FalsePositives []interface{} `json:"false_positives,omitempty"`
	From           string        `json:"from,omitempty"`
	ID             string        `json:"id,omitempty"`
	Immutable      bool          `json:"immutable,omitempty"`
	Interval       string        `json:"interval,omitempty"`
	License        string        `json:"license,omitempty"`
	MaxSignals     int           `json:"max_signals,omitempty"`
	Name           string        `json:"name,omitempty"`
	References     []interface{} `json:"references,omitempty"`
	RiskScore      int           `json:"risk_score,omitempty"`
	RuleID         string        `json:"rule_id,omitempty"`
	Severity       string        `json:"severity,omitempty"`
	Tags           []interface{} `json:"tags,omitempty"`
	To             string        `json:"to,omitempty"`
	Type           string        `json:"type,omitempty"`
	UpdatedAt      time.Time     `json:"updated_at,omitempty"`
	UpdatedBy      string        `json:"updated_by,omitempty"`
	Version        int           `json:"version,omitempty"`
}

type Signal struct {
	Depth         int           `json:"depth,omitempty"`
	OriginalEvent OriginalEvent `json:"original_event,omitempty"`
	OriginalTime  time.Time     `json:"original_time,omitempty"`
	Reason        string        `json:"reason,omitempty"`
	Rule          Rule          `json:"rule,omitempty"`
	Status        string        `json:"status,omitempty"`
}

type Execution struct {
	UUID string `json:"uuid,omitempty"`
}

type Meta struct {
	From             string `json:"from,omitempty"`
	KibanaSiemAppURL string `json:"kibana_siem_app_url,omitempty"`
}

type SeverityMapping struct {
	Severity string `json:"severity,omitempty"`
	Field    string `json:"field,omitempty"`
	Value    string `json:"value,omitempty"`
	Operator string `json:"operator,omitempty"`
}

type AlertSuppression struct {
	GroupBy []string `json:"group_by,omitempty"`
}

type Parameters struct {
	Description         string            `json:"description,omitempty"`
	RiskScore           int               `json:"risk_score,omitempty"`
	Severity            string            `json:"severity,omitempty"`
	License             string            `json:"license,omitempty"`
	Meta                Meta              `json:"meta,omitempty"`
	Author              []interface{}     `json:"author,omitempty"`
	FalsePositives      []interface{}     `json:"false_positives,omitempty"`
	From                string            `json:"from,omitempty"`
	RuleID              string            `json:"rule_id,omitempty"`
	MaxSignals          int               `json:"max_signals,omitempty"`
	RiskScoreMapping    []interface{}     `json:"risk_score_mapping,omitempty"`
	SeverityMapping     []SeverityMapping `json:"severity_mapping,omitempty"`
	Threat              []interface{}     `json:"threat,omitempty"`
	To                  string            `json:"to,omitempty"`
	References          []interface{}     `json:"references,omitempty"`
	Version             int               `json:"version,omitempty"`
	ExceptionsList      []interface{}     `json:"exceptions_list,omitempty"`
	Immutable           bool              `json:"immutable,omitempty"`
	RelatedIntegrations []interface{}     `json:"related_integrations,omitempty"`
	RequiredFields      []interface{}     `json:"required_fields,omitempty"`
	Setup               string            `json:"setup,omitempty"`
	Type                string            `json:"type,omitempty"`
	Language            string            `json:"language,omitempty"`
	Index               []string          `json:"index,omitempty"`
	Query               string            `json:"query,omitempty"`
	Filters             []interface{}     `json:"filters,omitempty"`
	AlertSuppression    AlertSuppression  `json:"alert_suppression,omitempty"`
}

type Params struct {
	Body string `json:"body,omitempty"`
}

type Action struct {
	Group        string `json:"group,omitempty"`
	ID           string `json:"id,omitempty"`
	Params       Params `json:"params,omitempty"`
	ActionTypeID string `json:"action_type_id,omitempty"`
}

type AlertRule struct {
	Category         string            `json:"category,omitempty"`
	Consumer         string            `json:"consumer,omitempty"`
	Execution        Execution         `json:"execution,omitempty"`
	Name             string            `json:"name,omitempty"`
	Producer         string            `json:"producer,omitempty"`
	RuleTypeID       string            `json:"rule_type_id,omitempty"`
	UUID             string            `json:"uuid,omitempty"`
	Tags             []interface{}     `json:"tags,omitempty"`
	Parameters       Parameters        `json:"parameters,omitempty"`
	Actions          []Action          `json:"actions,omitempty"`
	Author           []interface{}     `json:"author,omitempty"`
	CreatedAt        time.Time         `json:"created_at,omitempty"`
	CreatedBy        string            `json:"created_by,omitempty"`
	Description      string            `json:"description,omitempty"`
	Enabled          bool              `json:"enabled,omitempty"`
	ExceptionsList   []interface{}     `json:"exceptions_list,omitempty"`
	FalsePositives   []interface{}     `json:"false_positives,omitempty"`
	From             string            `json:"from,omitempty"`
	Immutable        bool              `json:"immutable,omitempty"`
	Interval         string            `json:"interval,omitempty"`
	Indices          []string          `json:"indices,omitempty"`
	License          string            `json:"license,omitempty"`
	MaxSignals       int               `json:"max_signals,omitempty"`
	References       []interface{}     `json:"references,omitempty"`
	RiskScoreMapping []interface{}     `json:"risk_score_mapping,omitempty"`
	RuleID           string            `json:"rule_id,omitempty"`
	SeverityMapping  []SeverityMapping `json:"severity_mapping,omitempty"`
	Threat           []interface{}     `json:"threat,omitempty"`
	To               string            `json:"to,omitempty"`
	Type             string            `json:"type,omitempty"`
	UpdatedAt        time.Time         `json:"updated_at,omitempty"`
	UpdatedBy        string            `json:"updated_by,omitempty"`
	Version          int               `json:"version,omitempty"`
	Meta             Meta              `json:"meta,omitempty"`
	RiskScore        int               `json:"risk_score,omitempty"`
	Severity         string            `json:"severity,omitempty"`
}

type Ancestor struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Index string `json:"index,omitempty"`
	Depth int    `json:"depth,omitempty"`
}

type AlertOriginalEvent struct {
	AgentIDStatus string    `json:"agent_id_status,omitempty"`
	Ingested      time.Time `json:"ingested,omitempty"`
	Outcome       string    `json:"outcome,omitempty"`
}

type Term struct {
	Field string      `json:"field,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Suppression struct {
	Terms     []Term    `json:"terms,omitempty"`
	Start     time.Time `json:"start,omitempty"`
	End       time.Time `json:"end,omitempty"`
	DocsCount int       `json:"docs_count,omitempty"`
}
type Alert struct {
	Rule           AlertRule          `json:"rule,omitempty"`
	OriginalTime   time.Time          `json:"original_time,omitempty"`
	Ancestors      []Ancestor         `json:"ancestors,omitempty"`
	Status         string             `json:"status,omitempty"`
	WorkflowStatus string             `json:"workflow_status,omitempty"`
	Depth          int                `json:"depth,omitempty"`
	Reason         string             `json:"reason,omitempty"`
	Severity       string             `json:"severity,omitempty"`
	RiskScore      int                `json:"risk_score,omitempty"`
	OriginalEvent  AlertOriginalEvent `json:"original_event,omitempty"`
	Suppression    Suppression        `json:"suppression,omitempty"`
	UUID           string             `json:"uuid,omitempty"`
}

type Kibana struct {
	Version  string   `json:"version,omitempty"`
	Alert    Alert    `json:"alert,omitempty"`
	SpaceIds []string `json:"space_ids,omitempty"`
}

type Parent struct {
	ID string `json:"id,omitempty"`
}

type Agent struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type Process struct {
	Args   []string `json:"args,omitempty"`
	Parent Parent   `json:"parent,omitempty"`
	Pid    int      `json:"pid,omitempty"`
	Title  string   `json:"title,omitempty"`
}

type DataStream struct {
	Namespace string `json:"namespace,omitempty"`
	Type      string `json:"type,omitempty"`
	Dataset   string `json:"dataset,omitempty"`
}

type Source struct {
	IP string `json:"ip,omitempty"`
}

type Processor struct {
	Name  string `json:"name,omitempty"`
	Event string `json:"event,omitempty"`
}

type Url struct {
	Path   string `json:"path,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	Port   int    `json:"port,omitempty"`
	Domain string `json:"domain,omitempty"`
	Full   string `json:"full,omitempty"`
}

type Observer struct {
	Hostname string `json:"hostname,omitempty"`
	Type     string `json:"type,omitempty"`
	Version  string `json:"version,omitempty"`
}

type Trace struct {
	ID string `json:"id,omitempty"`
}

type Node struct {
	Name string `json:"name,omitempty"`
}

type Runtime struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type Language struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type Service struct {
	Node     Node     `json:"node,omitempty"`
	Name     string   `json:"name,omitempty"`
	Runtime  Runtime  `json:"runtime,omitempty"`
	Language Language `json:"language,omitempty"`
}

type HostOs struct {
	Platform string `json:"platform,omitempty"`
}

type Host struct {
	Hostname     string   `json:"hostname,omitempty"`
	Os           HostOs   `json:"os,omitempty"`
	IP           []string `json:"ip,omitempty"`
	Name         string   `json:"name,omitempty"`
	Architecture string   `json:"architecture,omitempty"`
}

type Client struct {
	IP string `json:"ip,omitempty"`
}

type RequestHeaders struct {
	Origin          []string `json:"Origin,omitempty"`
	Locale          []string `json:"Locale,omitempty"`
	SecChUa         []string `json:"Sec-Ch-Ua,omitempty"`
	Email           []string `json:"Email,omitempty"`
	Accept          []string `json:"Accept,omitempty"`
	Requestid       []string `json:"Requestid,omitempty"`
	Authority       []string `json:"Authority,omitempty"`
	SecChUaPlatform []string `json:"Sec-Ch-Ua-Platform,omitempty"`
	Referer         []string `json:"Referer,omitempty"`
	UserAgent       []string `json:"User-Agent,omitempty"`
	TapAPIToken     []string `json:"Tap-Api-Token,omitempty"`
	Connection      []string `json:"Connection,omitempty"`
	SecFetchSite    []string `json:"Sec-Fetch-Site,omitempty"`
	SecFetchDest    []string `json:"Sec-Fetch-Dest,omitempty"`
	AcceptEncoding  []string `json:"Accept-Encoding,omitempty"`
	Traceparent     []string `json:"Traceparent,omitempty"`
	SecFetchMode    []string `json:"Sec-Fetch-Mode,omitempty"`
	PostmanToken    []string `json:"Postman-Token,omitempty"`
	AcceptLanguage  []string `json:"Accept-Language,omitempty"`
	SecChUaMobile   []string `json:"Sec-Ch-Ua-Mobile,omitempty"`
	ContentLength   []string `json:"Content-Length,omitempty"`
	ContentType     []string `json:"Content-Type,omitempty"`
}

type Request struct {
	Headers RequestHeaders `json:"headers,omitempty"`
	Method  string         `json:"method,omitempty"`
}

type ResponseHeaders struct {
	AccessControlAllowOrigin      []string `json:"Access-Control-Allow-Origin,omitempty"`
	AccessControlAllowCredentials []string `json:"Access-Control-Allow-Credentials,omitempty"`
	ContentEncoding               []string `json:"Content-Encoding,omitempty"`
	Traceid                       []string `json:"Traceid,omitempty"`
	ContentType                   []string `json:"Content-Type,omitempty"`
}

type Response struct {
	Headers    ResponseHeaders `json:"headers,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}

type HTTP struct {
	Request  Request  `json:"request,omitempty"`
	Response Response `json:"response,omitempty"`
	Version  string   `json:"version,omitempty"`
}

type Duration struct {
	Us int `json:"us,omitempty"`
}

type SpanCount struct {
	Dropped int `json:"dropped,omitempty"`
	Started int `json:"started,omitempty"`
}

type Transaction struct {
	Result              string    `json:"result,omitempty"`
	Duration            Duration  `json:"duration,omitempty"`
	RepresentativeCount int       `json:"representative_count,omitempty"`
	Name                string    `json:"name,omitempty"`
	ID                  string    `json:"id,omitempty"`
	SpanCount           SpanCount `json:"span_count,omitempty"`
	Type                string    `json:"type,omitempty"`
	Sampled             bool      `json:"sampled,omitempty"`
}

type Os struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Full    string `json:"full,omitempty"`
}
type UserAgent struct {
	Original string `json:"original,omitempty"`
	Os       Os     `json:"os,omitempty"`
	Name     string `json:"name,omitempty"`
	Device   struct {
		Name string `json:"name,omitempty"`
	} `json:"device,omitempty"`
	Version string `json:"version,omitempty"`
}

type Timestamp struct {
	Us int64 `json:"us,omitempty"`
}

type Ecs struct {
	Version string `json:"version,omitempty"`
}

type Event struct {
	AgentIDStatus string    `json:"agent_id_status,omitempty"`
	Ingested      time.Time `json:"ingested,omitempty"`
	Outcome       string    `json:"outcome,omitempty"`
	Kind          string    `json:"kind,omitempty"`
}

type AlertData struct {
	Signal      Signal      `json:"signal,omitempty"`
	ID          string      `json:"_id,omitempty"`
	Index       string      `json:"_index,omitempty"`
	Kibana      Kibana      `json:"kibana,omitempty"`
	Timestamp   time.Time   `json:"@timestamp,omitempty"`
	Parent      Parent      `json:"parent,omitempty"`
	Agent       Agent       `json:"agent,omitempty"`
	Process     Process     `json:"process,omitempty"`
	DataStream  DataStream  `json:"data_stream,omitempty"`
	Source      Source      `json:"source,omitempty"`
	Processor   Processor   `json:"processor,omitempty"`
	URL         Url         `json:"url,omitempty"`
	Observer    Observer    `json:"observer,omitempty"`
	Trace       Trace       `json:"trace,omitempty"`
	Service     Service     `json:"service,omitempty"`
	Host        Host        `json:"host,omitempty"`
	Client      Client      `json:"client,omitempty"`
	HTTP        HTTP        `json:"http,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	UserAgent   UserAgent   `json:"user_agent,omitempty"`
	Timestamp0  Timestamp   `json:"timestamp,omitempty"`
	Ecs         Ecs         `json:"ecs,omitempty"`
	Event       Event       `json:"event,omitempty"`
}
