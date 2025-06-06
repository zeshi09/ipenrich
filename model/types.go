package model

type EnrichedHash struct {
	LogFile string `json:"log_file"`
	Hash    string `json:"hash"`

	VTMalicious  int `json:"vt_malicious"`
	VTSuspicious int `json:"vt_suspicious"`
	VTHarmless   int `json:"vt_harmless"`
	VTUndetected int `json:"vt_undetected"`
}

type EnrichedIP struct {
	LogFile string `json:"log_file"`
	Ip      string `json:"ip"`
	Dns     string `json:"dns_name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Org     string `json:"org"`

	AbuseScore int `json:"abuse_score"`

	VTMalicious  int `json:"vt_malicious"`
	VTSuspicious int `json:"vt_suspicious"`
	VTHarmless   int `json:"vt_harmless"`
	VTUndetected int `json:"vt_undetected"`
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
