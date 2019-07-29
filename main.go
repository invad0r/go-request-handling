package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("this time is : " + tm))
}

func timeHandler2(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func timeHandler3(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

// shorter form for timeHandler3
func timeHandler4(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	})
}

// shorter form for timeHandler3 using implicit conversion
func timeHandler5(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}

func main() {
	fmt.Println("request handling..")
	mux := http.NewServeMux()

	// (1)
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	// (2)
	th := http.HandlerFunc(timeHandler2)
	mux.Handle("/time2", th)

	// shorter using go shortcute mux.HandleFuc method
	mux.HandleFunc("time22", timeHandler2)

	// (3)
	log.Println("Listening..")
	http.ListenAndServe(":3000", mux)
}
