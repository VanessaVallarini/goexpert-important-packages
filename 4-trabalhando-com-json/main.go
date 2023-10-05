package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//serialização e desserialização de json

type Conta struct {
	Numero int
	Saldo  int
}

//quando recebo uma chave diferente da minha struct uso as tags
//se eu quisesse omitir um campo do json basta usar a tag: `json:"-"`
//exemplo de tag para validar saldo == 0: `json:"s" validate:"gt=0"`
type Conta2 struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta) //guardo o valor pra mim
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

	//entregando o retorno pra alguém (api, arquivo...)
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println(err)
	}

	//agora vamos fazer ao contrário, ler um json e armazenar em uma struct
	jsonPuro := []byte(`{"Numero":1,"Saldo":100}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaX)

	//quando recebo uma chave diferente da minha struct
	jsonPuro = []byte(`{"n":1,"s":100}`)
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaX)
}
