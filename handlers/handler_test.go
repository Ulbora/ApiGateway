package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	//"reflect"
	cst "ApiGateway/cluster"
	e "ApiGateway/errors"
	mgr "ApiGateway/managers"
	"testing"
)

type Challenge struct {
	Question string `json:"question"`
	Key      string `json:"key"`
	Answer   string `json:"answer"`
}

func Test_parseQueryString(t *testing.T) {
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	q.Set("p2", "param2")
	rtn := parseQueryString(q)
	fmt.Print("querystring: ")
	fmt.Println(rtn)
	var tpass = false
	if rtn == "?p1=param1&p2=param2" || rtn == "?p2=param2&p1=param1" {
		tpass = true
	}
	if tpass != true {
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

func Test_paramsOK(t *testing.T) {
	p := new(passParams)
	rtn := paramsOK(p)
	if rtn != false {
		t.Fail()
	}
}

func Test_buildHeaders(t *testing.T) {
	pr, _ := http.NewRequest("POST", "/test", nil)
	pr.Header.Set("Content-Type", "application/json")
	sr, _ := http.NewRequest("POST", "/test", nil)
	buildHeaders(pr, sr)
	h := sr.Header
	var key string
	var val string
	for hn, v := range h {
		key = hn
		val = v[0]
	}
	fmt.Print("key: ")
	fmt.Println(key)
	fmt.Print("val: ")
	fmt.Println(val)
	if key != "Content-Type" || val != "application/json" {
		t.Fail()
	}
}

func TestHandler_buildRespHeaders(t *testing.T) {
	//w http.ResponseWriter
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	//pw := new(http.Response)
	//outw := new(http.ResponseWriter)
	//var outw http.ResponseWriter
	outw := httptest.NewRecorder()

	//fmt.Print("pw: ")
	//fmt.Println(pw)
	fmt.Print("outw: ")
	fmt.Println(outw)
	buildRespHeaders(w.Result(), outw)
	// if outw == nil {
	// 	t.Fail()
	// }
	outRes := outw.Result()
	fmt.Print("outRes: ")
	fmt.Println(outRes)

	fmt.Print("outRes.Header: ")
	fmt.Println(outRes.Header)
	if outRes.StatusCode != 200 {
		t.Fail()
	}
	h := outRes.Header
	if len(h) == 0 {
		t.Fail()
	}
	for hn, v := range h {
		fmt.Print("header: ")
		fmt.Println(hn)
		fmt.Print("value: ")
		fmt.Println(v[0])
		if v[0] != "application/json" {
			t.Fail()
		}
	}
}

func Test_processBody(t *testing.T) {
	var c Challenge
	c.Answer = "test"
	c.Key = "123"
	c.Question = "test"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	var p passParams
	p.r = r
	body, fail := processReqBody(&p)
	fmt.Print("body no err: ")
	fmt.Println(body)
	if fail && body != nil {
		t.Fail()
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func Test_processBodyMethod(t *testing.T) {
	var c Challenge
	c.Answer = "test"
	c.Key = "123"
	c.Question = "test"
	//aJSON, _ := json.Marshal(c)
	//var bd []byte
	r, _ := http.NewRequest("POST", "/test", errReader(0))
	var p passParams
	p.r = r
	body, fail := processReqBody(&p)
	fmt.Print("body with err: ")
	fmt.Println(body)
	if !fail && body != nil {
		t.Fail()
	}
}

func Test_getRequest(t *testing.T) {
	var c Challenge
	c.Answer = "test"
	c.Key = "123"
	c.Question = "test"
	aJSON, _ := json.Marshal(c)

	req, failed := getRequest("POST", "http://test", &aJSON)
	fmt.Print("failed: ")
	fmt.Println(failed)
	fmt.Print("req: ")
	fmt.Println(req)
	if failed && req == nil {
		t.Fail()
	}
}

func Test_getRequest2(t *testing.T) {

	req, failed := getRequest("GET", "http://test", nil)
	fmt.Print("failed: ")
	fmt.Println(failed)
	fmt.Print("req: ")
	fmt.Println(req)
	if failed && req == nil {
		t.Fail()
	}
}

func Test_getRequestBadUrl(t *testing.T) {

	req, failed := getRequest("GET", "://test", nil)
	fmt.Print("failed: ")
	fmt.Println(failed)
	fmt.Print("req: ")
	fmt.Println(req)
	if !failed {
		t.Fail()
	}
}

func Test_processServiceCall(t *testing.T) {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	fmt.Print("err: ")
	fmt.Println(err)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, e, failed := processServiceCall(req)
	fmt.Print("get resp: ")
	fmt.Println(resp)
	if failed && resp == nil {
		fmt.Print("service call error: ")
		fmt.Println(e)
		t.Fail()
	}
}

func Test_processServiceCallPost(t *testing.T) {
	var c Challenge
	c.Answer = "test"
	c.Key = "123"
	c.Question = "test"
	aJSON, _ := json.Marshal(c)

	req, err := http.NewRequest("POST", "http://google.com", bytes.NewBuffer(aJSON))
	fmt.Print("Post err: ")
	fmt.Println(err)
	fmt.Print("Post req: ")
	fmt.Println(req)
	resp, e, failed := processServiceCall(req)
	fmt.Print("Post resp: ")
	fmt.Println(resp)
	if failed && resp == nil {
		fmt.Print("service call error: ")
		fmt.Println(e)
		t.Fail()
	}
}

func Test_processServiceCallUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "http://google.tst", nil)
	fmt.Print("err: ")
	fmt.Println(err)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, e, failed := processServiceCall(req)
	fmt.Print("bad get resp: ")
	fmt.Println(resp)
	if !failed && resp != nil {
		fmt.Print("service call error: ")
		fmt.Println(e)
		t.Fail()
	}
}

func Test_tripBreaker(t *testing.T) {

	var p passParams
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.FailoverRouteName = "blue"
	p.rts.FailureThreshold = 1
	p.rts.HealthCheckTimeSeconds = 120
	p.rts.OpenFailCode = 400
	p.rts.RouteID = 22
	p.rts.URLID = 1395
	failed := tripBreaker(&p)
	if failed {
		fmt.Print("trip error: ")
		t.Fail()
	}
}

func Test_tripBreakerErr(t *testing.T) {

	var p passParams
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.FailoverRouteName = "blue"
	p.rts.FailureThreshold = 1
	p.rts.HealthCheckTimeSeconds = 120
	p.rts.OpenFailCode = 400
	p.rts.RouteID = 22
	p.rts.URLID = 1395
	failed := tripBreaker(&p)
	if !failed {
		fmt.Print("trip error: ")
		t.Fail()
	}
}

func Test_sendErrors(t *testing.T) {

	var p passParams
	var er e.GatewayErrors
	p.e = &er

	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.RouteID = 22
	p.rts.URLID = 33
	suc, code := sendErrors(&p, 400, "service call failed")
	// failed := sendErrors(&p, 400, "service call failed")
	if !suc {
		fmt.Print("send error success: ")
		fmt.Println(suc)
		fmt.Print("send error code: ")
		fmt.Println(code)
		t.Fail()
	}
}

func Test_sendErrorsFail(t *testing.T) {

	var p passParams
	var er e.GatewayErrors
	p.e = &er

	p.e.Host = "http://localhost:3011"
	//p.e.ClientID = 403
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.RouteID = 22
	p.rts.URLID = 33
	suc, code := sendErrors(&p, 400, "service call failed")
	// failed := sendErrors(&p, 400, "service call failed")
	fmt.Print("send error success: ")
	fmt.Println(suc)
	fmt.Print("send error code: ")
	fmt.Println(code)
	if suc {
		t.Fail()
	}
}

func Test_getPathParams(t *testing.T) {
	r, _ := http.NewRequest("GET", "/challenge?route=test&rname=blue&fpath=send", nil)
	var pmap = make(map[string]string)
	pmap["route"] = "route1"
	pmap["rname"] = "rname1"
	pmap["fpath"] = "fpath1"
	route, rname, fpath := getPathParams(pmap, r)

	// failed := sendErrors(&p, 400, "service call failed")
	fmt.Print("route: ")
	fmt.Println(route)
	fmt.Print("rname: ")
	fmt.Println(rname)
	fmt.Print("fpath: ")
	fmt.Println(fpath)
	//fmt.Println(code)
	if route != "route1" || rname != "rname1" || fpath != "fpath1" {
		t.Fail()
	}
}

func Test_getPathParamsNil(t *testing.T) {
	r, _ := http.NewRequest("GET", "/challenge?route=route1&rname=rname1&fpath=fpath1", nil)
	var pmap map[string]string
	//pmap["route"] = "route1"
	//pmap["rname"] = "rname1"
	//pmap["fpath"] = "fpath1"
	route, rname, fpath := getPathParams(pmap, r)

	// failed := sendErrors(&p, 400, "service call failed")
	fmt.Print("route: ")
	fmt.Println(route)
	fmt.Print("rname: ")
	fmt.Println(rname)
	fmt.Print("fpath: ")
	fmt.Println(fpath)
	//fmt.Println(code)
	if route != "route1" || rname != "rname1" || fpath != "fpath1" {
		t.Fail()
	}
}

func Test_getPathParamsNone(t *testing.T) {
	r, _ := http.NewRequest("GET", "/challenge", nil)
	var pmap map[string]string
	//pmap["route"] = "route1"
	//pmap["rname"] = "rname1"
	//pmap["fpath"] = "fpath1"
	route, rname, fpath := getPathParams(pmap, r)

	// failed := sendErrors(&p, 400, "service call failed")
	fmt.Print("route: ")
	fmt.Println(route)
	fmt.Print("rname: ")
	fmt.Println(rname)
	fmt.Print("fpath: ")
	fmt.Println(fpath)
	//fmt.Println(code)
	if route != "" || rname != "" || fpath != "" {
		t.Fail()
	}
}
