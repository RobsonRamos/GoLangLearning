package main

import "fmt"

func main(){
	var a []int // nao precisamos declarar o tamanho
	primos := []int{2, 3, 5, 7, 11, 13}
	nomes := []string{}

	fmt.Println(a, primos, nomes)
	
	b:= make([]int, 10)
	fmt.Println(b, len(b), cap(b))
	
	c := make([]int, 10, 20)
	fmt.Println(c, len(c), cap(c))

	d, e := 0, 10

	for d < e { 
		d += 1
	}

	fmt.Println(d)
	fmt.Println("____________________________________________________________")
	var i int

	for i = 0; i< 10; i++{
		fmt.Println("Valor de i", i)
	}
	fmt.Println("____________________________________________________________")

	var k int
	
	for k = 0; k < 10; {
		fmt.Println("Valor de k: ", k)	
		k += 1		
	}

	numeros := []int{1, 2, 3, 4, 5}

	for i := range numeros{
		numeros[i] *= 2 
	}
	fmt.Println("____________________________________________________________")
	for indice, valor := range numeros{
		fmt.Printf("indice[%d] = %d \n", indice, valor)
	}


	fmt.Println(numeros)

	var value int

	externo:
	for{
		for value = 0; value < 10; value++ {
			
			if value == 5 {
				break externo			
			}			
		}
	}
	fmt.Println("Valor de value : ", value)
}