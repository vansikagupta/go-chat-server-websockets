package main

import (
	"fmt"
	"net/http"
	"log"
	ws "github.com/vansikagupta/go-chat-server-websockets/pkg/websockets"
)

func serveHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HomePage")
}

func serveWs(w http.ResponseWriter, r *http.Request){
	conn, err := ws.Upgrade(w, r)
	if err != nil{
		log.Println(err)
	}

	ws.Reader(conn)
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