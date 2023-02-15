package model

import "time"

type Agent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Client struct {
	Ip string `json:"ip"`
}

type DataStream struct {
	Dataset   string `json:"dataset"`
	Namespace string `json:"namespace"`
	Type      string `json:"type"`
}

type Ecs struct {
	Version string `json:"version"`
}

type Event struct {
	AgentIdStatus string    `json:"agent_id_status"`
	Ingested      time.Time `json:"ingested"`
	Kind          string    `json:"kind"`
	Outcome       string    `json:"outcome"`
}

type Os struct {
	Platform string `json:"platform"`
}

type Host struct {
	Architecture string   `json:"architecture"`
	Hostname     string   `json:"hostname"`
	Ip           []string `json:"ip"`
	Name         string   `json:"name"`
	Os           Os       `json:"os"`
}

type RequestHeaders struct {
	Accept          []string `json:"Accept"`
	AcceptEncoding  []string `json:"Accept-Encoding"`
	AcceptLanguage  []string `json:"Accept-Language"`
	Authority       []string `json:"Authority"`
	Connection      []string `json:"Connection"`
	ContentLength   []string `json:"Content-Length"`
	ContentType     []string `json:"Content-Type"`
	Email           []string `json:"Email"`
	Locale          []string `json:"Locale"`
	Origin          []string `json:"Origin"`
	PostmanToken    []string `json:"Postman-Token"`
	Referer         []string `json:"Referer"`
	RequestId       []string `json:"Requestid"`
	SecChUa         []string `json:"Sec-Ch-Ua"`
	SecChUaMobile   []string `json:"Sec-Ch-Ua-Mobile"`
	SecChUaPlatform []string `json:"Sec-Ch-Ua-Platform"`
	SecFetchDest    []string `json:"Sec-Fetch-Dest"`
	SecFetchMode    []string `json:"Sec-Fetch-Mode"`
	SecFetchSite    []string `json:"Sec-Fetch-Site"`
	TapApiToken     []string `json:"Tap-Api-Token"`
	TraceParent     []string `json:"Traceparent"`
	UserAgent       []string `json:"User-Agent"`
}

type Request struct {
	Headers RequestHeaders `json:"headers"`
	Method  string         `json:"method"`
}

type ResponseHeaders struct {
	AccessControlAllowCredentials []string `json:"Access-Control-Allow-Credentials"`
	AccessControlAllowOrigin      []string `json:"Access-Control-Allow-Origin"`
	ContentEncoding               []string `json:"Content-Encoding"`
	ContentType                   []string `json:"Content-Type"`
	TraceId                       []string `json:"Traceid"`
}
type Response struct {
	Headers    ResponseHeaders `json:"headers"`
	StatusCode int             `json:"status_code"`
}
type Http struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	Version  string   `json:"version"`
}

type Ancestors struct {
	Depth int    `json:"depth"`
	Id    string `json:"id"`
	Index string `json:"index"`
	Type  string `json:"type"`
}

type OriginalEvent struct {
	AgentIdStatus string    `json:"agent_id_status"`
	Ingested      time.Time `json:"ingested"`
	Outcome       string    `json:"outcome"`
}

type Params struct {
	Body string `json:"body"`
}

type Actions struct {
	ActionTypeId string `json:"action_type_id"`
	Group        string `json:"group"`
	Id           string `json:"id"`
	Params       Params `json:"params"`
}
type Execution struct {
	Uuid string `json:"uuid"`
}

type Meta struct {
	From             string `json:"from"`
	KibanaSiemAppUrl string `json:"kibana_siem_app_url"`
}

type AlertSuppression struct {
	GroupBy []string `json:"group_by"`
}

type SeverityMapping struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Severity string `json:"severity"`
	Value    string `json:"value"`
}

