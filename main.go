package main

import (
	"fmt"
	"net/http"
	"os"
)

const port = ":8080"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		bar := os.Getenv("foo")

		_, err := fmt.Fprintf(rw, "Hello %s!", bar)
		if err != nil {
			rw.Write([]byte(err.Error()))
		}

	})

	fmt.Println("Starting app @", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

}
