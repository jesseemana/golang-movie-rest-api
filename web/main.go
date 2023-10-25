package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("handling web stuff in Go")
	urls()
	makeGetRequest()
}

func makeGetRequest() {
	const url = "http://localhost/8080/recipes"

	response, err := http.Get(url)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content lenght:", response.ContentLength)

	var response_string strings.Builder // for reading strings too

	data, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(data))

	count, err := response_string.Write(data)
	checkNilErr(err)
	fmt.Println(count)

	fmt.Println(response_string.String())
}

func urls() {
	var myurl string = "https://lco.dev:3000/learn?coursename=node_masterclass&payment=gl09asd0123"

	result, err := url.Parse(myurl)
	checkNilErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	params := result.Query()
	fmt.Println(params["coursename"])

	// Constructing a new url 
	urlparts := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=jesse",
	}

	fmt.Println("The newly constructed URL is:", urlparts)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
