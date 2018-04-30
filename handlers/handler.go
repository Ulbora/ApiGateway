package handlers

import (
	ch "ApiGateway/cache"
	cst "ApiGateway/cluster"
	e "ApiGateway/errors"
	mgr "ApiGateway/managers"
	//"fmt"
	//"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
	//"net/http"
)

//Handler Handler
type Handler struct {
	APIKey   string
	ClientID int64
	Cache    ch.RouteCache
}

type passParams struct {
	h     *Handler
	rts   *mgr.GatewayRouteURL
	fpath string
	code  *url.Values
	gwr   *mgr.GatewayRoutes
	clst  *cst.GatewayRoutes
	e     *e.GatewayErrors
	w     http.ResponseWriter
	r     *http.Request
}

type returnVals struct {
	rtnCode int
	rtn     string
	eTime1  time.Time
	sTime2  time.Time
}

func parseQueryString(vals url.Values) string {
	var rtn = ""
	var first = true
	for key, value := range vals {
		if first == true {
			first = false
			rtn += "?" + key + "=" + value[0]
		} else {
			rtn += "&" + key + "=" + value[0]
		}
	}
	return rtn
}

func getAPIGatewayURL() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY") != "" {
		rtn = os.Getenv("API_GATEWAY")
	} else {
		rtn = "http://localhost:3011"
	}
	return rtn
}

func paramsOK(p *passParams) bool {
	var rtn = true
	if p.clst == nil || p.code == nil || p.gwr == nil || p.h == nil || p.r == nil || p.rts == nil || p.w == nil {
		rtn = false
	}
	return rtn
}

func buildHeaders(pr *http.Request, sr *http.Request) {
	h := pr.Header
	for hn, v := range h {
		//fmt.Print("header: ")
		//fmt.Print(hn)
		//fmt.Print(" value: ")
		//fmt.Println(v[0])
		sr.Header.Set(hn, v[0])
	}
}

func buildRespHeaders(inw *http.Response, outw http.ResponseWriter) {
	//fmt.Print("resp mock: ")
	//fmt.Println(inw)
	h := inw.Header
	//fmt.Print("header mock: ")
	//fmt.Println(h)
	//var cnt = 0

	for hn, v := range h {
		// cnt++
		// fmt.Print("header to send: ")
		// fmt.Println(hn)
		// fmt.Print("value to send: ")
		// fmt.Println(v[0])
		// if cnt > 5 {
		// 	break
		// }
		outw.Header().Set(hn, v[0])
	}
}
