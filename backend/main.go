package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Init function")
	PerformGetRequest()
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
	con, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(con))

}
