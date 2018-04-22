package common

import (
	//cl "ApiGateway/cluster"
	"fmt"
	"net/http"
	//"reflect"
	"testing"
)

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

func Test_getRequestBadURL(t *testing.T) {
	var u = ""
	r, err := GetRequest(u, "bad method", nil)
	if err != true {
		t.Fail()
	} else if r != nil {
		t.Fail()
	}

}

func Test_getRequest(t *testing.T) {
	var u = "http://www.myapigateway.com"
	r, err := GetRequest(u, "GET", nil)
	if err == true {
		t.Fail()
	} else if r == nil {
		t.Fail()
	}

}

func Test_getRequestBadURL2(t *testing.T) {
	var u = ""
	var b = []byte("test")
	r, err := GetRequest(u, "bad method", &b)
	if err != true {
		t.Fail()
	} else if r != nil {
		t.Fail()
	}

}

func Test_getRequest2(t *testing.T) {
	var u = "http://www.myapigateway.com"
	var b = []byte("test")
	r, err := GetRequest(u, "POsT", &b)
	if err == true {
		t.Fail()
	} else if r == nil {
		t.Fail()
	}

}

func Test_getResponseNil(t *testing.T) {
	var p GatewayClusterRouteURL
	suc := ProcessRespose(nil, &p)
	if suc != false {
		t.Fail()
	}
}

func Test_getResponseMethod(t *testing.T) {
	var p GatewayClusterRouteURL
	//var r = new(http.Response)
	//r := httptest.NewRecorder()
	client := &http.Client{}
	req, reqperr := http.NewRequest("POST", "http://google.com", nil)
	fmt.Print("reqperr: ")
	fmt.Println(reqperr)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, reserr := client.Do(req)
	fmt.Print("reserr: ")
	fmt.Println(reserr)
	fmt.Print("resp: ")
	fmt.Println(resp)
	suc := ProcessRespose(resp, &p)
	if suc != false {
		t.Fail()
	}
}

func Test_getResponseArrayErr(t *testing.T) {
	var p GatewayClusterRouteURL
	//var r = new(http.Response)
	//r := httptest.NewRecorder()
	client := &http.Client{}
	req, reqperr := http.NewRequest("GET", "http://localhost:3011/rs/cluster/routes/challenge", nil)
	req.Header.Set("u-client-id", "403")
	req.Header.Set("u-api-key", "403")
	fmt.Print("reqperr: ")
	fmt.Println(reqperr)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, reserr := client.Do(req)
	fmt.Print("reserr: ")
	fmt.Println(reserr)
	fmt.Print("resp: ")
	fmt.Println(resp)
	suc := ProcessRespose(resp, &p)
	if suc != false {
		t.Fail()
	}
}

func Test_getResponse(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var r = new(http.Response)
	//r := httptest.NewRecorder()
	var rtn = make([]GatewayClusterRouteURL, 0)
	client := &http.Client{}
	req, reqperr := http.NewRequest("GET", "http://localhost:3011/rs/cluster/routes/challenge", nil)
	req.Header.Set("u-client-id", "403")
	req.Header.Set("u-api-key", "403")
	fmt.Print("reqperr: ")
	fmt.Println(reqperr)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, reserr := client.Do(req)
	fmt.Print("reserr: ")
	fmt.Println(reserr)
	fmt.Print("resp: ")
	fmt.Println(resp)
	defer resp.Body.Close()
	// decoder := json.NewDecoder(resp.Body)
	// error := decoder.Decode(&rtn)
	// if error != nil {
	// 	log.Println(error.Error())
	// }
	// fmt.Print("response: ")
	// fmt.Println(rtn)

	suc := ProcessRespose(resp, &rtn)
	if suc != true {
		t.Fail()
	} else {
		fmt.Print("response: ")
		fmt.Println(rtn)
	}
}
func Test_getJSONEncode(t *testing.T) {
	var p GatewayClusterRouteURL
	j := GetJSONEncode(&p)
	fmt.Print("j: ")
	fmt.Println(j)
	if j == nil {
		t.Fail()
	}
}

func Test_ProcessServiceCallUrl(t *testing.T) {
	var p GatewayClusterRouteURL
	//var gw GatewayRoutes
	req, _ := http.NewRequest("POST", "", nil)
	code := ProcessServiceCall(req, &p)
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
	//var gw GatewayRoutes
	//gw.ClientID = 403
	//gw.APIKey = "403"
	req, _ := http.NewRequest("GET", "http://localhost:3011/rs/cluster/routes/challenge", nil)
	//cid := strconv.FormatInt(gw.ClientID, 10)
	req.Header.Set("u-client-id", "403")
	req.Header.Set("u-api-key", "403")
	code := ProcessServiceCall(req, &rtn)
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
