package performance

import (
	cm "ApiGateway/common"
	"fmt"
	"net/http"
	"strconv"
)

//GatewayMonitor gateway GatewayMonitor
type GatewayMonitor struct {
	ClientID int64
	Host     string
}

//GatewayPerformResponse GatewayPerformResponse
type GatewayPerformResponse struct {
	Success bool `json:"success"`
}

//PerformLog PerformLog
type PerformLog struct {
	RouteID    int64 `json:"routeId"`
	RouteURIID int64 `json:"routeUriId"`
	Latency    int64 `json:"latency"`
}

//SavePerformance SavePerformance
func (m *GatewayMonitor) SavePerformance(obj interface{}) (*GatewayPerformResponse, int) {
	var code int
	var rtn GatewayPerformResponse
	var mURL = m.Host + "/rs/cluster/route/performance"
	//fmt.Print("mURL: ")
	//fmt.Println(mURL)
	j := cm.GetJSONEncode(obj)
	req, fail := cm.GetRequest(mURL, http.MethodPost, j)
	if !fail {
		cid := strconv.FormatInt(m.ClientID, 10)
		req.Header.Set("u-client-id", cid)
		req.Header.Set("Content-Type", "application/json")
		code = cm.ProcessServiceCall(req, &rtn)
	} else {
		fmt.Println("save error failed")
		code = http.StatusBadRequest
	}
	return &rtn, code
}
