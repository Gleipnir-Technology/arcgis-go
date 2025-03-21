package main

import "encoding/json"
import "flag"
import "fmt"
import "io"
import "net/http"
import "net/url"

type ArcGIS struct {
	ServiceRoot string
	TenantId string
	Token string
}

type AuthInfo struct {
	isTokenBasedSecurity bool
	tokenServiceUrl string
}
type RestInfo struct {
	CurrentVersion float64
	FullVersion string
	OwningSystemUrl string
	OwningTenant string
	AuthInfo AuthInfo
}

func (arcgis ArcGIS) info() (*RestInfo, error) {  
	u := fmt.Sprintf("%s/%s/arcgis/rest/info", arcgis.ServiceRoot, arcgis.TenantId)
	base, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
    params.Add("f", "json")
    params.Add("token", arcgis.Token)
    base.RawQuery = params.Encode()

	fmt.Println("Requesting", base.String())
	resp, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Received body", string(body[:]))
	var i RestInfo
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func main() {
	service_root := flag.String("root", "https://machine.domain.com", "The root URL for the ArcGIS service to contact.")
	tenant_id := flag.String("tenant", "abc123", "The tenant ID to use when communicating with the service.")
	token := flag.String("token", "secretstuff", "The API token to use to communicate with the service.")
	flag.Parse()

	arcgis := ArcGIS {*service_root, *tenant_id, *token}
	info, err := arcgis.info()
	if err != nil {
		fmt.Println("Failed: %s", err)
	}
	fmt.Println("Current version: ", info.CurrentVersion)
}
