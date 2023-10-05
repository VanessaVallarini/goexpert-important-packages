package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//WriteString grava strings
	//tamanho, err := f.WriteString("Hello, World!")

	//Write bytes, se eu não sei se a entrada será string uso esse
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
	f.Close()
	fmt.Println()

	//leitura
	//arquivo, err := os.Open("arquivo.txt")
	//facilitando a vida com Readfile
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))
	fmt.Println()

	//supondo que preciso ler um arquivo que é maior que a memória que eu tenho?
	//lendo pedaços do arquivo e carregando esses peddaços de pouco em pouco na memória
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10) //vai ler de 10 em 10 bytes
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break //se der pau para de ler o arquivo
		}
		fmt.Println(string(buffer[:n])) //vai juntando o resultado dos 10 bytes em um slice. n é a posição da leitura
	}

	//removendo um arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
