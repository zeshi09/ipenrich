package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	apiKey = os.Getenv("VT_API_KEY")
)

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

func FetchVTStats(ip string) string {
	if apiKey == "" {
		return "-"
	}
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/ip_addresses/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "error"
	}
	req.Header.Set("x-apikey", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()

	var result VTResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "error"
	}

	stats := result.Data.Attributes.LastAnalysisStats
	total := stats.Harmless + stats.Malicious + stats.Suspicious + stats.Undetected
	return fmt.Sprintf("%d/%d", stats.Malicious, total)
}
