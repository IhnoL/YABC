package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var bar_port = ""

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bar hit!")
	fmt.Fprintf(w, "   Bar   ")
}

func handleRequests() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":"+bar_port, nil))
}

func main() {
	bar_port, _ = os.LookupEnv("BAR_PORT")
	if bar_port == "" {
		fmt.Println("BAR_PORT not set. Using default")
		bar_port = "6060"
	}

	fmt.Printf("Bar serving on port: %s\n", bar_port)
	handleRequests()
}
