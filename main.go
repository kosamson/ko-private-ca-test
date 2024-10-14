package main

import (
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println("received hello request")
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte("This is an example server.\n")); err != nil {
		panic(err)
	}
	log.Println("sent hello response")
}

func main() {
	http.HandleFunc("/hello", HelloServer)

	log.Println("serving https")

	if err := http.ListenAndServeTLS(
		":8443",
		"ca-setup/kosamson.site.crt",
		"ca-setup/kosamson.site.key",
		nil,
	); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
