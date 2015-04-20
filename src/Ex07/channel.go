package main

import "fmt"

func main(){
	c := make(chan int)
	go produzir(c)
	valor := <-c // recebe o valor do canal
	fmt.Println(valor)
}

func produzir(c chan int) {
	c <- 33 // gera o valor para o canal
}
