package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeshi09/ipenrich/model"
)

func FetchGeoInfo(ip string) (string, string, string) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=country,city,org,query", ip)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", ""
	}
	defer resp.Body.Close()

	var info model.GeoInfo

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return "", "", ""
	}

	if info.Org == "" {
		info.Org = "NF"
	}

	return info.Country, info.City, info.Org
	// return map[string]string{
	// 	"country": info.Country,
	// 	"city":    info.City,
	// 	"org":     info.Org,
	// }
}
