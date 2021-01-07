package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("hii")
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: config,
	}
	netClient := &http.Client{
		Transport: transport,
	}
	baseURL := "https://youtube.com/pewdiepie"

	response, err := netClient.Get(baseURL)
	checkError(err)

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	response.Body.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
