package errors

import (
	"fmt"
	"testing"
)

func TestGatewayErrors_SaveErrorsUrl(t *testing.T) {
	var e GatewayErrors
	e.ClientID = 403
	e.Host = "://localhost:3011"
	var el ErrorLog
	el.ClientID = e.ClientID
	el.ErrCode = 400
	el.RouteID = 22
	el.RouteURIID = 33
	el.Message = "error from cluster gateway api"
	resp, code := e.SaveErrors(el)
	fmt.Print("resp in clear cluster: ")
	fmt.Println(resp)
	fmt.Print("code in clear cluster: ")
	fmt.Println(code)
	if code != 400 {
		t.Fail()
	}
}

func TestGatewayErrors_SaveErrorsReq(t *testing.T) {
	var e GatewayErrors
	e.ClientID = 4034
	e.Host = "http://localhost:3011"
	var el ErrorLog
	el.ClientID = e.ClientID
	el.ErrCode = 400
	el.RouteID = 22
	el.RouteURIID = 33
	el.Message = "error from cluster gateway api"
	resp, code := e.SaveErrors(el)
	fmt.Print("resp in error cluster: ")
	fmt.Println(resp)
	fmt.Print("code in error cluster: ")
	fmt.Println(code)
	if code != 400 || resp.Success {
		t.Fail()
	}
}

func TestGatewayErrors_SaveErrors(t *testing.T) {
	var e GatewayErrors
	e.ClientID = 403
	e.Host = "http://localhost:3011"
	var el ErrorLog
	el.ClientID = e.ClientID
	el.ErrCode = 400
	el.RouteID = 22
	el.RouteURIID = 33
	el.Message = "error from cluster gateway api"
	resp, code := e.SaveErrors(el)
	fmt.Print("resp in error cluster: ")
	fmt.Println(resp)
	fmt.Print("code in error cluster: ")
	fmt.Println(code)
	if code != 200 || !resp.Success {
		t.Fail()
	}
}
