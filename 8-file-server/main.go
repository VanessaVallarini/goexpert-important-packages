package main

//executar este arquivo: go run main.go

import (
	"log"
	"net/http"
)

//como poder criar um servidor de arquivos? (neste exemplo uma página simples html)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog"))
	})

	//diferente de antes estamos usando um log quando der erro
	log.Fatal(http.ListenAndServe(":8080", mux))

}
