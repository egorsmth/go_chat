package main

import (
	"log"
    "net/http"
    "sort"

	"github.com/gorilla/websocket"

	"./db"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func print_headers(w http.ResponseWriter, r *http.Request) {
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

	/*
	fmt.Fprintln(w, "<b>Response Headers:</b></br>")
	for _, k := range responseKeys {
		fmt.Fprintln(w, k, ":", k, "</br>")
	}
	*/
}
	

func main() {
	db.GetUser("qauvlf92vwvebalbhmi7e1h6ycujtlgd")
	// fs := http.FileServer(http.Dir("public"))
	// http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("serving file...")
	// 	fs.ServeHTTP(w, r)
	// 	log.Println("done serving...ZZZZ")
	// 	print_headers(w, r)
	// }))

	// http.HandleFunc("/ws", handleConnections)
	// go handleMessages()
	// log.Print("http server started on :8080")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: &v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: &v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
