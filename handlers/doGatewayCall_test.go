package handlers

import (
	"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"net/http/httptest"
	//"net/url"
	cst "ApiGateway/cluster"
	e "ApiGateway/errors"
	mgr "ApiGateway/managers"
	"testing"
)

type challenge struct {
	Answer string `json:"answer"`
	Key    string `json:"key"`
}

func TestGatewayPost_doGatewayCall(t *testing.T) {

	var p passParams
	p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusOK {
		t.Fail()
	}
}

func TestGatewayPost_doGatewayCallParam(t *testing.T) {

	var p passParams
	//p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusInternalServerError {
		t.Fail()
	}
}

// type errReader int

// func (errReader) Read(p []byte) (n int, err error) {
// 	return 0, errors.New("test error")
// }
func TestGatewayPost_doGatewayCallReq(t *testing.T) {

	var p passParams
	p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	//var c challenge
	//c.Answer = "test"
	//c.Key = "test"

	//aJSON, _ := json.Marshal(c)
	//r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r, _ := http.NewRequest("POST", "/test", errReader(0))
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGatewayPost_doGatewayCallBakUrl(t *testing.T) {

	var p passParams
	p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "://challenge.myapigateway.com"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGatewayPost_doGatewayCallBadServiceCall(t *testing.T) {

	var p passParams
	p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.tst"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGatewayPost_doGatewayCallNotFound(t *testing.T) {

	var p passParams
	p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://www.google.com"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusNotFound {
		t.Fail()
	}
}

func TestGatewayPost_doGatewayCallBadResponseBody(t *testing.T) {

	var p passParams
	p.h = new(Handler)
	var clstRt cst.GatewayRoutes

	p.clst = &clstRt
	p.clst.Host = "http://localhost:3011"
	//p.clst.ClientID = 403
	p.clst.APIKey = "403"
	p.gwr = new(mgr.GatewayRoutes)
	p.gwr.ClientID = 403

	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://localhost:3030/test"
	p.fpath = ""
	var q = make(url.Values, 0)
	//q.Set("p1", "param1")
	p.code = &q
	//var c challenge
	//c.Answer = "test"
	//c.Key = "test"

	//aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	p.rts.RouteID = 22
	p.rts.URLID = 33
	var er e.GatewayErrors
	p.e = &er
	p.e.Host = "http://localhost:3011"
	p.e.ClientID = 403
	rtn := doGatewayCall(&p)
	fmt.Print("rtn in doPostPutPatch resp body: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusInternalServerError {
		t.Fail()
	}
}
