package handlers

import (
	clst "ApiGateway/cluster"
	"fmt"
	"net/http"
	"time"
)

func doGatewayCall(p *passParams) *returnVals {
	var rtnVals returnVals
	if paramsOK(p) {
		var spath = p.rts.URL + "/" + p.fpath + parseQueryString(*p.code)
		//fmt.Print("spath: ")
		//fmt.Println(spath)
		body, failed := processReqBody(p)
		if !failed {
			req, reqFailed := getRequest(p.r.Method, spath, body)
			if !reqFailed {
				buildHeaders(p.r, req)
				rtnVals.eTime1 = time.Now()
				resp, sErr, failed := processServiceCall(req)
				rtnVals.sTime2 = time.Now()
				if !failed {
					defer resp.Body.Close()
					respbody, err := processResponse(resp) //:= ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Print("Resp Body err: ")
						fmt.Println(err)
						rtnVals.rtnCode = 500
						rtnVals.rtn = err.Error()
						tripBreaker(p)
						go sendErrors(p, rtnVals.rtnCode, err.Error())
					} else {
						rtnVals.rtn = string(respbody)
						//fmt.Print("Resp Body: ")
						//fmt.Println(rtnVals.rtn)
						rtnVals.rtnCode = resp.StatusCode
						if rtnVals.rtnCode != http.StatusOK {
							go sendErrors(p, rtnVals.rtnCode, resp.Status)
						} else {
							var b clst.ResetBreaker
							b.Route = p.rts.Route
							b.RouteURIID = p.rts.URLID
							p.clst.ResetBreaker(b)
						}
						buildRespHeaders(resp, p.w)
					}
				} else {
					rtnVals.rtnCode = 400
					rtnVals.rtn = sErr
					fmt.Print("Gateway err: ")
					fmt.Println(sErr)
					tripBreaker(p)
					go sendErrors(p, rtnVals.rtnCode, sErr)
				}
			} else {
				rtnVals.rtnCode = 400
				rtnVals.rtn = "Request error"
			}
		} else {
			rtnVals.rtnCode = 400
			rtnVals.rtn = "Request body error"
		}

	} else {
		fmt.Print("Missing parameters: ")
		fmt.Println(p)
		rtnVals.rtnCode = 500
		rtnVals.rtn = "Missing parameters"
	}
	return &rtnVals
}
