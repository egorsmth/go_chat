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
	http.Handle("/chat", http.HandlerFunc(controllers.Chat))
	http.Handle("/chat_rooms", http.HandlerFunc(controllers.ChatRooms))
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	//
	// http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("serving file...")
	// 	fs.ServeHTTP(w, r)
	// 	log.Println("done serving...ZZZZ")
	// 	print_headers(w, r)
	// }))

	http.HandleFunc("/ws", controllers.HandleConnections)
	// go handleMessages()
	log.Print("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer ws.Close()

// 	clients[ws] = true
// 	for {
// 		var msg Message
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			log.Fatal(err)
// 			delete(clients, ws)
// 			break
// 		}
// 		broadcast <- msg
// 	}
// }
