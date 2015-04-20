package main

import "fmt"


type Operacao interface{
	Calcular() int
}

type Soma struct{
	operando1, operando2 int
}

func (s Soma) Calcular() int{
	return s.operando1 + s.operando2
}

func (s Soma) String() string{
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}

type Subtracao struct{
	operando1, operando2 int
}

func (s Subtracao) Calcular() int{
	return s.operando1 - s.operando2
}

func (s Subtracao) String() string{
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2)
}

func main(){
	var soma Operacao
	soma = Soma{10, 20}
	fmt.Printf("%v = %d\n", soma, soma.Calcular())

	var subtracao Operacao
	subtracao = Subtracao{30, 15}
	fmt.Printf("%v = %d\n", subtracao, subtracao.Calcular())


	operacoes := make([]Operacao, 4)
	operacoes[0] = Soma{10, 20}
	operacoes[1] = Subtracao{15, 15}
	operacoes[2] = Soma{50, -10}
	operacoes[3] = Subtracao{20, 40}
	
	acumulador := 0

	for _, x := range operacoes{
		result := x.Calcular() 
		fmt.Printf("%v = %d\n", x, result)
		acumulador += result
	}

	fmt.Println("Valor acumulado =", acumulador)
}
