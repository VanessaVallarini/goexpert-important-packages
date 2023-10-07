package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//criar um servidor http para acessar o via cep
//executar este arquivo: go run main.go
//acessar url: curl localhost:8080 (bad request)
//acessar url: curl localhost:8080/?cep=86060660 (ok)
//acessar url: usar o Thunder client localhost:8080/?cep=8606066

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

func main() {
	//criando a url para nosso servidor
	//poderia usar função anônima, mas só complica
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	//w.Write([]byte("Hello, World!"))
	//})
	http.HandleFunc("/", BuscaCEPHandler)

	http.ListenAndServe(":8080", nil)
}

//toda request que eu receber os dados da request chegarão aqui
//e toda response tb será entregue aqui
func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	//tratamento para quando a requisição veio errada (no header)
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//validando se veio o param cep
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//adicionando dados no header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//w.Write([]byte("Hello, World!"))
	//result, err := json.Marshal(cep)
	//if err != nil {
	//w.WriteHeader(http.StatusInternalServerError)
	//return
	//}
	//w.Write(result)

	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCep, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%v/json", cep)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var c ViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
