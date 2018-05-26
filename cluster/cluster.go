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
	RouteID                int64  `json:"routeId"`
	Route                  string `json:"route"`
	URLID                  int64  `json:"urlId"`
	Name                   string `json:"name"`
	URL                    string `json:"url"`
	Active                 bool   `json:"active"`
	CircuitOpen            bool   `json:"circuitOpen"`
	OpenFailCode           int    `json:"openFailCode"`
	FailoverRouteName      string `json:"failoverRouteName"`
	FailureThreshold       int    `json:"failureThreshold"`
	HealthCheckTimeSeconds int    `json:"healthCheckTimeSeconds"`
}

//Breaker Breaker
type Breaker struct {
	FailureThreshold       int    `json:"failureThreshold"`
	HealthCheckTimeSeconds int    `json:"healthCheckTimeSeconds"`
	FailoverRouteName      string `json:"failoverRouteName"`
	OpenFailCode           int    `json:"openFailCode"`
	PartialOpen            bool   `json:"partialOpen"`
	RouteURIID             int64  `json:"routeUriId"`
	RestRouteID            int64  `json:"routeId"`
	ClientID               int64  `json:"clientId"`
	Route                  string `json:"route"`
}

//ResetBreaker ResetBreaker
type ResetBreaker struct {
	RouteURIID int64  `json:"routeUriId"`
	Route      string `json:"route"`
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
	//fmt.Print("clustURL: ")
	//fmt.Println(clustURL)
	req, fail := cm.GetRequest(clustURL, http.MethodGet, nil)
	if !fail {
		//var f2 bool
		cid := strconv.FormatInt(gw.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("u-api-key", gw.APIKey)
		code = cm.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("get failed")
		code = http.StatusBadRequest
	}
	//fmt.Print("rtn: ")
	//fmt.Println(rtn)
	return &rtn, code
}

//ClearClusterGwRoutes ClearClusterGwRoutes
func (gw *GatewayRoutes) ClearClusterGwRoutes(route string) (*GatewayClusterResponse, int) {
	var code int
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var rtn GatewayClusterResponse
	var clustURL = gw.Host + "/rs/cluster/routes/clear/" + route
	//fmt.Print("clustURL: ")
	//fmt.Println(clustURL)
	req, fail := cm.GetRequest(clustURL, http.MethodDelete, nil)
	if !fail {
		cid := strconv.FormatInt(gw.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("u-api-key", gw.APIKey)
		code = cm.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("clear failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}

//TripBreaker TripBreaker
func (gw *GatewayRoutes) TripBreaker(obj interface{}) (*GatewayClusterResponse, int) {
	var code int
	var rtn GatewayClusterResponse
	var tpURL = gw.Host + "/rs/cluster/route/trip"
	//fmt.Print("tpURL: ")
	//fmt.Println(tpURL)
	j := cm.GetJSONEncode(obj)
	req, fail := cm.GetRequest(tpURL, http.MethodPost, j)
	//fmt.Print("req: ")
	//fmt.Println(req)
	if !fail {
		cid := strconv.FormatInt(gw.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("u-api-key", gw.APIKey)
		req.Header.Set("Content-Type", "application/json")
		//fmt.Print("headers: ")
		//fmt.Println(req.Header)
		code = cm.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("trip breaker failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}

//ResetBreaker ResetBreaker
func (gw *GatewayRoutes) ResetBreaker(obj interface{}) (*GatewayClusterResponse, int) {
	var code int
	var rtn GatewayClusterResponse
	var rURL = gw.Host + "/rs/cluster/route/reset"
	//fmt.Print("rURL: ")
	//fmt.Println(rURL)
	j := cm.GetJSONEncode(obj)
	req, fail := cm.GetRequest(rURL, http.MethodPost, j)
	if !fail {
		cid := strconv.FormatInt(gw.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("u-api-key", gw.APIKey)
		req.Header.Set("Content-Type", "application/json")
		code = cm.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("trip breaker failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}
