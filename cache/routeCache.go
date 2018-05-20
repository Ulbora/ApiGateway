package cache

import (
	cl "ApiGateway/cluster"
	"sync"
)

//RouteCache RouteCache
type RouteCache struct {
	ClientID string
	//Route    string
}

var routeCache = make(map[string]*[]cl.GatewayClusterRouteURL)
var rrrCache = make(map[string]int)
var mu sync.Mutex

//SaveRoutes SaveRoutes
func (c *RouteCache) SaveRoutes(route string, rts *[]cl.GatewayClusterRouteURL) string {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + route
	routeCache[key] = rts
	return key
}

//GetRoutes GetRoutes
func (c *RouteCache) GetRoutes(route string) *[]cl.GatewayClusterRouteURL {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + route
	rtn := routeCache[key]
	return rtn
}

//SaveRrRate SaveRrRate
func (c *RouteCache) SaveRrRate(route string, rrr int) string {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + route + ":rrrate"
	rrrCache[key] = rrr
	return key
}

//GetRrRate GetRrRate
func (c *RouteCache) GetRrRate(route string) int {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + route + ":rrrate"
	rtn := rrrCache[key]
	return rtn
}
