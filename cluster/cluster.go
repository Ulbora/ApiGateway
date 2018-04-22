package cluster

import (
	//by "bytes"
	//"encoding/json"
	"fmt"
	//"log"
	cm "ApiGateway/common"
	"net/http"
	//"strconv"
)

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route    string
	APIKey   string
	ClientID int64
}

// //ServiceParam ServiceParam
// type ServiceParam interface {
// 	GetType() string
// }

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
	return "GCRouteURL"
}

// //GetClusterGwRoutes GetClusterGwRoutes
// func (gw *GatewayRoutes) GetClusterGwRoutes() *[]GatewayClusterRouteURL {
// 	var rtn = make([]GatewayClusterRouteURL, 0)
// 	return &rtn
// }

//ProcessServiceCall ProcessCall
func (gw *GatewayRoutes) ProcessServiceCall(req *http.Request, obj interface{}) (int, bool) {
	var code int
	var err bool
	//cid := strconv.FormatInt(gw.ClientID, 10)
	//req.Header.Set("u-client-id", cid)
	//req.Header.Set("u-api-key", gw.APIKey)
	//if req.Method == http.MethodPost || req.Method == http.MethodPut {
	//req.Header.Set("Content-Type", "application/json")
	//}
	client := &http.Client{}
	resp, cErr := client.Do(req)
	if cErr != nil {
		fmt.Print("Service err: ")
		fmt.Println(cErr)
		err = true
	} else {
		defer resp.Body.Close()
		suc := cm.ProcessRespose(resp, obj)
		if suc {
			code = resp.StatusCode
		}
		// decoder := json.NewDecoder(resp.Body)
		// error := decoder.Decode(&obj)
		// if error != nil {
		// 	log.Println(error.Error())
		// }
	}
	return code, err
}
