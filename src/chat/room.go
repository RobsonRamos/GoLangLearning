package main

import (
	"github.com/RobsonRamos/GoLangLearning/GoLangLearning/chat/trace"
	"log"
	"github.com/gorilla/websocket"
	"net/http"
)

type room struct{
	forward 	chan []byte
	join 		chan *Client
	leave 		chan *Client
	clients		map[*Client]bool
	tracer		trace.Tracer
}


func NewRoom() *room{

	return &room{
		forward : 	make(chan []byte),
		join : 		make(chan *Client),
		leave: 		make(chan *Client),
		clients:	make(map[*Client]bool),
		tracer : 	trace.Off(),
	}
}


func (r *room)run() {
	for{
		select{
			case client := <- r.join :
				// joining
				r.clients[client] = true
				r.tracer.Trace("New client joined")
			case client := <- r.leave : 
				// leaving
				delete(r.clients, client)
				close(client.send)
				r.tracer.Trace("Client left")
			case msg := <-  r.forward :
				// forward to all clients
				for client := range r.clients{
					select {
						case client.send <- msg :
							// send the message
							r.tracer.Trace("-- sent to client")
						default :
							// fail to send
							delete(r.clients, client)
							close(client.send)
							r.tracer.Trace("-- failed to send, cleaned up client")
					}
				}	
		}		
	}
}


const(
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ 
		ReadBufferSize : socketBufferSize, 
		WriteBufferSize: socketBufferSize,
}


func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request){
	socket, err := upgrader.Upgrade(w, req, nil) // upgrade the connection =)
	if err != nil{
		log.Fatal("ServeHTTP:", err)
		return 
	}

	client := &Client{
		socket : socket,
		send : make(chan []byte, messageBufferSize),
		room : r,
	}
	r.join <- client
	defer func(){ r.leave <- client}()
	go client.write()
	client.read()
}
