package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	//"net/http"
	ch "ApiGateway/cache"
	//e "ApiGateway/errors"
	//mgr "ApiGateway/managers"
	"testing"
)

func TestHandler_HandleGwRoute(t *testing.T) {
	var c ch.RouteCache
	var h Handler
	h.APIKey = "403"
	h.ClientID = 403
	h.Cache = c
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.HandleGwRoute(w, r)
	fmt.Print("resp in gateway: ")
	fmt.Println(w.Result())
	fmt.Print("body in gateway: ")
	fmt.Println(w.Body)
	fmt.Print("Code in gateway: ")
	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestHandler_HandleGwRouteNorout(t *testing.T) {
	var c ch.RouteCache
	var h Handler
	h.APIKey = "403"
	h.ClientID = 403
	h.Cache = c
	r, _ := http.NewRequest("GET", "/challenge?route=challenge1&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.HandleGwRoute(w, r)
	fmt.Print("resp in gateway: ")
	fmt.Println(w.Result())
	fmt.Print("body in gateway: ")
	fmt.Println(w.Body)
	fmt.Print("Code in gateway: ")
	fmt.Println(w.Code)
	if w.Code != 404 {
		t.Fail()
	}
}

//make this open after addedin circuit breaker
func TestHandler_HandleGwRouteOpen(t *testing.T) {
	var c ch.RouteCache
	var h Handler
	h.APIKey = "403"
	h.ClientID = 403
	h.Cache = c
	r, _ := http.NewRequest("GET", "/challenge?route=challenge&rname=red&fpath=rs/challenge/en_us?g=g&b=b", nil)
	w := httptest.NewRecorder()
	h.HandleGwRoute(w, r)
	fmt.Print("resp in gateway: ")
	fmt.Println(w.Result())
	fmt.Print("body in gateway: ")
	fmt.Println(w.Body)
	fmt.Print("Code in gateway for open circuit: ")
	fmt.Println(w.Code)
	if w.Code != 500 && w.Code != 400 {
		t.Fail()
	}
}
