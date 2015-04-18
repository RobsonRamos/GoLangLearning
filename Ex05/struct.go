package main

import "fmt" 

type Arquivo struct{
	nome		string
	tamanho 	float64
	caracteres 	int
	palavras 	int
	linhas 		int
}

func (arquivo *Arquivo) TamanhoMedioDePalavra() float64{
	return float64(arquivo.caracteres) / float64(arquivo.palavras)
}


func (arquivo *Arquivo) MediaDePalavrasPorLinha() float64{
	return float64(arquivo.palavras) / float64(arquivo.linhas)
}

func main(){
	
	arquivo := Arquivo{"artigo.txt", 12.68, 12986, 1862, 220}
	codigoFonte := Arquivo{ nome : "programa.go", tamanho : 1.12} 
	fmt.Println(arquivo)
	fmt.Println(codigoFonte)

	fmt.Printf("Média de palavras por linha: %.2f\n", arquivo.MediaDePalavrasPorLinha())
	fmt.Printf("Tamanho médio de palavra: %.2f\n", arquivo.TamanhoMedioDePalavra())

	fmt.Printf("%s\t%.2fKB\n", arquivo.nome, arquivo.tamanho)
	fmt.Printf("%s\t%.2fKB\n",codigoFonte.nome, codigoFonte.tamanho)

	// Pointer to Struct	
	ponteiroArquivo := &Arquivo{tamanho: 7.29, nome: "arquivo.txt"}
	fmt.Printf("%s\t%.2fKB\n", ponteiroArquivo.nome, ponteiroArquivo.tamanho)

	
}
