package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"math/rand"
)

type Point struct {
	Lat float64   `json:"lat"`
	Lon float64   `json:"lon"`
}

type Object struct {
	Point
	Title string `json:"title"`
}

const host = ":8001"
const path = "/location"
const contentType = "Content-Type"
const contentTypeValue = "application/json"

func LocationHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.Header().Set(contentType, contentTypeValue)
		w.Write([]byte(fmt.Sprintf("{\"datetime\": \"%s\"}\n", time.Now())))

	} else if req.Method == "POST" {

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}

		var point Point
		err = json.Unmarshal(body, &point)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("<- point: lat %f, lon %f\n", point.Lat, point.Lon)


		if rand.Intn(2) == 0 { // 0 or 1
			var o Object
			o.Lat, o.Lon, o.Title = point.Lat, point.Lon, fmt.Sprintf("Some object #%d", rand.Int())
			o.Lat += rand.Float64() / 500
			o.Lon += rand.Float64() / 500

			w.Header().Set(contentType, contentTypeValue)
			body, err = json.Marshal(o)
			if err != nil {
				log.Println(err)
			}
			w.Write(body)
			fmt.Print("-> ")
			fmt.Println(string(body))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	fmt.Println(host)
	http.HandleFunc(path, LocationHandler)
	log.Fatal(http.ListenAndServe(host, nil))
}
