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

func FetchVTStats(ip string) map[string]int {
	if apiKey == "" {
		return nil
	}
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/ip_addresses/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("x-apikey", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var result model.VTResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil
	}

	return map[string]int{
		"harmless":   result.Data.Attributes.LastAnalysisStats.Harmless,
		"malicious":  result.Data.Attributes.LastAnalysisStats.Malicious,
		"suspicious": result.Data.Attributes.LastAnalysisStats.Suspicious,
		"undetected": result.Data.Attributes.LastAnalysisStats.Undetected,
	}

}
