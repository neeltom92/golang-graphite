package main

import (
	"fmt"
	graphite "github.com/marpaia/graphite-golang"
	"log"
	"net/http"
)

func init() {

	host := "graphite.server.endpoint" // replace this endpoint with the actual Graphite server endpoint

	port := 2003 // Graphite server port

	var err error

	Graphite, err = graphite.NewGraphite(host, port)

	fmt.Println(err)

}

var Graphite *graphite.Graphite

func sliCalculation(apiEndpoint string) {
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	value := 1
	if resp.StatusCode == 200 {
		value = 1
		err = Graphite.SimpleSend("http.response.code", fmt.Sprintf("%v", value))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		value = 0
		err = Graphite.SimpleSend("http.response.code", fmt.Sprintf("%v", value))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	sliCalculation("https://api.example.com") // replace the endpoint with actual endpoint
}