type Parameters struct {
	AlertSuppression    AlertSuppression  `json:"alert_suppression"`
	Author              []interface{}     `json:"author"`
	Description         string            `json:"description"`
	ExceptionsList      []interface{}     `json:"exceptions_list"`
	FalsePositives      []interface{}     `json:"false_positives"`
	Filters             []interface{}     `json:"filters"`
	From                string            `json:"from"`
	Immutable           bool              `json:"immutable"`
	Index               []string          `json:"index"`
	Language            string            `json:"language"`
	License             string            `json:"license"`
	MaxSignals          int               `json:"max_signals"`
	Meta                Meta              `json:"meta"`
	Query               string            `json:"query"`
	References          []interface{}     `json:"references"`
	RelatedIntegrations []interface{}     `json:"related_integrations"`
	RequiredFields      []interface{}     `json:"required_fields"`
	RiskScore           int               `json:"risk_score"`
	RiskScoreMapping    []interface{}     `json:"risk_score_mapping"`
	RuleId              string            `json:"rule_id"`
	Setup               string            `json:"setup"`
	Severity            string            `json:"severity"`
	SeverityMapping     []SeverityMapping `json:"severity_mapping"`
	Threat              []interface{}     `json:"threat"`
	To                  string            `json:"to"`
	Type                string            `json:"type"`
	Version             int               `json:"version"`
}
type SignalRule struct {
	Actions          []Actions         `json:"actions"`
	Author           []interface{}     `json:"author"`
	Category         string            `json:"category"`
	Consumer         string            `json:"consumer"`
	CreatedAt        time.Time         `json:"created_at"`
	CreatedBy        string            `json:"created_by"`
	Description      string            `json:"description"`
	Enabled          bool              `json:"enabled"`
	ExceptionsList   []interface{}     `json:"exceptions_list"`
	Execution        Execution         `json:"execution"`
	FalsePositives   []interface{}     `json:"false_positives"`
	From             string            `json:"from"`
	Immutable        bool              `json:"immutable"`
	Indices          []string          `json:"indices"`
	Interval         string            `json:"interval"`
	License          string            `json:"license"`
	MaxSignals       int               `json:"max_signals"`
	Meta             Meta              `json:"meta"`
	Name             string            `json:"name"`
	Parameters       Parameters        `json:"parameters"`
	Producer         string            `json:"producer"`
	References       []interface{}     `json:"references"`
	RiskScore        int               `json:"risk_score"`
	RiskScoreMapping []interface{}     `json:"risk_score_mapping"`
	RuleId           string            `json:"rule_id"`
	RuleTypeId       string            `json:"rule_type_id"`
	Severity         string            `json:"severity"`
	SeverityMapping  []SeverityMapping `json:"severity_mapping"`
	Tags             []interface{}     `json:"tags"`
	Threat           []interface{}     `json:"threat"`
	To               string            `json:"to"`
	Type             string            `json:"type"`
	UpdatedAt        time.Time         `json:"updated_at"`
	UpdatedBy        string            `json:"updated_by"`
	Uuid             string            `json:"uuid"`
	Version          int               `json:"version"`
}

type Rule struct {
	Author         []interface{} `json:"author"`
	CreatedAt      time.Time     `json:"created_at"`
	CreatedBy      string        `json:"created_by"`
	Description    string        `json:"description"`
	Enabled        bool          `json:"enabled"`
	FalsePositives []interface{} `json:"false_positives"`
	From           string        `json:"from"`
	Id             string        `json:"id"`
	Immutable      bool          `json:"immutable"`
	Interval       string        `json:"interval"`
	License        string        `json:"license"`
	MaxSignals     int           `json:"max_signals"`
	Name           string        `json:"name"`
	References     []interface{} `json:"references"`
	RiskScore      int           `json:"risk_score"`
	RuleId         string        `json:"rule_id"`
	Severity       string        `json:"severity"`
	Tags           []interface{} `json:"tags"`
	To             string        `json:"to"`
	Type           string        `json:"type"`
	UpdatedAt      time.Time     `json:"updated_at"`
	UpdatedBy      string        `json:"updated_by"`
	Version        int           `json:"version"`
}

type Terms struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type Suppression struct {
	DocsCount int       `json:"docs_count"`
	End       time.Time `json:"end"`
	Start     time.Time `json:"start"`
	Terms     []Terms   `json:"terms"`
}

type Alert struct {
	Ancestors      []Ancestors   `json:"ancestors"`
	Depth          int           `json:"depth"`
	OriginalEvent  OriginalEvent `json:"original_event"`
	OriginalTime   time.Time     `json:"original_time"`
	Reason         string        `json:"reason"`
	RiskScore      int           `json:"risk_score"`
	Rule           Rule          `json:"rule"`
	Severity       string        `json:"severity"`
	Status         string        `json:"status"`
	Suppression    Suppression   `json:"suppression"`
	Uuid           string        `json:"uuid"`
	WorkflowStatus string        `json:"workflow_status"`
}

