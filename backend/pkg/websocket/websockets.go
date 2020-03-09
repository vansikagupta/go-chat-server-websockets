package websocket

import ( 
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

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){
	//returns a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, err
}

