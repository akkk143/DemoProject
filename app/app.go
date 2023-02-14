package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const serviceRoute = "/demo/project"

var srv *http.Server
var messages []string
var count int

// Start starts the http server
func Start() {
	messages = append(messages, "Hello World!")
	count = 0
	createServer()
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		panic(err)
	}
}

func createServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc(serviceRoute+"/printmsg", printMsg).Methods("GET")
	r.HandleFunc(serviceRoute+"/customquery/alert", customQueryAlert).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func printMsg(w http.ResponseWriter, r *http.Request) {
	for _, m := range messages {
		_, err := w.Write([]byte(m))
		if err != nil {
			log.Printf("couldnt write response error [%s]\n", err)
		}
	}
}

type ServiceData struct {
	ServiceName string
	APIUrl      string
	StatusCode  int
}

func customQueryAlert(w http.ResponseWriter, r *http.Request) {
	var body test

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		messages = append(messages, err.Error())
	}

	bytes, err := json.Marshal(body)

	if err != nil {
		messages = append(messages, err.Error())
	}

	data := test{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		messages = append(messages, err.Error())
	}
	api := ServiceData{}

	api.ServiceName = data.AlertData.Service.Name
	api.StatusCode = data.AlertData.Http.Response.StatusCode
	api.APIUrl = data.AlertData.Url.Full

	str := fmt.Sprint(api.APIUrl, " , ", api.StatusCode, " , ", api.ServiceName)

	if len(messages) > 5 {
		messages = make([]string, 0)
	}
	messages = append(messages, str)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	} else {
		log.Println("hello")
	}
}

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

type AlertData struct {
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

type test struct {
	AlertData AlertData `json:"alert_data"`
}
