package main

import "net/http"

//criar um servidor http para acessar o via cep
//executar este arquivo: go run main.go
//acessar url: curl localhost:8080 (bad request)
//acessar url: curl localhost:8080/?cep=86060660 (ok)

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

	//adicionando dados no header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Hello, World!"))
}
