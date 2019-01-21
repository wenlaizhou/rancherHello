package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello rancher!"))
		writer.WriteHeader(200)
	})

	http.ListenAndServe(":8077", nil)
}
