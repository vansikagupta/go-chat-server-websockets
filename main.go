package main

import (
	"fmt"
	"net/http"
	"log"
	ws "github.com/vansikagupta/go-chat-server-websockets/pkg/websocket"
	"github.com/vansikagupta/go-chat-server-websockets/utils"
)

func serveHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HomePage")
}

func serveWs(w http.ResponseWriter, r *http.Request){
	conn, err := ws.Upgrade(w, r)
	if err != nil{
		log.Println(err)
	}
	go ws.Writer(conn)
	ws.Reader(conn)
}

func setUpRoutes(){
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
}

func main(){
	fmt.Println("Go Websockets!!")
	values := make(chan int)
	defer close(values)
	go utils.GenerateValue(values)

	value := <-values
	fmt.Println(value)
	setUpRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))

	
}