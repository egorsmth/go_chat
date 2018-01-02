package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/websocket"

	"./models"
	"./shared"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

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

type chatResponse struct {
	User      *models.User
	ChatRooms *[]models.UserToChatRoom
}

func chat(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println(err)
		// redirect tp login if no session
	}
	user, err := models.GetUserFromSession(sid.Value)
	if err != nil {
		log.Println("In main GetUserFromSession error:", err)
		// redirect tp login if no session
	}

	t := template.New("chat")                           // Create a template.
	t = template.Must(t.ParseFiles("public/chat.html")) // Parse template file.
	if err != nil {
		log.Println("parse file err:", err)
	}
	cr := chatResponse{}
	cr.User = user
	cr.ChatRooms, err = models.GetChatRooms(user)
	err = t.Execute(w, cr)
	if err != nil {
		log.Println("template Execute err:", err)
	}

	//printHeaders(w, r)

}

func main() {
	shared.Init("user=root password=root dbname=social_net sslmode=disable")
	http.Handle("/chat", http.HandlerFunc(chat))
	//http.Handle("/chat-room", http.HandlerFunc(chatRoom))
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	//
	// http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("serving file...")
	// 	fs.ServeHTTP(w, r)
	// 	log.Println("done serving...ZZZZ")
	// 	print_headers(w, r)
	// }))

	// http.HandleFunc("/ws", handleConnections)
	// go handleMessages()
	log.Print("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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
			log.Fatal(err)
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
				log.Fatal(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
