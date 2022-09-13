package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var bar_url = ""
var foo_port = ""

func get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Foo hit!")
	fmt.Fprintln(w, "\n-----------   Foo"+get(bar_url)+"!  -----------------\n")
}

func handleRequests() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":"+foo_port, nil))
}

func main() {
	fetchEnvVariables()
	fmt.Printf("Foo serving on port: %s\n", foo_port)
	handleRequests()
}

func fetchEnvVariables() {
	bar_url = os.Getenv("BAR_URL")
	if bar_url == "" {
		log.Println("BAR_URL not set. Using default")
		bar_url = "http://localhost:6060"
	}
	foo_port = os.Getenv("FOO_PORT")
	if foo_port == "" {
		log.Println("BAR_PORT not set. Using default")
		foo_port = "5050"
	}
	fmt.Printf("Bar url is: %s\n", bar_url)
}
