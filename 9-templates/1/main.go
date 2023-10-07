package main

import (
	"os"
	"text/template"
)

//aplicação web com páginas dinâmicas (ex: email...)
//go tem sistema de templates imbutidos :)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso) //imprimindo na tela
	if err != nil {
		panic(err)
	}
}
