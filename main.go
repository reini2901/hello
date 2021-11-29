package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		_, err := fmt.Fprint(rw, "Hello world")
		if err != nil {
			rw.Write([]byte(err.Error()))
		}

	})

	fmt.Println("Starting app @", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

}
