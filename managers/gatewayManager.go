package managers

import (
	"fmt"
	"strconv"
	//cl "ApiGateway/cluster"
	ch "ApiGateway/cache"
)

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route    string
	APIKey   string
	ClientID int64
}

//GatewayRouteURL url
type GatewayRouteURL struct {
	RouteID      int64  `json:"routeId"`
	Route        string `json:"route"`
	URLID        int64  `json:"urlId"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Active       bool   `json:"active"`
	CircuitOpen  bool   `json:"circuitOpen"`
	OpenFailCode int    `json:"openFailCode"`
}

//GetGatewayRoute GetGatewayRoute
func (gw *GatewayRoutes) GetGatewayRoute(getActive bool, routeName string) *GatewayRouteURL {
	var rtnVal GatewayRouteURL
	var c ch.RouteCache
	c.ClientID = strconv.FormatInt(gw.ClientID, 10)
	crts := c.GetRoutes(routeName)
	fmt.Print("crts: ")
	fmt.Println(crts)
	if crts != nil {
		// work with cached routes and the delete
	} else {
		//get routes from server and then cache
	}
	return &rtnVal
}
