package main

import "fmt"

func main(){
	fib := []int{ 0, 1, 2, 3 ,4, 5}

	fmt.Println(fib)
	fmt.Println(fib[:3])
	fmt.Println(fib[:2])
	fmt.Println(fib[2:])
	fmt.Println(fib[:len(fib)])

	novo := fib[1:3]
	fmt.Println("Novo : ", novo)

	fib[2] = 13
	fmt.Println(fib)
	fmt.Println(novo)

	// Append adiciona ao final do slice
	s := make([]int, 0)
	s = append(s, 23)

	// Adicionando elemento no inicio do slice
	n:= []int{ 22 }
	n = append(n, s...)
	
	fmt.Println(n)

	// Podemos fazer assim tb
	x := append([]int{ 22 }, s...)
	fmt.Println(x)
}
