package main

import (
	"fmt"
	"net/http"
	"os"
)

const port = ":8080"
const foo = "foo"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		bar := os.Getenv(foo)

		_, err := fmt.Fprintf(rw, "Hello %s!", bar)
		if err != nil {
			rw.Write([]byte(err.Error()))
		}

	})

	http.HandleFunc("/liveness", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("liveness probe called")
	})

	http.HandleFunc("/readiness", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("readyness probe called")
	})

	fmt.Println("Starting app @", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

}
