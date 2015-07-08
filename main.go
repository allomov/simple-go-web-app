package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	fmt.Fprintf(response, "<h1>Hey there!</h1>\n")
}

func main() {
	var host = flag.String("host", "", "IP of host to run webserver on")
	var port = flag.Int("port", 8080, "Port to run webserver on")

	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
