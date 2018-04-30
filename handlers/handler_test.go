package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	//"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

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
