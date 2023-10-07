package main

//executar este arquivo: go run main.go
//acessar url: curl localhost:8080/
//acessar url: curl localhost:8080/blog

import (
	"net/http"
)

//Toda vez que executamos http.ListenAndServe(":8080", nil), o go sobe algo que chamado de multiplexer
//Nada mais é que um componente onde taxamos as rotas nele
//nessa lina http.HandleFunc("/", BuscaCEPHandler) estamos ataxando um handler nessa rota
//mas em qual miltiplexer estou fazendo isso? No padrão do go
//vantagem: assim é mais simples
//Mas e se... Eu quiser ter mais de 1 server? Mais rotas? Um pacote externo adicionar um handler no nosso servidor sem seu saber?
//Então, pra termos mais controle dos nossos handlers podemos criar nosso próprio servermux
//outra vantagem, posso ter portas diferentes

func main() {
	//vamos atachar nossos handler no server abaixo
	mux := http.NewServeMux()

	//adicionando rotas
	//função anônima
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World!"))
	//})
	//ou normal
	mux.HandleFunc("/", HomeHandler)

	//trabalhando com vários vários HandleFunc
	mux.Handle("/blog", blog{title: "MY BLOG"})

	http.ListenAndServe(":8080", mux) //não é mais nil, agora tenho meu próprio mux
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

//também posso trabalhar dessa forma, onde tenho uma função que serve para vários HandleFunc
type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
