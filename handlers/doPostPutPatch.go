package handlers

import (
	clst "ApiGateway/cluster"
	e "ApiGateway/errors"
	//"bytes"
	"fmt"
	//"io/ioutil"
	"net/http"
	"time"
)

func doPostPutPatch(p *passParams) *returnVals {
	//fmt.Print("found routes: ")
	//fmt.Println(rts)
	var rtnVals returnVals
	var rtn string
	var rtnCode int

	//var sTime1 = time.Now()
	var sTime2 time.Time
	var eTime1 time.Time
	if paramsOK(p) {
		//var eTime2 time.Time
		var spath = p.rts.URL + "/" + p.fpath + parseQueryString(*p.code)
		//fmt.Print("spath: ")
		//fmt.Println(spath)
		//body := r.Body.Read()
		body, failed := processReqBody(p)
		if !failed {

			// body, err := ioutil.ReadAll(p.r.Body)
			// if err != nil {
			// 	fmt.Println(err)
			// } //else {
			//fmt.Print("Body: ")
			//fmt.Println(string(body))
			//}
			req, reqFailed := getRequest(p.r.Method, spath, body)
			if !reqFailed {
				buildHeaders(p.r, req)
				client := &http.Client{}
				eTime1 = time.Now()

				// add processServiceCall to handler.go

				resp, cErr := client.Do(req)
				sTime2 = time.Now()
				if cErr != nil {
					fmt.Print("Gateway err: ")
					fmt.Println(cErr)
					rtnCode = 400
					rtn = cErr.Error()
					var b clst.Breaker
					b.ClientID = p.gwr.ClientID
					b.FailoverRouteName = p.rts.FailoverRouteName
					b.FailureThreshold = p.rts.FailureThreshold
					b.HealthCheckTimeSeconds = p.rts.HealthCheckTimeSeconds
					b.OpenFailCode = p.rts.OpenFailCode
					b.RestRouteID = p.rts.RouteID
					b.RouteURIID = p.rts.URLID

					p.clst.TripBreaker(&b)
					//p.clst.ClearClusterGwRoutes(p.rts.Route)
					//cbk := p.h.CbDB.GetBreaker(p.b)
					//fmt.Print("cbk: ")
					//fmt.Println(cbk)
					//p.h.CbDB.Trip(cbk)
					//go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, 400, cErr.Error(), p.rts.RouteID, p.rts.URLID)
					var el e.ErrorLog
					el.ClientID = p.gwr.ClientID
					el.ErrCode = rtnCode
					el.RouteID = p.rts.RouteID
					el.RouteURIID = p.rts.URLID
					el.Message = cErr.Error()
					go p.e.SaveErrors(el)
				} else {
					defer resp.Body.Close()
					respbody, err := processResponse(resp) //:= ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Print("Resp Body err: ")
						fmt.Println(err)
						rtnCode = 500
						rtn = err.Error()
						var b clst.Breaker
						b.ClientID = p.gwr.ClientID
						b.FailoverRouteName = p.rts.FailoverRouteName
						b.FailureThreshold = p.rts.FailureThreshold
						b.HealthCheckTimeSeconds = p.rts.HealthCheckTimeSeconds
						b.OpenFailCode = p.rts.OpenFailCode
						b.RestRouteID = p.rts.RouteID
						b.RouteURIID = p.rts.URLID

						p.clst.TripBreaker(&b)
						//p.clst.ClearClusterGwRoutes(p.rts.Route)
						//cbk := p.h.CbDB.GetBreaker(p.b)
						//fmt.Print("cbk: ")
						//fmt.Println(cbk)
						//p.h.CbDB.Trip(cbk)
						//go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, 400, cErr.Error(), p.rts.RouteID, p.rts.URLID)
						var el e.ErrorLog
						el.ClientID = p.gwr.ClientID
						el.ErrCode = rtnCode
						el.RouteID = p.rts.RouteID
						el.RouteURIID = p.rts.URLID
						el.Message = cErr.Error()
						go p.e.SaveErrors(el)
						//cbk := p.h.CbDB.GetBreaker(p.b)
						//p.h.CbDB.Trip(cbk)
						//go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, 500, err.Error(), p.rts.RouteID, p.rts.URLID)
					} else {
						rtn = string(respbody)
						//fmt.Print("Resp Body: ")
						//fmt.Println(rtn)
						rtnCode = resp.StatusCode
						if rtnCode != http.StatusOK {
							//go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, rtnCode, resp.Status, p.rts.RouteID, p.rts.URLID)
							var el e.ErrorLog
							el.ClientID = p.gwr.ClientID
							el.ErrCode = rtnCode
							el.RouteID = p.rts.RouteID
							el.RouteURIID = p.rts.URLID
							el.Message = resp.Status
							go p.e.SaveErrors(el)
						} else {
							//go p.h.CbDB.Reset(p.gwr.ClientID, p.rts.URLID)
							var b clst.ResetBreaker
							b.Route = p.rts.Route
							b.RouteURIID = p.rts.URLID
							p.clst.ResetBreaker(b)
						}
						buildRespHeaders(resp, p.w)
						//w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
					}
				}
			} else {
				rtnCode = 400
				rtn = "Request error"
			}

			// req, rErr := http.NewRequest(p.r.Method, spath, bytes.NewBuffer(*body))
			// if rErr != nil {
			// 	fmt.Print("request err: ")
			// 	fmt.Println(rErr)
			// 	rtnCode = 400
			// 	rtn = rErr.Error()
			// } else {

			// }
		} else {
			rtnCode = 400
			rtn = "Request body error"
		}

	}
	rtnVals.rtnCode = rtnCode
	rtnVals.rtn = rtn
	rtnVals.eTime1 = eTime1
	rtnVals.sTime2 = sTime2
	return &rtnVals
}
