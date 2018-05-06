package errors

import (
	cm "ApiGateway/common"
	"fmt"
	"net/http"
	"strconv"
)

//GatewayErrors gateway errors
type GatewayErrors struct {
	ClientID int64
	Host     string
}

//ErrorLog ErrorLog
type ErrorLog struct {
	ClientID   int64  `json:"clientId"`
	ErrCode    int    `json:"errorCode"`
	Message    string `json:"message"`
	RouteID    int64  `json:"routeId"`
	RouteURIID int64  `json:"routeUriId"`
}

//GatewayErrorResponse GatewayErrorResponse
type GatewayErrorResponse struct {
	Success bool `json:"success"`
}

//SaveErrors SaveErrors
func (e *GatewayErrors) SaveErrors(obj interface{}) (*GatewayErrorResponse, int) {
	var code int
	var rtn GatewayErrorResponse
	var errURL = e.Host + "/rs/cluster/route/error"
	//fmt.Print("errURL: ")
	//fmt.Println(errURL)
	j := cm.GetJSONEncode(obj)
	req, fail := cm.GetRequest(errURL, http.MethodPost, j)
	if !fail {
		cid := strconv.FormatInt(e.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("Content-Type", "application/json")
		code = cm.ProcessServiceCall(req, &rtn)
		if code != http.StatusOK {
			fmt.Print("errURL: ")
			fmt.Println(errURL)
			fmt.Print("Save errors rtn: ")
			fmt.Println(rtn)
			fmt.Print("save error code: ")
			fmt.Println(code)

		}
	} else {
		fmt.Println("save error failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}
