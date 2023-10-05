package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

//executar este arquivo com: go run main.go 86060660 ...ceps
//criar o arquivo binário cep: go build -o cep main.go
//executar após criar arquivo binário ./cep 86060660
//ler o arquivo: cat cidade.txt

func main() {
	//loop para pegar tudo o que eu digitar
	for _, cep := range os.Args[1:] {
		url := fmt.Sprintf("http://viacep.com.br/ws/%v/json", cep)
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(data)

		//cereja do bolo... gravar em um arquivo
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(
			fmt.Sprintf("CEP: %s, Uf: %s,Bairro: %s,Complemento: %s,Logradouro: %s,Localidade: %s, Ddd: %s",
				data.Cep, data.Uf, data.Bairro, data.Complemento, data.Logradouro, data.Localidade, data.Ddd))
		fmt.Println("Arquivo criado com sucesso!")

	}

}
