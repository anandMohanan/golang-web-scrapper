package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}
)

func main() {

	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("Provide a url")
		os.Exit(1)
	}

	baseURL := arguments[0]
	fmt.Println(baseURL)

	getURL(baseURL)

}

func getURL(href string) {
	fmt.Printf(href)
	response, err := netClient.Get(href)
	checkError(err)
	defer response.Body.Close()

	links, err := extractlinks.All(response.Body)

	checkError(err)

	for _, link := range links {
		// fmt.Printf("index %v -- link %+v \n", i, link)
		getURL(link.Href)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
