package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	url     = "http://app.modelo.io"
	timeout = 5
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
		url := request.FormValue("target")
		log.Println(url)
		c := http.Client{
			Timeout: time.Duration(timeout * time.Second),
		}
		r, err := c.Get(url)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(writer, "status_code %v", 0)

		} else {
			fmt.Fprintf(writer, "status_code %v", r.StatusCode)
			r.Body.Close()
		}
	})

	err := http.ListenAndServe(":22222", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
