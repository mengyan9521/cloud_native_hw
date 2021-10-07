package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloHandler)
	fmt.Println("Server started at port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {

	header := res.Header()

	header.Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	fmt.Fprint(res, `{"status": "Good"}`)

}
