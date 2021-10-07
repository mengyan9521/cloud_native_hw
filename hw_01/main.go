/*完成作业的内容
1. 接收客户端 request，并将 request 中带的 header 写入 response header*/

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

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		for _, v2 := range v {
			w.Header().Set(k, v2)
			fmt.Printf("%s : %s \n", k, v2)
		}
	}
}
