package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	connnections map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		connnections: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWebSocket(ws *websocket.Conn) {
	fmt.Println("Incomming connection", ws.RemoteAddr())

	s.connnections[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read err", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("Thanks! "))
	}

}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWebSocket))
	http.ListenAndServe(":3000", nil)
}
