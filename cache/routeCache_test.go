package cache

import (
	cl "ApiGateway/cluster"
	"testing"
)

func TestRouteCache_SaveRoutes(t *testing.T) {
	var c RouteCache
	c.ClientID = "123"
	//c.Route = "content"
	var routes = make([]cl.GatewayClusterRouteURL, 0)
	var r1 cl.GatewayClusterRouteURL
	r1.Route = "content"
	r1.RouteID = 1
	r1.URLID = 11
	r1.URL = "content/test"
	routes = append(routes, r1)

	var r2 cl.GatewayClusterRouteURL
	r2.Route = "content"
	r2.RouteID = 1
	r2.URLID = 11
	r2.URL = "content/test"
	routes = append(routes, r2)

	rtn := c.SaveRoutes("content", &routes)
	if rtn != "123:content" {
		t.Fail()
	}
}

func TestRouteCache_GetRoutes(t *testing.T) {
	var c RouteCache
	c.ClientID = "123"
	rtn := c.GetRoutes("content")
	if len(*rtn) != 2 || (*rtn)[0].RouteID != 1 {
		t.Fail()
	}
}

func TestRouteCache_SaveRrRate(t *testing.T) {
	var c RouteCache
	c.ClientID = "123"

	rtn := c.SaveRrRate("content", 1)
	if rtn != "123:content:rrrate" {
		t.Fail()
	}
}

func TestRouteCache_GetRrRate(t *testing.T) {
	var c RouteCache
	c.ClientID = "123"
	rtn := c.GetRrRate("content")
	if rtn != 1 {
		t.Fail()
	}
}
