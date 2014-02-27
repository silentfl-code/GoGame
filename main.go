package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func main() {
	//основная часть
	http.Handle("/ws", websocket.Handler(WebSocketServer))
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	//log.Println(url)
	http.ServeFile(w, r, url)
}

func WebSocketServer(ws *websocket.Conn) {
	for {
		var in []byte
		err := websocket.Message.Receive(ws, &in)
		log.Printf("Receive: %v\n", string(in))
		if err != nil {
			log.Println(err)
			ws.Close()
			return
		}
		//msg, _ := json.Marshal(struct{ id, value string }{"0", "pong"})
		msg := "pong"
		if string(in) == "ping" {
			websocket.JSON.Send(ws, msg)
		} else {
			websocket.JSON.Send(ws, in)
		}
	}
}
