package main

import "fmt"

func main(){
	// Declaracao
	vazio1 := map[int]string{}
	vazio2 := make(map[int]string)
	mapaGrande := make(map[int]string, 4096)
	
	capitais := map[string]string{
		"SP" : "Sao Paulo"
		"RJ" : "Rio de Janeiro"
	}

	fmt.Println(len(capitais))
	capitais["RN"] = "Natal"
	capitais["AM"] = "Manaus"
	capitais["SE"] = "Aracaju"
	fmt.Println(capitais)
	
}
