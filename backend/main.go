package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Init function")
	// PerformGetRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.StatusCode)
	fmt.Println("Content Length:", res.ContentLength)

	var responseString strings.Builder
	con, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(con)
	fmt.Println(byteCount)

	// fmt.Println(string(con))

}

func PerformPostRequest() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
	{
	 "name":" sathwik",
	 "age": 22,
	 "city": "bangalore",
	 "email": "fshaghsu@gmail.com"
	}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	cont, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(cont))
}
