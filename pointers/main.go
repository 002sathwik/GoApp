package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://httpbin.org/get"

func main() {
	fmt.Println("Init function")

	res, err := http.Get(url)

	errs(err)

	fmt.Printf("Responce is of type :%T\n", res)
	defer res.Body.Close()

	databytes, err := ioutil.ReadAll(res.Body)
	errs(err)

	content := string(databytes)
	fmt.Print(content)

}

func errs(err error) {
	if err != nil {
		panic(err)
	}
}
