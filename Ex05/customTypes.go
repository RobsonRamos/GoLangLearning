package main

import "fmt"

type ListaDeCompras []string


func main(){

	lista := make(ListaDeCompras, 6)
	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Azeite"
	lista[3] = "Atum"
	lista[4] = "Frango"
	lista[5] = "Chocolate"
	fmt.Println(lista)	
}

func (lista ListaDeCompras) Categorizar()( []string, []string, []string){
	
	var vegetais, carnes, outros []string

	for _, e := range lista{
		switch e {
			case "Alface", "Pepino" : 
				vegetais = append(vegetais, e)

			case "Atum", "Frango" : 
				carnes = append(carnes, e)

			default : 
				outros = append(outros, e)
		}
	}	
	return vegetais, carnes, outros
}
