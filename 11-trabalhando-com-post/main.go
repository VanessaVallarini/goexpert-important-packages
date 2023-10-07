package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name: "vanessa"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) //copio o body da resposta e colo na saída
}
