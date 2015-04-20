package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	socket 		*websocket.Conn
	send 		chan []byte
	room 		*room
}
