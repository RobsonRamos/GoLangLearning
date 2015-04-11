package main

import "fmt"

func main(){
	var a [3]int
	numeros := [5]int{1, 2, 3, 4, 5}
	primos := [...]int{2, 3, 5, 7, 11}
	nomes := [2]string{}

	fmt.Println(a, numeros, primos, nomes)

	multiB := [2][2]int{{2, 13}, {-1, 6}}
	fmt.Println(multiB)
}


