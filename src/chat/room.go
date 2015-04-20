package main

type room struct{
	forward 	chan []byte
	join 		chan *Client
	leave 		chan *Client
	clients		map[*Client]bool
}
