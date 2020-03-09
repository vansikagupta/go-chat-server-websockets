package websocket

import(
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

type Client struct{
	ID string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct{
	ClientID string
	Type int   //binary or text
    Body string 
}

func (c *Client) Read(){
	defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
	}()
	
	for{
		//Read message from the connection
		msgType, p, err := c.Conn.ReadMessage()
		//p is slice of byte, []bytes
		//messageType is an int with value websocket.BinaryMessage or websocket.TextMessage
		if err != nil {
			log.Println(err)
			return
		}

		//write the message back to the connection pool
		message := Message{ClientID: c.ID, Type: msgType, Body: string(p)}
        c.Pool.Broadcast <- message
        fmt.Printf("Message Received: %+v\n", message)
	}
}