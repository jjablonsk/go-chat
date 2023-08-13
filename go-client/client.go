package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

func main() {
	nick := os.Args[1:]
	origin := "From UNKNOWN"
	if nick != nil {
		origin = fmt.Sprintf("From %s", nick)
	}
	url := "ws://localhost:3000/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	go func(ws *websocket.Conn) {
		watchForInput(ws)
	}(ws)

	go func(ws *websocket.Conn) {
		printMsg(ws)
	}(ws)

	<-make(chan bool)
}

func watchForInput(ws *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		message = strings.TrimSpace(message)
		_, err = ws.Write([]byte(message))
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
		fmt.Printf("Received: %s.\n", msg[:n])
	}
}
