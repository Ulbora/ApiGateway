package handlers

import (
//"bytes"
//"encoding/json"
//"fmt"
//"net/http"
//"net/http/httptest"
//"net/url"
//"testing"
)

type challenge struct {
	Answer string `json:"answer"`
	Key    string `json:"key"`
}

var tgpcid int64 = 46

// //var rrID int64

// //var routeErr int64
// //var routeURLErrID int64
// var connectedForTgp bool
// var gwTgp mgr.GatewayDB

// //var gwRoutesErr mgr.GatewayRoutes
// //var hrr Handler

// func TestGatewayPost_Connect(t *testing.T) {
// 	gwTgp.DbConfig.Host = "localhost:3306"
// 	gwTgp.DbConfig.DbUser = "admin"
// 	gwTgp.DbConfig.DbPw = "admin"
// 	gwTgp.DbConfig.DatabaseName = "ulbora_api_gateway"
// 	connectedForTgp = gwTgp.ConnectDb()
// 	if connectedForTgp != true {
// 		t.Fail()
// 	}
// 	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
// 	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
// 	//cp.Host = "http://localhost:3010"
// 	//testMode = true
// 	//hrr.DbConfig = gwRR.DbConfig
// }

// func TestGatewayPost_doPostPutPatchReq(t *testing.T) {
// 	var p passParams
// 	p.h = new(Handler)
// 	var cbr cb.CircuitBreaker
// 	cbr.DbConfig = gwTgp.DbConfig
// 	p.h.CbDB = cbr
// 	p.b = new(cb.Breaker)
// 	p.gwr = new(mgr.GatewayRoutes)
// 	p.rts = new(mgr.GatewayRouteURL)
// 	p.rts.URL = "http://challenge.myapigateway.com"
// 	p.fpath = "rs/challenge"
// 	var q = make(url.Values, 0)
// 	q.Set("p1", "param1")
// 	p.code = &q

// 	aJSON, _ := json.Marshal(nil)
// 	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	p.r = r
// 	w := httptest.NewRecorder()
// 	p.w = w

// 	//["p1"] = ["param1"]

// 	rtn := doPostPutPatch(&p)
// 	fmt.Print("doPost Res: ")
// 	fmt.Println(rtn)
// 	if rtn.rtnCode != http.StatusBadRequest {
// 		t.Fail()
// 	}
// }

// func TestGatewayPost_doPostPutPatchReq2(t *testing.T) {
// 	var p passParams
// 	p.h = new(Handler)
// 	var cbr cb.CircuitBreaker
// 	cbr.DbConfig = gwTgp.DbConfig
// 	p.h.CbDB = cbr
// 	p.b = new(cb.Breaker)
// 	p.gwr = new(mgr.GatewayRoutes)
// 	p.rts = new(mgr.GatewayRouteURL)
// 	//p.rts.URL = "http://challenge.myapigateway.com"
// 	p.fpath = "rs/challenge"
// 	var q = make(url.Values, 0)
// 	q.Set("p1", "param1")
// 	p.code = &q
// 	var c challenge
// 	c.Answer = "test"
// 	c.Key = "test"

// 	aJSON, _ := json.Marshal(c)
// 	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	p.r = r
// 	w := httptest.NewRecorder()
// 	p.w = w

// 	//["p1"] = ["param1"]

// 	rtn := doPostPutPatch(&p)
// 	fmt.Print("doPost bad req2 Res: ")
// 	fmt.Println(rtn)
// 	if rtn.rtnCode != http.StatusBadRequest {
// 		t.Fail()
// 	}
// }

// func TestGatewayPost_doPostPutPatchMedia(t *testing.T) {
// 	var p passParams
// 	p.h = new(Handler)
// 	var cbr cb.CircuitBreaker
// 	cbr.DbConfig = gwTgp.DbConfig
// 	p.h.CbDB = cbr
// 	p.b = new(cb.Breaker)
// 	p.gwr = new(mgr.GatewayRoutes)
// 	p.rts = new(mgr.GatewayRouteURL)
// 	p.rts.URL = "http://challenge.myapigateway.com"
// 	p.fpath = "rs/challenge"
// 	var q = make(url.Values, 0)
// 	q.Set("p1", "param1")
// 	p.code = &q
// 	var c challenge
// 	c.Answer = "test"
// 	c.Key = "test"

// 	aJSON, _ := json.Marshal(c)
// 	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
// 	//r.Header.Set("Content-Type", "application/json")
// 	p.r = r
// 	w := httptest.NewRecorder()
// 	p.w = w

// 	//["p1"] = ["param1"]

// 	rtn := doPostPutPatch(&p)
// 	fmt.Print("doPost Res: ")
// 	fmt.Println(rtn)
// 	if rtn.rtnCode != http.StatusUnsupportedMediaType {
// 		t.Fail()
// 	}
// }

// func TestGatewayPost_doPostPutPatch(t *testing.T) {
// 	var p passParams
// 	p.h = new(Handler)
// 	var cbr cb.CircuitBreaker
// 	cbr.DbConfig = gwTgp.DbConfig
// 	p.h.CbDB = cbr
// 	p.b = new(cb.Breaker)
// 	p.gwr = new(mgr.GatewayRoutes)
// 	p.rts = new(mgr.GatewayRouteURL)
// 	p.rts.URL = "http://challenge.myapigateway.com"
// 	p.fpath = "rs/challenge"
// 	var q = make(url.Values, 0)
// 	q.Set("p1", "param1")
// 	p.code = &q
// 	var c challenge
// 	c.Answer = "test"
// 	c.Key = "test"

// 	aJSON, _ := json.Marshal(c)
// 	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
// 	r.Header.Set("Content-Type", "application/json")
// 	p.r = r
// 	w := httptest.NewRecorder()
// 	p.w = w

// 	//["p1"] = ["param1"]

// 	rtn := doPostPutPatch(&p)
// 	fmt.Print("doPost Res: ")
// 	fmt.Println(rtn)
// 	if rtn.rtnCode != http.StatusOK {
// 		t.Fail()
// 	}
// }
