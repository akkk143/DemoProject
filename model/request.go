package model

import "time"

type Url struct {
	Path   string `json:"path"`
	Scheme string `json:"scheme"`
	Port   int    `json:"port"`
	Domain string `json:"domain"`
	Full   string `json:"full"`
}
type Service struct {
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
}

type Http struct {
	Request struct {
		Headers struct {
			Origin          []string `json:"Origin"`
			Locale          []string `json:"Locale"`
			SecChUa         []string `json:"Sec-Ch-Ua"`
			Email           []string `json:"Email"`
			Accept          []string `json:"Accept"`
			Requestid       []string `json:"Requestid"`
			Authority       []string `json:"Authority"`
			SecChUaPlatform []string `json:"Sec-Ch-Ua-Platform"`
			TapApiToken     []string `json:"Tap-Api-Token"`
			Referer         []string `json:"Referer"`
			UserAgent       []string `json:"User-Agent"`
			Connection      []string `json:"Connection"`
			SecFetchSite    []string `json:"Sec-Fetch-Site"`
			SecFetchDest    []string `json:"Sec-Fetch-Dest"`
			AcceptEncoding  []string `json:"Accept-Encoding"`
			Traceparent     []string `json:"Traceparent"`
			SecFetchMode    []string `json:"Sec-Fetch-Mode"`
			PostmanToken    []string `json:"Postman-Token"`
			AcceptLanguage  []string `json:"Accept-Language"`
			SecChUaMobile   []string `json:"Sec-Ch-Ua-Mobile"`
			ContentLength   []string `json:"Content-Length"`
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
}

type Transaction struct {
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
}

type Test struct {
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
	Timestamp time.Time `json:"@timestamp"`
	Parent    struct {
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
	Agent struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"agent"`
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
	Client struct {
		Ip string `json:"ip"`
	} `json:"client"`
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
				Referer         []string `json:"Referer"`
				TapApiToken     []string `json:"Tap-Api-Token"`
				Connection      []string `json:"Connection"`
				UserAgent       []string `json:"User-Agent"`
				SecFetchDest    []string `json:"Sec-Fetch-Dest"`
				SecFetchSite    []string `json:"Sec-Fetch-Site"`
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
