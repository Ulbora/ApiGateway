package performance

import (
	"fmt"
	"testing"
)

func TestGatewayMonitor_SavePerformanceUrl(t *testing.T) {
	var m GatewayMonitor
	m.ClientID = 403
	m.Host = "://localhost:3011"
	var ml PerformLog
	ml.RouteID = 22
	ml.RouteURIID = 33
	ml.Latency = 10000
	resp, code := m.SavePerformance(ml)
	fmt.Print("resp in monitor cluster: ")
	fmt.Println(resp)
	fmt.Print("code in monitor cluster: ")
	fmt.Println(code)
	if code != 400 {
		t.Fail()
	}
}

func TestGatewayMonitor_SavePerformance(t *testing.T) {
	var m GatewayMonitor
	m.ClientID = 403
	m.Host = "http://localhost:3011"
	var ml PerformLog
	ml.RouteID = 22
	ml.RouteURIID = 33
	ml.Latency = 10000
	resp, code := m.SavePerformance(ml)
	fmt.Print("resp in monitor cluster: ")
	fmt.Println(resp)
	fmt.Print("code in monitor cluster: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}
