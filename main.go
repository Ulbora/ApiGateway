package main

import (
	hd "ApiGateway/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Api Gateway running inside network on port 3020!")
	router := mux.NewRouter()
	var h hd.Handler
	//gateway routes
	router.HandleFunc("/np/{route}/{rname}/{fpath:[^.]+}", h.HandleGwRoute)
	router.HandleFunc("/{route}/{fpath:[^ ]+}", h.HandleGwRoute)
	//disgard -- router.HandleFunc("/{route}/{fpath:[^.]+}", handleGwRoute)
	router.HandleFunc("/{route}", h.HandleGwRoute)
	http.ListenAndServe(":3020", router)

}
