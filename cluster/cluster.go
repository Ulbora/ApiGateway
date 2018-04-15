package cluster

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route    string
	APIKey   string
	ClientID int64
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

// //GetClusterGwRoutes GetClusterGwRoutes
// func (gw *GatewayRoutes) GetClusterGwRoutes() *[]GatewayClusterRouteURL {
// 	var rtn = make([]GatewayClusterRouteURL, 0)

// 	return &rtn
// }
