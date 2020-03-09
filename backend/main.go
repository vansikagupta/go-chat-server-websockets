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

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	username,_ := r.URL.Query()["username"]
	fmt.Println("Websocket Endpoint Hit by client",username[0])
	conn, err := websocket.Upgrade(w, r)
	if err != nil{
		log.Println(err)
	}
	
	//new client created and registered everytime end point is hit 
	client := &websocket.Client{
		ID: username[0],
		Conn: conn,// new connection
		Pool: pool, // same pool
	}

	pool.Register <- client
	client.Read()
}

func setUpRoutes(){
	
	pool := websocket.NewPool()
	go pool.Start()// we don't want to create new pool everytime end point is hit
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main(){
	fmt.Println("Go Websockets!!")
	setUpRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))

	
}