package flavor
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"gclassec/goClientCompute/openstack"
	"net/http"
	"gclassec/goClientCompute/util"
	"net/url"
	"fmt"
)
type Service struct {
	Session openstack.Session
	Client  http.Client
	URL     string
}
type Response struct {
	Name		string		`json:"name"`
	RAM		int64		`json:"ram"`
	VCPU		int64		`json:"vcpus"`
	Disk		int64		`json:"disk"`
	FlavorID	string		`json:"id"`
}
type DetailResponse struct{
	Name		string		`json:"name"`
	RAM		int64		`json:"ram"`
	VCPU		int64		`json:"vcpus"`
	Disk		int64		`json:"disk"`
	FlavorID	string		`json:"id"`
}

type QueryParameters struct {
	FlavorID		string
	Name			string
	RAM			int64
	VCPU			int64
	Disk			int64

}

type SortDirection string
const (
	// Desc specifies the sort direction to be descending.
	Desc SortDirection = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc SortDirection = "asc"
)
func (flavorService Service) Flavors() (flavor []Response, err error) {
	return flavorService.QueryFlavors(nil)
}

func (flavorService Service) FlavorsDetail() (flavor []DetailResponse, err error) {
	return flavorService.QueryFlavorsDetail(nil)
}

func (flavorService Service) QueryFlavors(queryParameters *QueryParameters) ([]Response, error) {
	flavorContainer := flavorResponse{}
	err := flavorService.queryFlavors(false /*includeDetails*/, &flavorContainer, queryParameters)
	if err != nil {
		return nil, err
	}

	return flavorContainer.Servers, nil
}


func (flavorService Service) QueryFlavorsDetail(queryParameters *QueryParameters) ([]DetailResponse, error) {
	flavorDetailContainer := flavorDetailResponse{}
	err := flavorService.queryFlavors(true /*includeDetails*/, &flavorDetailContainer, queryParameters)
	if err != nil {
		return nil, err
	}

	return flavorDetailContainer.Servers, nil
}

func (flavorService Service) queryFlavors(includeDetails bool, flavorResponseContainer interface{}, queryParameters *QueryParameters) error {
	urlPostFix := "/flavors"
	if includeDetails {
		urlPostFix = urlPostFix + "/detail"
	}
	fmt.Println(urlPostFix)
	reqURL, err := buildQueryURL(flavorService, queryParameters, urlPostFix)
	if err != nil {
		return err
	}

	fmt.Println("RequestURL",reqURL)

	var headers http.Header = http.Header{}
	headers.Set("Accept", "application/json")
	resp, err := flavorService.Session.Get(reqURL.String(), nil, &headers)
	fmt.Println("********************",headers)
	if err != nil {
		return err
	}

	err = util.CheckHTTPResponseStatusCode(resp)
	if err != nil {
		return err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response body printing.")
	fmt.Println(rbody)
	if err != nil {
		return errors.New("aaa")
	}
	if err = json.Unmarshal(rbody, &flavorResponseContainer); err != nil {
		return err
	}
	return nil
}

func buildQueryURL(flavorService Service, queryParameters *QueryParameters, computePartialURL string) (*url.URL, error) {
	reqURL, err := url.Parse(flavorService.URL)
	if err != nil {
		return nil, err
	}

	if queryParameters != nil {
		values := url.Values{}
		if queryParameters.Name != "" {
			values.Set("flavor_name", queryParameters.Name)
		}
		if len(values) > 0 {
			reqURL.RawQuery = values.Encode()
		}
	}
	reqURL.Path += computePartialURL

	return reqURL, nil
}

type flavorDetailResponse struct {
	Servers []DetailResponse `json:"flavors"`
}

type flavorResponse struct {
	Servers []Response `json:"flavors"`
}