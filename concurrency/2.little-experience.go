package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitCh chan struct{} // 0 bytes
	msgCh  chan string
}

func newServer() *Server {
	return &Server{
		quitCh: make(chan struct{}),
		msgCh:  make(chan string),
	}
}

func (s *Server) start() {
	fmt.Println("Starting server...")
	s.loop()
}

func (s *Server) printMessage(message string) {
	fmt.Println("Output the message: ", message)
}

func (s *Server) sendMessage(message string) {
	s.msgCh <- message
}

func (s *Server) quit() {
	close(s.quitCh)
}

func (s *Server) loop() {
mainloop: // name for "for loop"
	for {
		select {
		case <-s.quitCh:
			// quit the channel
			break mainloop
		case msg := <-s.msgCh:
			// do some operation with data get from channel
			s.printMessage(msg)
		}
	}
}

func littleExperience() {
	server := newServer()
	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	fmt.Println("Hello")
	server.start()

}
