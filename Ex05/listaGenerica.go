package main

import "fmt"

type ListaGenerica []interface{ }

func (lista *ListaGenerica) RemoverIndice (indice int){
	l := *lista
	removido := l[indice]
	*lista = append(l[0 : indice], l[indice +1 :]...)
	return removido
}

func main(){
	
}