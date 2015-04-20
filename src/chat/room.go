package main

import (
	"log"
	"github.com/gorilla/websocket"
	"net/http"
)

type room struct{
	forward 	chan []byte
	join 		chan *Client
	leave 		chan *Client
	clients		map[*Client]bool
}


func NewRoom() *room{

	return &room{
		forward : 	make(chan []byte),
		join : 		make(chan *Client),
		leave: 		make(chan *Client),
		clients:	make(map[*Client]bool),
	}
}


func (r *room)run() {
	for{
		select{
			case client := <- r.join :
				// joining
				r.clients[client] = true
			case client := <- r.leave : 
				// leaving
				delete(r.clients, client)
				close(client.send)
			case msg := <-  r.forward :
				// forward to all clients
				for client := range r.clients{
					select {
						case client.send <- msg :
							// send the message
						default :
							// fail to send
							delete(r.clients, client)
							close(client.send)
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
