package cluster

import (
	"strconv"
	//"encoding/json"
	"fmt"
	//"log"
	"net/http"
	//"net/http/httptest"
	//"net/http"
	//"reflect"
	"testing"
)

// func Test_getType(t *testing.T) {
// 	var p GatewayClusterRouteURL
// 	res := p.GetType()
// 	if res != "GCRouteURL" {
// 		t.Fail()
// 	}

// }

func Test_ProcessServiceCallUrl(t *testing.T) {
	var p GatewayClusterRouteURL
	var gw GatewayRoutes
	req, _ := http.NewRequest("POST", "", nil)
	code := gw.ProcessServiceCall(req, &p)
	fmt.Print("req code:")
	fmt.Println(code)
	//fmt.Print("req suc:")
	//fmt.Println(err)
	if code != 400 {
		t.Fail()
	}
}

func Test_ProcessServiceCall(t *testing.T) {
	//var p GatewayClusterRouteURL
	var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	req, _ := http.NewRequest("GET", "http://localhost:3011/rs/cluster/routes/challenge", nil)
	cid := strconv.FormatInt(gw.ClientID, 10)
	req.Header.Set("u-client-id", cid)
	req.Header.Set("u-api-key", gw.APIKey)
	code := gw.ProcessServiceCall(req, &rtn)
	fmt.Print("req code in processCall:")
	fmt.Println(code)
	fmt.Print("req failed in processCall:")
	//fmt.Println(failed)
	if code != 200 {
		t.Fail()
	}
	fmt.Print("rtn in processCall:")
	fmt.Println(rtn)
	if len(rtn) == 0 {
		t.Fail()
	}
}

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
