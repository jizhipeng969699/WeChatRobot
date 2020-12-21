package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := io.WriteString(writer, "hello world")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n)
	})

	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal(err)
	}

}
