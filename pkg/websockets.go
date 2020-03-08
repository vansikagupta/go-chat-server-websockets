package websockets

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)
//to upgrade HTTP endpoint to websocket endpoint
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){
	//returns a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, err
}


func Reader(conn *websocket.Conn){
	for{
		//Read message from the connection
		msgType, p, err := conn.ReadMessage()
		//p is slice of byte, []bytes
		//messageType is an int with value websocket.BinaryMessage or websocket.TextMessage
		if err != nil {
			log.Println(err)
			return
		}
		//display on server
		fmt.Println(string(p))

		//write the message back to the same connection
		if err := conn.WriteMessage(msgType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Test(){
	fmt.Println("Test")
}

