package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

//requisições externas via http (busca cep)
//qd o sistema precisa ser rápido a forma de fazer isso é saber estabelecer os limites de chamadas externas
//ex: minha api é mega rápida, porém, o ViaCep demora muito pra responder, logo minha api se torna lenta
//precisamos fazer esse tipo de tratamento

func main() {
	c := http.Client{Timeout: time.Microsecond} //passando o tempo de duração da request para o objeto http
	resp, err := c.Get("http://google.com")     //executando o meu client
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
