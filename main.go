package main

import (
	"fmt"
	"log"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	say, err := cowsay.Say(
		cowsay.Phrase("Hello"),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, say)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
