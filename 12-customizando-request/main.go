package main

import (
	"io/ioutil"
	"net/http"
)

//ese formato é muito usado quando quero customizar minha request. Ex: passar um bearer token....

func main() {

	c := http.Client{} //client

	req, err := http.NewRequest("GET", "http://google.com", nil) //objeto de request
	if err != nil {
		panic(err)
	}

	//adicionando um header na minha request
	req.Header.Set("Accept", "application/json") //o retorno eu aceito um json

	//meu client faz a request
	resp, err := c.Do(req)
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
