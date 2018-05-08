package handlers

import (
	ch "ApiGateway/cache"
	cst "ApiGateway/cluster"
	e "ApiGateway/errors"
	mgr "ApiGateway/managers"
	"bytes"
	"fmt"
	"io/ioutil"
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
		if first {
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

func processReqBody(p *passParams) (*[]byte, bool) {
	var fail = false
	var body []byte
	if p.r.Method == http.MethodPost || p.r.Method == http.MethodPut || p.r.Method == http.MethodPatch {
		var err error
		//make this fail
		body, err = ioutil.ReadAll(p.r.Body)
		if err != nil {
			fail = true
			fmt.Print("process body err: ")
			fmt.Println(err)
		}
	}
	return &body, fail
}

func getRequest(method string, url string, body *[]byte) (*http.Request, bool) {
	var failed = false
	var req *http.Request
	var rErr error
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req, rErr = http.NewRequest(method, url, bytes.NewBuffer(*body))
	} else {
		req, rErr = http.NewRequest(method, url, nil)
	}
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
		failed = true
		//rtnCode = 400
		//rtn = rErr.Error()
	}
	return req, failed
}

func processServiceCall(req *http.Request) (*http.Response, string, bool) {
	client := &http.Client{}
	resp, cErr := client.Do(req)
	var e string
	var failed bool
	if cErr != nil {
		fmt.Print("Service call err: ")
		fmt.Println(cErr)
		e = cErr.Error()
		failed = true
	}
	return resp, e, failed
}

func tripBreaker(p *passParams) bool {
	var rtn bool
	var b cst.Breaker
	b.ClientID = p.gwr.ClientID
	b.FailoverRouteName = p.rts.FailoverRouteName
	b.FailureThreshold = p.rts.FailureThreshold
	b.HealthCheckTimeSeconds = p.rts.HealthCheckTimeSeconds
	b.OpenFailCode = p.rts.OpenFailCode
	b.RestRouteID = p.rts.RouteID
	b.RouteURIID = p.rts.URLID

	resp, code := p.clst.TripBreaker(&b)
	if !resp.Success || code != http.StatusOK {
		fmt.Print("trip resp: ")
		fmt.Println(resp)

		fmt.Print("trip code: ")
		fmt.Println(code)
		rtn = true
	}
	return rtn
}

func sendErrors(p *passParams, errCode int, errMessage string) (bool, int) {
	//var rtn = false
	var el e.ErrorLog
	el.ClientID = p.gwr.ClientID
	el.ErrCode = errCode
	el.RouteID = p.rts.RouteID
	el.RouteURIID = p.rts.URLID
	el.Message = errMessage
	//go p.e.SaveErrors(el)
	resp, code := p.e.SaveErrors(el)
	//fmt.Print("send err resp: ")
	//fmt.Println(resp)

	//fmt.Print("send err code: ")
	//fmt.Println(code)
	return resp.Success, code
}

func getPathParams(vars map[string]string, r *http.Request) (string, string, string) {
	var route string
	var rName string
	var fpath string
	if vars != nil {
		route = vars["route"]
		rName = vars["rname"]
		fpath = vars["fpath"]
	} else {
		route = r.URL.Query().Get("route")
		rName = r.URL.Query().Get("rname")
		fpath = r.URL.Query().Get("fpath")
	}
	return route, rName, fpath
}