type Kibana struct {
	Alert    Alert    `json:"alert"`
	SpaceIds []string `json:"space_ids"`
	Version  string   `json:"version"`
}

type Observer struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Version  string `json:"version"`
}

type Parent struct {
	Id string `json:"id"`
}

type Process struct {
	Args   []string `json:"args"`
	Parent Parent   `json:"parent"`
	Pid    int      `json:"pid"`
	Title  string   `json:"title"`
}

type Processor struct {
	Event string `json:"event"`
	Name  string `json:"name"`
}

type Language struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Node struct {
	Name string `json:"name"`
}

type Runtime struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Service struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
	Node     Node     `json:"node"`
	Runtime  Runtime  `json:"runtime"`
}

type SignalOriginalEvent struct {
	Outcome string `json:"outcome"`
}

type Signal struct {
	Depth         int                 `json:"depth"`
	OriginalEvent SignalOriginalEvent `json:"original_event"`
	OriginalTime  time.Time           `json:"original_time"`
	Reason        string              `json:"reason"`
	Rule          SignalRule          `json:"rule"`
	Status        string              `json:"status"`
}

type Source struct {
	Ip string `json:"ip"`
}

type TimeStamp struct {
	Us int64 `json:"us"`
}

type Trace struct {
	Id string `json:"id"`
}

type Duration struct {
	Us int `json:"us"`
}

type SpanCount struct {
	Dropped int `json:"dropped"`
	Started int `json:"started"`
}

type Transaction struct {
	Duration            Duration  `json:"duration"`
	Id                  string    `json:"id"`
	Name                string    `json:"name"`
	RepresentativeCount int       `json:"representative_count"`
	Result              string    `json:"result"`
	Sampled             bool      `json:"sampled"`
	SpanCount           SpanCount `json:"span_count"`
	Type                string    `json:"type"`
}

