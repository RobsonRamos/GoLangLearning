package main

import "fmt"

func main(){
	c := make(chan int, 4)

	go produzir(c)
	consumir(c)
}

func produzir(c chan <- int) {
	c <- 1 
	c <- 2
	c <- 3
	close(c)
}

func consumir(c <- chan int){
	for valor := range c {
		fmt.Println(valor)
	}
}
