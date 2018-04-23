package managers

import (
	"fmt"
	"testing"
)

func TestGatewayRoutes_GetGatewayRoutes(t *testing.T) {
	var gw GatewayRoutes
	gw.ClientID = 403
	gw.APIKey = "403"
	chrt := gw.GetGatewayRoute(true, "challenge")
	fmt.Print("route: ")
	fmt.Println(chrt)
}
