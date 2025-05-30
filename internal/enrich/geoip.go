package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeshi09/ipenrich/model"
)

func FetchGeoInfo(ip string) map[string]string {
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=country,city,org,query", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var info model.GeoInfo

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil
	}

	if info.Org == "" {
		info.Org = "NF"
	}

	return map[string]string{
		"country": info.Country,
		"city":    info.City,
		"org":     info.Org,
	}
}
