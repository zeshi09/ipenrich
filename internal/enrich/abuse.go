package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	abuseApiKey = os.Getenv("ABUSEIPDB_API_KEY")
)

type abuseInfo struct {
	Data struct {
		AbuseConfidenseScore int `json:"abuseConfidenceScore"`
	} `json:"data"`
}

func FetchAbuseScore(ip string) string {
	if abuseApiKey == "" {
		return "-"
	}

	url := fmt.Sprintf("https://api.abuseipdb.com/api/v2/check?ipAddress=%s", ip)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Key", abuseApiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "error"
	}
	defer resp.Body.Close()

	var result abuseInfo

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "error"
	}

	return fmt.Sprintf("%d", result.Data.AbuseConfidenseScore)
}