type Url struct {
	Domain string `json:"domain"`
	Full   string `json:"full"`
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type Device struct {
	Name string `json:"name"`
}

type UserAgentOs struct {
	Full    string `json:"full"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type UserAgent struct {
	Device   Device      `json:"device"`
	Name     string      `json:"name"`
	Original string      `json:"original"`
	Os       UserAgentOs `json:"os"`
	Version  string      `json:"version"`
}

type AlertDatas struct {
	Timestamp   time.Time   `json:"@timestamp"`
	Id          string      `json:"_id"`
	Index       string      `json:"_index"`
	Agent       Agent       `json:"agent"`
	Client      Client      `json:"client"`
	DataStream  DataStream  `json:"data_stream"`
	Ecs         Ecs         `json:"ecs"`
	Event       Event       `json:"event"`
	Host        Host        `json:"host"`
	Http        Http        `json:"http"`
	Kibana      Kibana      `json:"kibana"`
	Observer    Observer    `json:"observer"`
	Parent      Parent      `json:"parent"`
	Process     Process     `json:"process"`
	Processor   Processor   `json:"processor"`
	Service     Service     `json:"service"`
	Signal      Signal      `json:"signal"`
	Source      Source      `json:"source"`
	Timestamp1  TimeStamp   `json:"timestamp"`
	Trace       Trace       `json:"trace"`
	Transaction Transaction `json:"transaction"`
	Url         Url         `json:"url"`
	UserAgent   UserAgent   `json:"user_agent"`
}

type AlertData struct {
	Signal struct {
		Depth         int `json:"depth"`
		OriginalEvent struct {
			Outcome string `json:"outcome"`
		} `json:"original_event"`
		OriginalTime time.Time `json:"original_time"`
		Reason       string    `json:"reason"`
		Rule         struct {
			Author         []interface{} `json:"author"`
			CreatedAt      time.Time     `json:"created_at"`
			CreatedBy      string        `json:"created_by"`
			Description    string        `json:"description"`
			Enabled        bool          `json:"enabled"`
			FalsePositives []interface{} `json:"false_positives"`
			From           string        `json:"from"`
			Id             string        `json:"id"`
			Immutable      bool          `json:"immutable"`
			Interval       string        `json:"interval"`
			License        string        `json:"license"`
			MaxSignals     int           `json:"max_signals"`
			Name           string        `json:"name"`
			References     []interface{} `json:"references"`
			RiskScore      int           `json:"risk_score"`
			RuleId         string        `json:"rule_id"`
			Severity       string        `json:"severity"`
			Tags           []interface{} `json:"tags"`
			To             string        `json:"to"`
			Type           string        `json:"type"`
			UpdatedAt      time.Time     `json:"updated_at"`
			UpdatedBy      string        `json:"updated_by"`
			Version        int           `json:"version"`
		} `json:"rule"`
		Status string `json:"status"`
	} `json:"signal"`
	Id     string `json:"_id"`
	Index  string `json:"_index"`
	Kibana struct {
		Version string `json:"version"`
		Alert   struct {
			Rule struct {
				Category  string `json:"category"`
				Consumer  string `json:"consumer"`
				Execution struct {
					Uuid string `json:"uuid"`
				} `json:"execution"`
				Name       string        `json:"name"`
				Producer   string        `json:"producer"`
				RuleTypeId string        `json:"rule_type_id"`
				Uuid       string        `json:"uuid"`
				Tags       []interface{} `json:"tags"`
				Parameters struct {
					Description string `json:"description"`
					RiskScore   int    `json:"risk_score"`
					Severity    string `json:"severity"`
					License     string `json:"license"`
					Meta        struct {
						From             string `json:"from"`
						KibanaSiemAppUrl string `json:"kibana_siem_app_url"`
					} `json:"meta"`
					Author           []interface{} `json:"author"`
					FalsePositives   []interface{} `json:"false_positives"`
					From             string        `json:"from"`
					RuleId           string        `json:"rule_id"`
					MaxSignals       int           `json:"max_signals"`
					RiskScoreMapping []interface{} `json:"risk_score_mapping"`
					SeverityMapping  []struct {
						Severity string `json:"severity"`
						Field    string `json:"field"`
						Value    string `json:"value"`
						Operator string `json:"operator"`
					} `json:"severity_mapping"`
					Threat              []interface{} `json:"threat"`
					To                  string        `json:"to"`
					References          []interface{} `json:"references"`
					Version             int           `json:"version"`
					ExceptionsList      []interface{} `json:"exceptions_list"`
					Immutable           bool          `json:"immutable"`
					RelatedIntegrations []interface{} `json:"related_integrations"`
					RequiredFields      []interface{} `json:"required_fields"`
					Setup               string        `json:"setup"`
					Type                string        `json:"type"`
					Language            string        `json:"language"`
					Index               []string      `json:"index"`
					Query               string        `json:"query"`
					Filters             []interface{} `json:"filters"`
					AlertSuppression    struct {
						GroupBy []string `json:"group_by"`
					} `json:"alert_suppression"`
				} `json:"parameters"`
				Actions []struct {
					Group  string `json:"group"`
					Id     string `json:"id"`
					Params struct {
						Body string `json:"body"`
					} `json:"params"`
					ActionTypeId string `json:"action_type_id"`
				} `json:"actions"`
				Author           []interface{} `json:"author"`
				CreatedAt        time.Time     `json:"created_at"`
				CreatedBy        string        `json:"created_by"`
				Description      string        `json:"description"`
				Enabled          bool          `json:"enabled"`
				ExceptionsList   []interface{} `json:"exceptions_list"`
				FalsePositives   []interface{} `json:"false_positives"`
				From             string        `json:"from"`
				Immutable        bool          `json:"immutable"`
				Interval         string        `json:"interval"`
				Indices          []string      `json:"indices"`
				License          string        `json:"license"`
				MaxSignals       int           `json:"max_signals"`
				References       []interface{} `json:"references"`
				RiskScoreMapping []interface{} `json:"risk_score_mapping"`
				RuleId           string        `json:"rule_id"`
				SeverityMapping  []struct {
					Severity string `json:"severity"`
					Field    string `json:"field"`
					Value    string `json:"value"`
					Operator string `json:"operator"`
				} `json:"severity_mapping"`
				Threat    []interface{} `json:"threat"`
				To        string        `json:"to"`
				Type      string        `json:"type"`
				UpdatedAt time.Time     `json:"updated_at"`
				UpdatedBy string        `json:"updated_by"`
				Version   int           `json:"version"`
				Meta      struct {
					From             string `json:"from"`
					KibanaSiemAppUrl string `json:"kibana_siem_app_url"`
				} `json:"meta"`
				RiskScore int    `json:"risk_score"`
				Severity  string `json:"severity"`
			} `json:"rule"`
			OriginalTime time.Time `json:"original_time"`
			Ancestors    []struct {
				Id    string `json:"id"`
				Type  string `json:"type"`
				Index string `json:"index"`
				Depth int    `json:"depth"`
			} `json:"ancestors"`
			Status         string `json:"status"`
			WorkflowStatus string `json:"workflow_status"`
			Depth          int    `json:"depth"`
			Reason         string `json:"reason"`
			Severity       string `json:"severity"`
			RiskScore      int    `json:"risk_score"`
			OriginalEvent  struct {
				AgentIdStatus string    `json:"agent_id_status"`
				Ingested      time.Time `json:"ingested"`
				Outcome       string    `json:"outcome"`
			} `json:"original_event"`
			Suppression struct {
				Terms []struct {
					Field string      `json:"field"`
					Value interface{} `json:"value"`
				} `json:"terms"`
				Start     time.Time `json:"start"`
				End       time.Time `json:"end"`
				DocsCount int       `json:"docs_count"`
			} `json:"suppression"`
			Uuid string `json:"uuid"`
		} `json:"alert"`
		SpaceIds []string `json:"space_ids"`
	} `json:"kibana"`
	Timestamp string `json:"timestamp"`
	Parent    struct {
		Id string `json:"id"`
	} `json:"parent"`
	Agent struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"agent"`
	Process struct {
		Args   []string `json:"args"`
		Parent struct {
			Pid int `json:"pid"`
		} `json:"parent"`
		Pid   int    `json:"pid"`
		Title string `json:"title"`
	} `json:"process"`
	DataStream struct {
		Namespace string `json:"namespace"`
		Type      string `json:"type"`
		Dataset   string `json:"dataset"`
	} `json:"data_stream"`
	Source struct {
		Ip string `json:"ip"`
	} `json:"source"`
	Processor struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	} `json:"processor"`
	Url struct {
		Path   string `json:"path"`
		Scheme string `json:"scheme"`
		Port   int    `json:"port"`
		Domain string `json:"domain"`
		Full   string `json:"full"`
	} `json:"url"`
	Observer struct {
		Hostname string `json:"hostname"`
		Type     string `json:"type"`
		Version  string `json:"version"`
	} `json:"observer"`
	Trace struct {
		Id string `json:"id"`
	} `json:"trace"`
	Service struct {
		Node struct {
			Name string `json:"name"`
		} `json:"node"`
		Name    string `json:"name"`
		Runtime struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"runtime"`
		Language struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"language"`
	} `json:"service"`
	Host struct {
		Hostname string `json:"hostname"`
		Os       struct {
			Platform string `json:"platform"`
		} `json:"os"`
		Ip           []string `json:"ip"`
		Name         string   `json:"name"`
		Architecture string   `json:"architecture"`
	} `json:"host"`
	Http struct {
		Request struct {
			Headers struct {
				Origin          []string `json:"Origin"`
				Locale          []string `json:"Locale"`
				SecChUa         []string `json:"Sec-Ch-Ua"`
				Email           []string `json:"Email"`
				Requestid       []string `json:"Requestid"`
				Accept          []string `json:"Accept"`
				Authority       []string `json:"Authority"`
				SecChUaPlatform []string `json:"Sec-Ch-Ua-Platform"`
				UserAgent       []string `json:"User-Agent"`
				Connection      []string `json:"Connection"`
				Referer         []string `json:"Referer"`
				TapApiToken     []string `json:"Tap-Api-Token"`
				SecFetchSite    []string `json:"Sec-Fetch-Site"`
				SecFetchDest    []string `json:"Sec-Fetch-Dest"`
				AcceptEncoding  []string `json:"Accept-Encoding"`
				Traceparent     []string `json:"Traceparent"`
				SecFetchMode    []string `json:"Sec-Fetch-Mode"`
				PostmanToken    []string `json:"Postman-Token"`
				AcceptLanguage  []string `json:"Accept-Language"`
				ContentLength   []string `json:"Content-Length"`
				SecChUaMobile   []string `json:"Sec-Ch-Ua-Mobile"`
				ContentType     []string `json:"Content-Type"`
			} `json:"headers"`
			Method string `json:"method"`
		} `json:"request"`
		Response struct {
			Headers struct {
				AccessControlAllowOrigin      []string `json:"Access-Control-Allow-Origin"`
				AccessControlAllowCredentials []string `json:"Access-Control-Allow-Credentials"`
				ContentEncoding               []string `json:"Content-Encoding"`
				Traceid                       []string `json:"Traceid"`
				ContentType                   []string `json:"Content-Type"`
			} `json:"headers"`
			StatusCode int `json:"status_code"`
		} `json:"response"`
		Version string `json:"version"`
	} `json:"http"`
	Client struct {
		Ip string `json:"ip"`
	} `json:"client"`
	Transaction struct {
		Result   string `json:"result"`
		Duration struct {
			Us int `json:"us"`
		} `json:"duration"`
		RepresentativeCount int    `json:"representative_count"`
		Name                string `json:"name"`
		Id                  string `json:"id"`
		SpanCount           struct {
			Dropped int `json:"dropped"`
			Started int `json:"started"`
		} `json:"span_count"`
		Type    string `json:"type"`
		Sampled bool   `json:"sampled"`
	} `json:"transaction"`
	UserAgent struct {
		Original string `json:"original"`
		Os       struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Full    string `json:"full"`
		} `json:"os"`
		Name   string `json:"name"`
		Device struct {
			Name string `json:"name"`
		} `json:"device"`
		Version string `json:"version"`
	} `json:"user_agent"`
	Timestamp1 struct {
		Us int64 `json:"us"`
	} `json:"timestamp"`
	Ecs struct {
		Version string `json:"version"`
	} `json:"ecs"`
	Event struct {
		AgentIdStatus string    `json:"agent_id_status"`
		Ingested      time.Time `json:"ingested"`
		Outcome       string    `json:"outcome"`
		Kind          string    `json:"kind"`
	} `json:"event"`
}

