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

func Test_ClearClusterResetBreakerUrl(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "://localhost:3011"
	var b ResetBreaker
	b.Route = "challenge"
	b.RouteURIID = 33
	resp, code := gw.ResetBreaker(b)
	fmt.Print("resp in reset breaker: ")
	fmt.Println(resp)
	fmt.Print("code in reset breaker: ")
	fmt.Println(code)
	if code != 400 {
		t.Fail()
	}
}

func Test_ClearClusterResetBreaker1(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	var b ResetBreaker
	b.Route = "challenge"
	b.RouteURIID = 33
	resp, code := gw.ResetBreaker(b)
	fmt.Print("resp in reset breaker: ")
	fmt.Println(resp)
	fmt.Print("code in reset breaker: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}

func Test_ClearClusterTripBreakerUrl(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "://localhost:3011"
	var b Breaker
	b.ClientID = gw.ClientID
	b.FailoverRouteName = "blue"
	b.FailureThreshold = 1
	b.HealthCheckTimeSeconds = 120
	b.OpenFailCode = 400
	b.RestRouteID = 22
	b.RouteURIID = 33

	resp, code := gw.TripBreaker(b)
	fmt.Print("resp in trip breaker: ")
	fmt.Println(resp)
	fmt.Print("code in trip breaker: ")
	fmt.Println(code)
	if code != 400 {
		t.Fail()
	}
}

func Test_ClearClusterTripBreaker1(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	var b Breaker
	b.ClientID = gw.ClientID
	b.FailoverRouteName = "green"
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.OpenFailCode = 400
	b.RestRouteID = 22
	b.RouteURIID = 33

	resp, code := gw.TripBreaker(b)
	fmt.Print("resp in trip breaker: ")
	fmt.Println(resp)
	fmt.Print("code in trip breaker: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}

func Test_ClearClusterTripBreaker2(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	var b Breaker
	b.ClientID = gw.ClientID
	b.FailoverRouteName = "green"
	b.FailureThreshold = 1
	b.HealthCheckTimeSeconds = 120
	b.OpenFailCode = 400
	b.RestRouteID = 22
	b.RouteURIID = 33

	resp, code := gw.TripBreaker(b)
	fmt.Print("resp in trip breaker: ")
	fmt.Println(resp)
	fmt.Print("code in trip breaker: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}

func Test_GetClusterGwRoutes2(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	resp, code := gw.GetClusterGwRoutes("challenge")
	fmt.Print("resp in cluster after trip: ")
	fmt.Println(resp)
	fmt.Print("code in cluster afer trip: ")
	fmt.Println(code)
	if code != 200 || len(*resp) == 0 || !(*resp)[1].CircuitOpen {
		t.Fail()
	}
}

func Test_ClearClusterResetBreaker2(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	var b ResetBreaker
	b.Route = "challenge"
	b.RouteURIID = 33
	resp, code := gw.ResetBreaker(b)
	fmt.Print("resp in reset breaker: ")
	fmt.Println(resp)
	fmt.Print("code in reset breaker: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}

func Test_GetClusterGwRoutes3(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var rtn = make([]GatewayClusterRouteURL, 0)
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	gw.Host = "http://localhost:3011"
	resp, code := gw.GetClusterGwRoutes("challenge")
	fmt.Print("resp in cluster after trip: ")
	fmt.Println(resp)
	fmt.Print("code in cluster afer trip: ")
	fmt.Println(code)
	if code != 200 || len(*resp) == 0 || (*resp)[1].CircuitOpen {
		t.Fail()
	}
}
