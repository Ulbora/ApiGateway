package handlers

import (
	cst "ApiGateway/cluster"
	e "ApiGateway/errors"
	mgr "ApiGateway/managers"
	pm "ApiGateway/performance"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

//HandleGwRoute HandleGwRoute
func (h *Handler) HandleGwRoute(w http.ResponseWriter, r *http.Request) {
	var rtns *returnVals
	var rtn string
	var rtnCode int
	var gw mgr.GatewayRoutes
	gw.APIKey = h.APIKey
	gw.Cache = h.Cache
	gw.ClientID = h.ClientID

	var sTime1 = time.Now()
	var sTime2 time.Time
	var eTime1 time.Time
	var eTime2 time.Time

	vars := mux.Vars(r)
	route, rName, fpath := getPathParams(vars, r)
	fmt.Print("route in gateway: ")
	fmt.Println(route)
	fmt.Print("rName: ")
	fmt.Println(rName)
	fmt.Print("fpath: ")
	fmt.Println(fpath)

	code := r.URL.Query()
	var activeRoute = true
	if rName != "" {
		activeRoute = false
	}
	rts := gw.GetGatewayRoute(activeRoute, route, rName)
	fmt.Print("rts: ")
	fmt.Println(rts)
	if rts.URL == "" {
		fmt.Println("No route found in gateway")
		rtnCode = http.StatusNotFound // rts.OpenFailCode
		rtn = "bad route"
		fmt.Print("found routes: ")
		fmt.Println(rts)
	} else if rts.CircuitOpen {
		fmt.Println("Circuit breaker is open for this route")
		rtnCode = rts.OpenFailCode
		rtn = "Circuit open"
		gw.ReadAndStore(route)
		fmt.Print("found route: ")
		fmt.Println(rts)
	} else {
		var clstRt cst.GatewayRoutes
		var er e.GatewayErrors
		var p passParams
		p.clst = &clstRt
		p.clst.Host = getAPIGatewayURL()
		p.clst.ClientID = h.ClientID
		p.clst.APIKey = h.APIKey
		p.code = &code
		p.fpath = fpath
		p.gwr = &gw
		p.gwr.ClientID = h.ClientID
		p.h = h
		p.r = r
		p.e = &er
		p.e.Host = getAPIGatewayURL()
		p.e.ClientID = h.ClientID
		p.rts = rts
		p.w = w
		rtns = doGatewayCall(&p)
		rtnCode = rtns.rtnCode
		rtn = rtns.rtn
		eTime1 = rtns.eTime1
		sTime2 = rtns.sTime2
		fmt.Print("rtns from gateway: ")
		fmt.Println(rtns)

		// switch r.Method {

		// }
	}

	eTime2 = time.Now()
	dif1 := eTime1.Sub(sTime1)
	dif2 := eTime2.Sub(sTime2)
	tots := dif1.Seconds() + dif2.Seconds()
	//sec := dif.Seconds()
	//fmt.Print("latency sec: ")
	//fmt.Println(tots)
	pms := (tots * 1000000)
	//fmt.Print("latency micros: ")
	//fmt.Println(pms)
	rms := int64(pms + .5)

	var m pm.GatewayMonitor
	m.ClientID = h.ClientID
	m.Host = getAPIGatewayURL()
	var ml pm.PerformLog
	ml.RouteID = rts.RouteID
	ml.RouteURIID = rts.URLID
	ml.Latency = rms
	go m.SavePerformance(ml)
	w.WriteHeader(rtnCode)
	fmt.Fprint(w, rtn)

}
