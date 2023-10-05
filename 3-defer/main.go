package main

import (
	"io"
	"net/http"
)

//é algo que usamos para atrasar algo
func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	//assim que finalizar tudo o que tem na função main a linha abaixo será executada
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

	println()
	println("OUTRO EXEMPLO")
	println("Primeira linha")
	defer println("Segunda linha") //essa será executada por último
	println("Terceira linha")
}
