package main

import (
	"fmt"
	"os"
)

// Usando valores de retorno nomeados 
func PrecoFinal(precoCusto float64) (precoDolar, precoReal float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34
	precoDolar = precoCusto * fatorLucro // ja foi declarada na assinatura, nao precisa do ':=' 
	precoReal = precoDolar * taxaConversao
	return		// Como declaramos na assinatura as variaveis que iremos retornar
}

// Exemplo de varadic function
func CriarArquivos(dirBase string, arquivos ...string){
	
	for _, nome := range arquivos{

		caminhoArquivo := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt")
		arq, err := os.Create(caminhoArquivo)

		defer arq.Close()

		if err != nil{
			fmt.Printf("Erro ao criar arquivo %s: %v \n", nome, err)
			os.Exit(1)
		}
		fmt.Printf("Arquivo %s criado.\n", arq.Name())
	}
}

func main(){
	precoDolar, precoReal := PrecoFinal(34.99)
	fmt.Printf("Preço final em dólar: %.2f\n" + "Preço final em reais: %.2f\n", precoDolar, precoReal)

	tmp := os.TempDir()
	CriarArquivos(tmp)
	CriarArquivos(tmp, "teste1")
	CriarArquivos(tmp, "teste2", "teste3", "teste4")
}
