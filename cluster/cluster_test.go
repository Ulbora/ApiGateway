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

func Test_getType(t *testing.T) {
	var p GatewayClusterRouteURL
	res := p.GetType()
	if res != "GCRouteURL" {
		t.Fail()
	}

}

func Test_ProcessServiceCallUrl(t *testing.T) {
	var p GatewayClusterRouteURL
	var gw GatewayRoutes
	req, _ := http.NewRequest("POST", "", nil)
	code, err := gw.ProcessServiceCall(req, &p)
	fmt.Print("req code:")
	fmt.Println(code)
	fmt.Print("req suc:")
	fmt.Println(err)
	if code != 0 || !err {
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
	code, err := gw.ProcessServiceCall(req, &rtn)
	fmt.Print("req code:")
	fmt.Println(code)
	fmt.Print("req suc:")
	fmt.Println(err)
	if code != 200 || err {
		t.Fail()
	}
	fmt.Print("rtn:")
	fmt.Println(rtn)
	if len(rtn) == 0 {
		t.Fail()
	}
}
