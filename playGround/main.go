package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const MyUrl string = "http://127.0.0.1:5500/learnGo/HTML_extractor/index.html"

func main() {
	res, err := url.Parse(MyUrl)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.Get(MyUrl)

	if err != nil {
		fmt.Println(err)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Println(resp.Header["Date"])
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.Scheme)
}
