package main

import "fmt"

func separar(nums []int, i, p chan <- int, pronto chan <- bool){
	for _, n := range nums {
		if n % 2 == 0{
			p <- n
		} else {
			i <- n
		} 
	}
	pronto <- true
}

func main(){
	i, p := make(chan int), make(chan int)
	pronto := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}
	go separar(nums , i, p, pronto)
	fim := false
	for !fim {
		select {
			case v1 := <- p 	: fmt.Println("Par : ", v1)
			case v2 := <- i 	: fmt.Println("Impar : ", v2)
			case fim = <- pronto	: fmt.Println("Terminou")
		}
	}	
}
