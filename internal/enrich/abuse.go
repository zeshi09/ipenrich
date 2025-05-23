package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zeshi09/ipenrich/model"
)

var (
	abuseApiKey = os.Getenv("ABUSEIPDB_API_KEY")
)

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

	var result model.AbuseInfo

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "error"
	}

	return fmt.Sprintf("%d", result.Data.AbuseConfidenseScore)
}
