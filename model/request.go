package model

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
