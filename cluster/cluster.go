package cluster

import (
	cm "ApiGateway/common"
	"fmt"
	"net/http"
	"strconv"
)

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route    string
	APIKey   string
	ClientID int64
	Host     string
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

//GatewayClusterResponse ClusterResponse
type GatewayClusterResponse struct {
	Success bool `json:"success"`
}

//GetClusterGwRoutes GetClusterGwRoutes
func (gw *GatewayRoutes) GetClusterGwRoutes(route string) (*[]GatewayClusterRouteURL, int) {
	var code int
	var rtn = make([]GatewayClusterRouteURL, 0)
	var clustURL = gw.Host + "/rs/cluster/routes/" + route
	fmt.Print("clustURL: ")
	fmt.Println(clustURL)
	req, fail := cm.GetRequest(clustURL, http.MethodGet, nil)
	if !fail {
		//var f2 bool
		cid := strconv.FormatInt(gw.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("u-api-key", gw.APIKey)
		code = gw.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("get failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}

//ClearClusterGwRoutes ClearClusterGwRoutes
func (gw *GatewayRoutes) ClearClusterGwRoutes(route string) (*GatewayClusterResponse, int) {
	var code int
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var rtn GatewayClusterResponse
	var clustURL = gw.Host + "/rs/cluster/routes/clear/" + route
	fmt.Print("clustURL: ")
	fmt.Println(clustURL)
	req, fail := cm.GetRequest(clustURL, http.MethodDelete, nil)
	if !fail {
		cid := strconv.FormatInt(gw.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("u-api-key", gw.APIKey)
		code = gw.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("clear failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}

//ProcessServiceCall ProcessCall
func (gw *GatewayRoutes) ProcessServiceCall(req *http.Request, obj interface{}) int {
	var code int
	client := &http.Client{}
	resp, cErr := client.Do(req)
	if cErr != nil {
		fmt.Print("Service err: ")
		fmt.Println(cErr)
		code = http.StatusBadRequest
	} else {
		defer resp.Body.Close()
		cm.ProcessRespose(resp, obj)
		code = resp.StatusCode
	}
	return code
}
