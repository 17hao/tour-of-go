// handler echoes the HTTP request
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\t%s\t%s\n", r.Method, r.URL, r.Proto)
	fmt.Fprint(w, "[HEADERS]\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\t%s = %s\n", k, v)
	}
	fmt.Fprintf(w, "[HOST]\n\t%s\n", r.Host)
	fmt.Fprintf(w, "[REMOTE HOST]\n\t%s\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fmt.Fprint(w, "[QUERY]\n")
	for k, v := range r.Form {
		fmt.Fprintf(w, "\t%s = %s\n", k, v)
	}
}
