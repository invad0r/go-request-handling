package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().Format(format)
		w.Write([]byte("the time is: " + time))
	}
	return http.HandlerFunc(fn)
}

func main() {
	fmt.Println("request handling..")
	// mux := http.NewServeMux()		// poses security risk, global variable which other packages can access !

	var format string = time.RFC1123
	th := timeHandler(format)

	http.Handle("/time", th)

	log.Println("Listening..")
	http.ListenAndServe(":3000", nil)
}
