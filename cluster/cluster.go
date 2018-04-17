package cluster

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route    string
	APIKey   string
	ClientID int64
}

//ServiceParam ServiceParam
type ServiceParam interface {
	GetType() string
}

//GatewayClusterRouteURL url
type GatewayClusterRouteURL struct {
	RouteID           int64  `json:"routeId"`
	Route             string `json:"route"`
	URLID             int64  `json:"urlId"`
	Name              string `json:"name"`
	URL               string `json:"url"`
	Active            bool   `json:"active"`
	CircuitOpen       bool   `json:"circuitOpen"`
	OpenFailCode      int    `json:"openFailCode"`
	FailoverRouteName string `json:"failoverRouteName"`
}

//GetType GetType
func (gw *GatewayClusterRouteURL) GetType() string {
	return "GatewayClusterRouteURL"
}

//GetClusterGwRoutes GetClusterGwRoutes
func (gw *GatewayRoutes) GetClusterGwRoutes() *[]GatewayClusterRouteURL {
	var rtn = make([]GatewayClusterRouteURL, 0)
	return &rtn
}

//ProcessCall ProcessCall
func (gw *GatewayRoutes) ProcessCall(req *http.Request, obj ServiceParam) int {
	cid := strconv.FormatInt(gw.ClientID, 10)
	req.Header.Set("u-client-id", cid)
	req.Header.Set("u-api-key", gw.APIKey)
	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}
	client := &http.Client{}
	resp, cErr := client.Do(req)
	if cErr != nil {
		fmt.Print("Service GET err: ")
		fmt.Println(cErr)
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&obj)
		if error != nil {
			log.Println(error.Error())
		}
	}
	return resp.StatusCode
}

func getRequest(url string, method string) (*http.Request, bool) {
	//var gURL = c.Host + "/rs/content/get/" + id + "/" + clientID
	//fmt.Println(gURL)
	//resp, err := http.Get(gURL)
	var err = false
	req, rErr := http.NewRequest(method, url, nil)
	if rErr != nil {
		err = true
		fmt.Print("request err: ")
		fmt.Println(rErr)
	}
	return req, err
}
