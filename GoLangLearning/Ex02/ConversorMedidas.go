package main

import (
"fmt"
"os"
"strconv"
)

func main(){
	if len(os.Args) < 3{
		fmt.Println("Uso: conversor <conversor> <unidade>")
		os.Exit(1)
	}
	unidadeOrigem := os.Args[len(os.Args) -1]
	valoresOrigem := os.Args[1 : len(os.Args) -1]

	var unidadeDestino string
	

	if unidadeOrigem == "celsius"{
		unidadeDestino = "farenheint"	
	} else if unidadeOrigem == "quilometros"{
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s nao eh uma unidade conhecida!")
		os.Exit(1)
	}

	for i, v := range valoresOrigem	{
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posicao %d nao eh um numero valido!", v, i)
			os.Exit(1)
		}		
		var valorDestino float64
	
		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem * 1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)		
	}
	
}
