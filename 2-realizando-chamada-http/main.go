package main

import (
	"io"
	"net/http"
)

func main() {
	//acessar o site do google e baixar o código fonte
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close() //se eu abro um arquivo devo fechar para não vazar recursos
}
