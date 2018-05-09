package main

import (
	ch "ApiGateway/cache"
	hd "ApiGateway/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

var h hd.Handler
var c ch.RouteCache

func main() {
	var cid string
	var key string
	cid = os.Args[1]
	key = os.Args[2]
	//fmt.Print("cid: ")
	//fmt.Println(cid)

	//fmt.Print("key: ")
	//fmt.Println(key)

	h.Cache = c
	if cid == "" {
		cid = getClientID()
	}

	if key == "" {
		key = getAPIKey()
	}
	h.ClientID, _ = strconv.ParseInt(cid, 10, 0)
	h.APIKey = key

	//fmt.Print("cid: ")
	//fmt.Println(h.ClientID)

	//fmt.Print("key: ")
	//fmt.Println(h.APIKey)

	fmt.Println("Api Gateway running inside network on port 3020!")
	router := mux.NewRouter()
	//gateway routes
	router.HandleFunc("/np/{route}/{rname}/{fpath:[^.]+}", h.HandleGwRoute)
	router.HandleFunc("/{route}/{fpath:[^ ]+}", h.HandleGwRoute)
	//disgard -- router.HandleFunc("/{route}/{fpath:[^.]+}", handleGwRoute)
	router.HandleFunc("/{route}", h.HandleGwRoute)
	http.ListenAndServe(":3020", router)

}

func getClientID() string {
	return os.Getenv("API_GATEWAY_CLIENT_ID")
}

func getAPIKey() string {
	return os.Getenv("API_GATEWAY_API_KEY")
}
