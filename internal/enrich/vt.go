package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zeshi09/ipenrich/model"
)

var (
	apiKey = os.Getenv("VT_API_KEY")
)

func FetchVTStatsIP(ip string) (int, int, int, int) {
	if apiKey == "" {
		return 0, 0, 0, 0
	}
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/ip_addresses/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, 0, 0, 0
	}
	req.Header.Set("x-apikey", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, 0, 0, 0
	}
	defer resp.Body.Close()

	var result model.VTResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, 0, 0
	}

	return result.Data.Attributes.LastAnalysisStats.Harmless, result.Data.Attributes.LastAnalysisStats.Malicious, result.Data.Attributes.LastAnalysisStats.Suspicious, result.Data.Attributes.LastAnalysisStats.Undetected

}
