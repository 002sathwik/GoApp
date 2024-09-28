package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://httpbin.org:3000/get?course=go&jedi=master"

func main() {
	fmt.Println("Init function")

	fmt.Println(myurl)
	res, _ := url.Parse(myurl)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Println(qparams["course"])
	fmt.Println(qparams["jedi"])

	for _, val := range qparams {
		fmt.Println("Params are", val)
	}

	partsOfurl := &url.URL{
		Scheme:   "https",
		Host:     "httpbin.org",
		Path:     "/get",
		RawQuery: "course=go&jedi=master",
	}

	aUrl := partsOfurl.String()

	fmt.Println(aUrl)

}

func errs(err error) {
	if err != nil {
		panic(err)
	}
}
