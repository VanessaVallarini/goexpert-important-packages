package main

import (
	"context"
	"io/ioutil"
	"net/http"
)

//exemplo: tenho uma requisição que vai demorar um pouco
//porém, enquanto a requisição está sendo realizada o meu sistema está fazendo uma conta
//e dependendo do resultado da conta eu nem vá precisar mais do resultado da requisição
//então, temos opções:
//1 - mesmo não precisando mais do resultado da requisição eu finalizo ela
//2 - ou, cancela a requisição

func main() {
	//criando um contexto com um timeout. Se passar de 1 segundo cancela
	ctx := context.Background()

	//uma das formas de cancelar o contexto
	//ctx, cancel := context.WithTimeout(ctx, time.Second) //ISSO É SUPER ÚTIL PARA CONSULTAS NO BANCO DE DADOS

	//no fim das contas, tudo isso serve para não precisarmos setar o timeout no client:
	//c := http.Client{Timeout: time.Microsecond}

	ctx, cancel := context.WithCancel(ctx)

	//outra forma de cancelar o contexto (boa prática) é executando a função cancel():
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req) //o DefaultClient elimina a necessidade de fazer antes o client := http.Client{}
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
