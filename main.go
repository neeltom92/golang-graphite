package main

import (
	"fmt"
	graphite "github.com/marpaia/graphite-golang"
	"log"
	"net/http"
)

func init() {

	host := "172.31.183.235"

	port := 2008

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

	fmt.Println("The status code we got is:", resp.StatusCode)
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
	sliCalculation("https://test.exmaple.com")
	fmt.Println("\n")

}
