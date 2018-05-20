package managers

import (
	ch "ApiGateway/cache"
	cl "ApiGateway/cluster"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

var c ch.RouteCache
var gw GatewayRoutes

func TestGatewayRoutes_GetGatewayRoutes(t *testing.T) {
	//var c ch.RouteCache
	//var gw GatewayRoutes
	os.Setenv("CACHE_REFRESH_RATE", "2")
	gw.ClientID = 403
	gw.APIKey = "403"
	c.ClientID = strconv.FormatInt(gw.ClientID, 10)
	gw.Cache = c
	chrt := gw.GetGatewayRoute(true, "challenge", "")
	fmt.Print("route: ")
	fmt.Println(chrt)
	if !chrt.Active {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRoutesNotActive(t *testing.T) {
	//var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	c.ClientID = strconv.FormatInt(gw.ClientID, 10)
	gw.Cache = c
	chrt := gw.GetGatewayRoute(false, "challenge", "blue")
	fmt.Print("route: ")
	fmt.Println(chrt)
	time.Sleep(time.Second * 2)
	cr := c.GetRoutes("challenge")
	fmt.Print("cashed routes after read: ")
	fmt.Println(cr)
	if len(*cr) == 0 {
		t.Fail()
	}
	if chrt.Active {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRoutesNotActive2(t *testing.T) {
	//var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	c.ClientID = strconv.FormatInt(gw.ClientID, 10)
	gw.Cache = c
	chrt := gw.GetGatewayRoute(false, "challenge", "blue")
	fmt.Print("route: ")
	fmt.Println(chrt)
	time.Sleep(time.Second * 2)
	cr := c.GetRoutes("challenge")
	fmt.Print("cashed routes after read: ")
	fmt.Println(cr)
	if len(*cr) == 0 {
		t.Fail()
	}
	if chrt.Active {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRoutesNotActive3(t *testing.T) {
	//var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	c.ClientID = strconv.FormatInt(gw.ClientID, 10)
	gw.Cache = c
	chrt := gw.GetGatewayRoute(false, "challenge", "blue")
	fmt.Print("route: ")
	fmt.Println(chrt)
	time.Sleep(time.Second * 2)
	cr := c.GetRoutes("challenge")
	fmt.Print("cashed routes after read: ")
	fmt.Println(cr)
	os.Setenv("CACHE_REFRESH_RATE", "")
	if len(*cr) == 0 {
		t.Fail()
	}
	if chrt.Active {
		t.Fail()
	}
}

func TestGatewayRoutes_readAndStore(t *testing.T) {
	//var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	c.ClientID = strconv.FormatInt(gw.ClientID, 10)
	gw.Cache = c
	rts := gw.ReadAndStore("challenge")
	fmt.Print("read routes: ")
	fmt.Println(rts)
	if len(*rts) == 0 {
		t.Fail()
	}
}

func Test_getAPIGatewayURL(t *testing.T) {
	h := getAPIGatewayURL()
	if h != "http://localhost:3011" {
		t.Fail()
	}
}

func Test_getAPIGatewayURLEnv(t *testing.T) {
	os.Setenv("API_GATEWAY", "test")
	h := getAPIGatewayURL()
	if h != "test" {
		t.Fail()
	}
}

func Test_parseGatewayRoutes(t *testing.T) {
	//var rts = make([]cl.GatewayClusterRouteURL, 0)
	var r1 cl.GatewayClusterRouteURL
	r1.Active = true
	r1.CircuitOpen = false
	r1.Name = "challenge"
	r1.OpenFailCode = 401
	r1.Route = "challenge"
	r1.RouteID = 1
	r1.URL = "test1"
	r1.URLID = 1
	res := parseGatewayRoutes(r1)
	if res.Active != true || res.CircuitOpen != false || res.Name != "challenge" || res.OpenFailCode != 401 ||
		res.Route != "challenge" || res.RouteID != 1 || res.URL != "test1" || res.URLID != 1 {
		t.Fail()
	}
}

func Test_getCacheRefreshRate(t *testing.T) {
	h := getCacheRefreshRate()
	if h != 10 {
		t.Fail()
	}
}

func Test_getCacheRefreshRateEV(t *testing.T) {
	os.Setenv("CACHE_REFRESH_RATE", "20")
	h := getCacheRefreshRate()
	fmt.Print("refresh rate: ")
	fmt.Println(h)
	if h != 20 {
		t.Fail()
	}
}
