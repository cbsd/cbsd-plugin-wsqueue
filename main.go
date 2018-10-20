package main

import (
	"fmt"
	"flag"
	"log"
	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"

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

	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	message := []byte(*Message_Ptr)

	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Send: [%s]\n", message)
}
