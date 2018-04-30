package managers

import (
	//"fmt"
	"os"
	"strconv"
	//cl "ApiGateway/cluster"
	ch "ApiGateway/cache"
	cl "ApiGateway/cluster"
)

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route            string
	APIKey           string
	ClientID         int64
	Cache            ch.RouteCache
	CacheRefreshRate int
}

//GatewayRouteURL url
type GatewayRouteURL struct {
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

//GetGatewayRoute GetGatewayRoute
func (gw *GatewayRoutes) GetGatewayRoute(getActive bool, route string, routeName string) *GatewayRouteURL {
	var rtnVal GatewayRouteURL
	var rtn *[]cl.GatewayClusterRouteURL
	crts := gw.Cache.GetRoutes(route)
	//fmt.Print("crts: ")
	//fmt.Println(crts)
	if crts != nil {
		// work with cached routes and the delete
		rtn = crts
		gw.handleRefresh(route)
	} else {
		rtn = gw.readAndStore(route)
		//fmt.Print("code: ")
		//fmt.Println(code)
	}
	if len(*rtn) > 0 && getActive {
		for r := range *rtn {
			if (*rtn)[r].Active {
				rtnVal = parseGatewayRoutes((*rtn)[r])
				break
			}
		}
	} else if len(*rtn) > 0 {
		for r := range *rtn {
			if (*rtn)[r].Name == routeName {
				rtnVal = parseGatewayRoutes((*rtn)[r])
				break
			}
		}
	}
	return &rtnVal
}

func getAPIGatewayURL() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY") != "" {
		rtn = os.Getenv("API_GATEWAY")
	} else {
		rtn = "http://localhost:3011"
	}
	return rtn
}

func getCacheRefreshRate() int {
	var rtn = 10
	if os.Getenv("CACHE_REFRESH_RATE") != "" {
		rr := os.Getenv("CACHE_REFRESH_RATE")
		rtn, _ = strconv.Atoi(rr)
	}
	return rtn
}

func parseGatewayRoutes(rt cl.GatewayClusterRouteURL) GatewayRouteURL {
	var rtn GatewayRouteURL
	rtn.Active = rt.Active
	rtn.CircuitOpen = rt.CircuitOpen
	rtn.Name = rt.Name
	rtn.OpenFailCode = rt.OpenFailCode
	rtn.Route = rt.Route
	rtn.RouteID = rt.RouteID
	rtn.URL = rt.URL
	rtn.URLID = rt.URLID
	rtn.FailoverRouteName = rt.FailoverRouteName
	rtn.FailureThreshold = rt.FailureThreshold
	rtn.HealthCheckTimeSeconds = rt.HealthCheckTimeSeconds
	return rtn
}

func (gw *GatewayRoutes) readAndStore(route string) *[]cl.GatewayClusterRouteURL {
	var clst cl.GatewayRoutes
	clst.ClientID = gw.ClientID
	clst.APIKey = gw.APIKey
	clst.Host = getAPIGatewayURL()
	rtn, _ := clst.GetClusterGwRoutes(route)
	//fmt.Print("resp: ")
	//fmt.Println(rtn)
	gw.Cache.SaveRoutes(route, rtn)
	return rtn
}

func (gw *GatewayRoutes) handleRefresh(route string) {
	//fmt.Print("CacheRefreshRate before if: ")
	//fmt.Println(gw.CacheRefreshRate)
	//fmt.Print("getCacheRefreshRate before if: ")
	//fmt.Println(getCacheRefreshRate())
	if gw.CacheRefreshRate >= getCacheRefreshRate() {
		gw.CacheRefreshRate = 0
		//fmt.Print("CacheRefreshRate: ")
		//fmt.Println(gw.CacheRefreshRate)
		go func(gwr *GatewayRoutes, rt string) {
			gwr.readAndStore(rt)
		}(gw, route)
	} else {
		gw.CacheRefreshRate++
		//fmt.Print("CacheRefreshRate: ")
		//fmt.Println(gw.CacheRefreshRate)
	}
}