type Alerts struct {
	Timestamp string `json:"timestamp"`
	Id        string `json:"_id"`
	Index     string `json:"_index"`
	Agent     struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"agent"`
	Client struct {
		Ip string `json:"ip"`
	} `json:"client"`
	DataStream struct {
		Dataset   string `json:"dataset"`
		Namespace string `json:"namespace"`
		Type      string `json:"type"`
	} `json:"data_stream"`
	Ecs struct {
		Version string `json:"version"`
	} `json:"ecs"`
	Event struct {
		AgentIdStatus string    `json:"agent_id_status"`
		Ingested      time.Time `json:"ingested"`
		Kind          string    `json:"kind"`
		Outcome       string    `json:"outcome"`
	} `json:"event"`
	Host struct {
		Architecture string   `json:"architecture"`
		Hostname     string   `json:"hostname"`
		Ip           []string `json:"ip"`
		Name         string   `json:"name"`
		Os           struct {
			Platform string `json:"platform"`
		} `json:"os"`
	} `json:"host"`
	Http struct {
		Request struct {
			Headers struct {
				Accept          []string `json:"Accept"`
				AcceptEncoding  []string `json:"Accept-Encoding"`
				AcceptLanguage  []string `json:"Accept-Language"`
				Authority       []string `json:"Authority"`
				Connection      []string `json:"Connection"`
				ContentLength   []string `json:"Content-Length"`
				ContentType     []string `json:"Content-Type"`
				Email           []string `json:"Email"`
				Locale          []string `json:"Locale"`
				Origin          []string `json:"Origin"`
				PostmanToken    []string `json:"Postman-Token"`
				Referer         []string `json:"Referer"`
				Requestid       []string `json:"Requestid"`
				SecChUa         []string `json:"Sec-Ch-Ua"`
				SecChUaMobile   []string `json:"Sec-Ch-Ua-Mobile"`
				SecChUaPlatform []string `json:"Sec-Ch-Ua-Platform"`
				SecFetchDest    []string `json:"Sec-Fetch-Dest"`
				SecFetchMode    []string `json:"Sec-Fetch-Mode"`
				SecFetchSite    []string `json:"Sec-Fetch-Site"`
				TapApiToken     []string `json:"Tap-Api-Token"`
				Traceparent     []string `json:"Traceparent"`
				UserAgent       []string `json:"User-Agent"`
			} `json:"headers"`
			Method string `json:"method"`
		} `json:"request"`
		Response struct {
			Headers struct {
				AccessControlAllowCredentials []string `json:"Access-Control-Allow-Credentials"`
				AccessControlAllowOrigin      []string `json:"Access-Control-Allow-Origin"`
				ContentEncoding               []string `json:"Content-Encoding"`
				ContentType                   []string `json:"Content-Type"`
				Traceid                       []string `json:"Traceid"`
			} `json:"headers"`
			StatusCode int `json:"status_code"`
		} `json:"response"`
		Version string `json:"version"`
	} `json:"http"`
	Kibana struct {
		Alert struct {
			Ancestors []struct {
				Depth int    `json:"depth"`
				Id    string `json:"id"`
				Index string `json:"index"`
				Type  string `json:"type"`
			} `json:"ancestors"`
			Depth         int `json:"depth"`
			OriginalEvent struct {
				AgentIdStatus string    `json:"agent_id_status"`
				Ingested      time.Time `json:"ingested"`
				Outcome       string    `json:"outcome"`
			} `json:"original_event"`
			OriginalTime time.Time `json:"original_time"`
			Reason       string    `json:"reason"`
			RiskScore    int       `json:"risk_score"`
			Rule         struct {
				Actions []struct {
					ActionTypeId string `json:"action_type_id"`
					Group        string `json:"group"`
					Id           string `json:"id"`
					Params       struct {
						Body string `json:"body"`
					} `json:"params"`
				} `json:"actions"`
				Author         []interface{} `json:"author"`
				Category       string        `json:"category"`
				Consumer       string        `json:"consumer"`
				CreatedAt      time.Time     `json:"created_at"`
				CreatedBy      string        `json:"created_by"`
				Description    string        `json:"description"`
				Enabled        bool          `json:"enabled"`
				ExceptionsList []interface{} `json:"exceptions_list"`
				Execution      struct {
					Uuid string `json:"uuid"`
				} `json:"execution"`
				FalsePositives []interface{} `json:"false_positives"`
				From           string        `json:"from"`
				Immutable      bool          `json:"immutable"`
				Indices        []string      `json:"indices"`
				Interval       string        `json:"interval"`
				License        string        `json:"license"`
				MaxSignals     int           `json:"max_signals"`
				Meta           struct {
					From             string `json:"from"`
					KibanaSiemAppUrl string `json:"kibana_siem_app_url"`
				} `json:"meta"`
				Name       string `json:"name"`
				Parameters struct {
					AlertSuppression struct {
						GroupBy []string `json:"group_by"`
					} `json:"alert_suppression"`
					Author         []interface{} `json:"author"`
					Description    string        `json:"description"`
					ExceptionsList []interface{} `json:"exceptions_list"`
					FalsePositives []interface{} `json:"false_positives"`
					Filters        []interface{} `json:"filters"`
					From           string        `json:"from"`
					Immutable      bool          `json:"immutable"`
					Index          []string      `json:"index"`
					Language       string        `json:"language"`
					License        string        `json:"license"`
					MaxSignals     int           `json:"max_signals"`
					Meta           struct {
						From             string `json:"from"`
						KibanaSiemAppUrl string `json:"kibana_siem_app_url"`
					} `json:"meta"`
					Query               string        `json:"query"`
					References          []interface{} `json:"references"`
					RelatedIntegrations []interface{} `json:"related_integrations"`
					RequiredFields      []interface{} `json:"required_fields"`
					RiskScore           int           `json:"risk_score"`
					RiskScoreMapping    []interface{} `json:"risk_score_mapping"`
					RuleId              string        `json:"rule_id"`
					Setup               string        `json:"setup"`
					Severity            string        `json:"severity"`
					SeverityMapping     []struct {
						Field    string `json:"field"`
						Operator string `json:"operator"`
						Severity string `json:"severity"`
						Value    string `json:"value"`
					} `json:"severity_mapping"`
					Threat  []interface{} `json:"threat"`
					To      string        `json:"to"`
					Type    string        `json:"type"`
					Version int           `json:"version"`
				} `json:"parameters"`
				Producer         string        `json:"producer"`
				References       []interface{} `json:"references"`
				RiskScore        int           `json:"risk_score"`
				RiskScoreMapping []interface{} `json:"risk_score_mapping"`
				RuleId           string        `json:"rule_id"`
				RuleTypeId       string        `json:"rule_type_id"`
				Severity         string        `json:"severity"`
				SeverityMapping  []struct {
					Field    string `json:"field"`
					Operator string `json:"operator"`
					Severity string `json:"severity"`
					Value    string `json:"value"`
				} `json:"severity_mapping"`
				Tags      []interface{} `json:"tags"`
				Threat    []interface{} `json:"threat"`
				To        string        `json:"to"`
				Type      string        `json:"type"`
				UpdatedAt time.Time     `json:"updated_at"`
				UpdatedBy string        `json:"updated_by"`
				Uuid      string        `json:"uuid"`
				Version   int           `json:"version"`
			} `json:"rule"`
			Severity    string `json:"severity"`
			Status      string `json:"status"`
			Suppression struct {
				DocsCount int       `json:"docs_count"`
				End       time.Time `json:"end"`
				Start     time.Time `json:"start"`
				Terms     []struct {
					Field string      `json:"field"`
					Value interface{} `json:"value"`
				} `json:"terms"`
			} `json:"suppression"`
			Uuid           string `json:"uuid"`
			WorkflowStatus string `json:"workflow_status"`
		} `json:"alert"`
		SpaceIds []string `json:"space_ids"`
		Version  string   `json:"version"`
	} `json:"kibana"`
	Observer struct {
		Hostname string `json:"hostname"`
		Type     string `json:"type"`
		Version  string `json:"version"`
	} `json:"observer"`
	Parent struct {
		Id string `json:"id"`
	} `json:"parent"`
	Process struct {
		Args   []string `json:"args"`
		Parent struct {
			Pid int `json:"pid"`
		} `json:"parent"`
		Pid   int    `json:"pid"`
		Title string `json:"title"`
	} `json:"process"`
	Processor struct {
		Event string `json:"event"`
		Name  string `json:"name"`
	} `json:"processor"`
	Service struct {
		Language struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"language"`
		Name string `json:"name"`
		Node struct {
			Name string `json:"name"`
		} `json:"node"`
		Runtime struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"runtime"`
	} `json:"service"`
	Signal struct {
		Depth         int `json:"depth"`
		OriginalEvent struct {
			Outcome string `json:"outcome"`
		} `json:"original_event"`
		OriginalTime time.Time `json:"original_time"`
		Reason       string    `json:"reason"`
		Rule         struct {
			Author         []interface{} `json:"author"`
			CreatedAt      time.Time     `json:"created_at"`
			CreatedBy      string        `json:"created_by"`
			Description    string        `json:"description"`
			Enabled        bool          `json:"enabled"`
			FalsePositives []interface{} `json:"false_positives"`
			From           string        `json:"from"`
			Id             string        `json:"id"`
			Immutable      bool          `json:"immutable"`
			Interval       string        `json:"interval"`
			License        string        `json:"license"`
			MaxSignals     int           `json:"max_signals"`
			Name           string        `json:"name"`
			References     []interface{} `json:"references"`
			RiskScore      int           `json:"risk_score"`
			RuleId         string        `json:"rule_id"`
			Severity       string        `json:"severity"`
			Tags           []interface{} `json:"tags"`
			To             string        `json:"to"`
			Type           string        `json:"type"`
			UpdatedAt      time.Time     `json:"updated_at"`
			UpdatedBy      string        `json:"updated_by"`
			Version        int           `json:"version"`
		} `json:"rule"`
		Status string `json:"status"`
	} `json:"signal"`
	Source struct {
		Ip string `json:"ip"`
	} `json:"source"`
	Timestamp1 struct {
		Us int64 `json:"us"`
	} `json:"timestamp"`
	Trace struct {
		Id string `json:"id"`
	} `json:"trace"`
	Transaction struct {
		Duration struct {
			Us int `json:"us"`
		} `json:"duration"`
		Id                  string `json:"id"`
		Name                string `json:"name"`
		RepresentativeCount int    `json:"representative_count"`
		Result              string `json:"result"`
		Sampled             bool   `json:"sampled"`
		SpanCount           struct {
			Dropped int `json:"dropped"`
			Started int `json:"started"`
		} `json:"span_count"`
		Type string `json:"type"`
	} `json:"transaction"`
	Url struct {
		Domain string `json:"domain"`
		Full   string `json:"full"`
		Path   string `json:"path"`
		Port   int    `json:"port"`
		Scheme string `json:"scheme"`
	} `json:"url"`
	UserAgent struct {
		Device struct {
			Name string `json:"name"`
		} `json:"device"`
		Name     string `json:"name"`
		Original string `json:"original"`
		Os       struct {
			Full    string `json:"full"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		Version string `json:"version"`
	} `json:"user_agent"`
}
