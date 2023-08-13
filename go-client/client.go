package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

type Message struct {
	Sender string `json:"sender"`
	Msg    []byte `json:"msg"`
}

func main() {
	nick := os.Args[1]
	origin := "ws://localhost:3000/"
	url := "ws://localhost:3000/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	go func(ws *websocket.Conn) {
		watchForInput(ws, nick)
	}(ws)

	go func(ws *websocket.Conn) {
		printMsg(ws)
	}(ws)

	<-make(chan bool)
}

func watchForInput(ws *websocket.Conn, sender string) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		message := Message{
			Sender: sender,
			Msg:    []byte(strings.TrimSpace(text)),
		}
		json_message, _ := json.Marshal(message)

		_, err = ws.Write(json_message)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func printMsg(ws *websocket.Conn) {
	for true {
		var msg = make([]byte, 512)
		var n int
		var err error
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		var message Message
		json.Unmarshal([]byte(msg[:n]), &message)
		fmt.Printf("%s say: %s \n", message.Sender, message.Msg)
	}
}
