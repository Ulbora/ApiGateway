package cluster

import (
	"fmt"
	"testing"
)

func Test_GetClusterGwRoutesReq(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "://localhost:3011"
	resp, code := gw.GetClusterGwRoutes("challenge")
	fmt.Print("resp in cluster: ")
	fmt.Println(resp)
	fmt.Print("code in cluster: ")
	fmt.Println(code)
	if code != 400 {
		t.Fail()
	}
}

func Test_GetClusterGwRoutesResp(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://google.com"
	resp, code := gw.GetClusterGwRoutes("challenge")
	fmt.Print("resp in cluster: ")
	fmt.Println(resp)
	fmt.Print("code in cluster: ")
	fmt.Println(code)
	if code != 404 {
		t.Fail()
	}
}
func Test_GetClusterGwRoutes(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	resp, code := gw.GetClusterGwRoutes("challenge")
	fmt.Print("resp in cluster: ")
	fmt.Println(resp)
	fmt.Print("code in cluster: ")
	fmt.Println(code)
	if code != 200 && len(*resp) == 0 {
		t.Fail()
	}
}

func Test_ClearClusterGwRoutesUrl(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "://localhost:3011"
	resp, code := gw.ClearClusterGwRoutes("challenge")
	fmt.Print("resp in clear cluster: ")
	fmt.Println(resp)
	fmt.Print("code in clear cluster: ")
	fmt.Println(code)
	if code != 400 {
		t.Fail()
	}
}

func Test_ClearClusterGwRoutes(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	resp, code := gw.ClearClusterGwRoutes("challenge")
	fmt.Print("resp in clear cluster: ")
	fmt.Println(resp)
	fmt.Print("code in clear cluster: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}
