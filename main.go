package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/egorsmth/go_chat/controllers"
	"github.com/egorsmth/go_chat/shared"
)

func printHeaders(w http.ResponseWriter, r *http.Request) {
	// adding debug header to test (strong/weak) ETags in combination with NGINX
	w.Header().Set("ETag", "HelloWorld")

	var requestKeys []string
	for k := range r.Header {
		requestKeys = append(requestKeys, k)
	}
	sort.Strings(requestKeys)

	var responseKeys []string
	for k := range w.Header() {
		responseKeys = append(responseKeys, k)
	}

	sort.Strings(responseKeys)

	log.Println("request.RequestURI:", r.RequestURI)
	log.Println("request.RemoteAddr:", r.RemoteAddr)
	log.Println("request.TLS:", r.TLS)

	log.Println("Request Headers:")
	for _, k := range requestKeys {
		log.Println(k, ":", r.Header[k])
	}

	fmt.Fprintln(w, "<b>Response Headers:</b></br>")
	for _, k := range responseKeys {
		fmt.Fprintln(w, k, ":", k, "</br>")
	}
}

func main() {
	shared.Init("user=root password=root dbname=social_net sslmode=disable")
	http.HandleFunc("/chat/", controllers.Chat)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/chat/ws", controllers.HandleConnections)
	log.Print("http server started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
