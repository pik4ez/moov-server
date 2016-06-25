package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const logPath = "/tmp/moov.log"

// LocationHandler handles location requests (GET and POST).
func LocationHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// TODO implement
		io.WriteString(w, "not supported\n")
		return
	}
	if req.Method == "POST" {
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		body = append(body, []byte("\n")...)
		if _, err = logFile.Write(body); err != nil {
			log.Fatal(err)
		}
		io.WriteString(w, "POST accepted\n")
	}
}

func main() {
	http.HandleFunc("/location", LocationHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
