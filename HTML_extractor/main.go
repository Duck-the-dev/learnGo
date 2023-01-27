package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

var wg = sync.WaitGroup{}

func main() {

	extractor()
	wg.Wait()
}

func extractor() {
	wg.Add(1)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)

	}
	content := string(data)
	file.WriteString(content)
	wg.Done()
}
