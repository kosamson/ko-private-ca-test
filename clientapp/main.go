package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("SSL_CERT_DIR: %s", os.Getenv("SSL_CERT_DIR"))
	log.Printf("KO_DATA_PATH: %s", os.Getenv("KO_DATA_PATH"))

	doRequest("https://host.docker.internal:8443/hello")
	doRequest("https://httpbingo.org/ip")
}

func doRequest(addr string) {
	log.Println("starting request")

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to do http request: %v", err)
	}

	defer res.Body.Close()

	log.Printf("response status code: %v", res.StatusCode)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	log.Printf("response body: %v", string(bodyBytes))
}
