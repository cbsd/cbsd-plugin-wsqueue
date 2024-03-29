package main

import (
	"fmt"
	"flag"
	"log"
	"github.com/gorilla/websocket"
)

var origin = "http://127.0.0.1/"

func main() {
	Message_Ptr := flag.String("message", "", "a string")
	Ws_Ptr := flag.String("ws_url", "", "a string")

	flag.Parse()

	if len(*Ws_Ptr) < 1 {
		log.Println("Empty ws_url, use -ws_url=ws://localhost:8023/clonos/jailscontainers/")
		log.Fatal(1)
	}

	if len(*Message_Ptr) < 1 {
		log.Println("Empty message, use -message=\"message to queue\"")
		log.Fatal(1)
	}

	var url=*Ws_Ptr
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	message := []byte(*Message_Ptr)

	err = ws.WriteMessage(websocket.TextMessage, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Send: [%s]\n", message)
}
