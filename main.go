package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/vansikagupta/go-chat-server-websockets/pkg/websocket"
)

func serveHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HomePage")
}

func serveWs(w http.ResponseWriter, r *http.Request){
	conn, err := websocket.Upgrade(w, r)
	if err != nil{
		log.Println(err)
	}

	websocket.Reader(conn)
}

func setUpRoutes(){
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
}

func main(){
	fmt.Println("Go Websockets!!")
	setUpRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))
}