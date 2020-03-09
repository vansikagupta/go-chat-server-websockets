package websocket

import(
	"fmt"
	"log"
)

type Pool struct{
	 Register 	chan *Client
	 Unregister chan *Client
	 Clients 	map[*Client]bool // map Cleint active/not-active
	 Broadcast	chan Message
}

//returns a new connection Pool
func NewPool() *Pool {
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) Start(){
	for{
		//Go’s select lets you wait on multiple channel operations
		select{
		case client := <-pool.Register:
			pool.Clients[client] = true //adding new entry in Clients map
			fmt.Println("No of clients connected : ", len(pool.Clients))
			//broadcast "New user joined" message to all clients
			for client, _ := range pool.Clients{
				fmt.Println(client.ID)
				client.Conn.WriteMessage(1, []byte("New User joined..."))
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)//removing client from pool
			fmt.Println("No of clients connected : ", len(pool.Clients))
			//broadcast "User disconnected" message to all clients
			for client, _ := range pool.Clients{
				fmt.Println(client.ID)
				client.Conn.WriteMessage(1, []byte("User disconnected"))
			}

		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all connected clients")
			for client, _ := range pool.Clients{
				fmt.Println(client.ID)
				msg := message.Body + " From: "+ message.ClientID
				if err := client.Conn.WriteMessage(message.Type, []byte(msg)); err != nil{
					log.Println(err)
                    return
				}
			}
		}
	}
}