package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	datos := strings.NewReader(`{"nombre":"Luis"}`)
	res, err := http.Post("https://httpbin.org/post", "application/json", datos)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Proto, res.Status)
	fmt.Println("Content-Type:", res.Header.Get("Content-Type"))

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
