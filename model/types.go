package model

type EnrichedIP struct {
	LogFile string
	Ip      string
	Geo     map[string]string
	Abuse   int
	Vt      map[string]int
}

type GeoInfo struct {
	Query   string `json:"query"`
	Country string `json:"country"`
	City    string `json:"city"`
	Org     string `json:"org"`
}

type AbuseInfo struct {
	Data struct {
		AbuseConfidenseScore int `json:"abuseConfidenceScore"`
	} `json:"data"`
}

type VTResponse struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Harmless   int `json:"harmless"`
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}
