package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	curso0 := Curso{"Fundamentals", 10}
	curso1 := Curso{"Go", 40}
	curso2 := Curso{"Python", 80}
	curso3 := Curso{"Java", 1000}
	cursos := Cursos{}
	cursos = append(cursos, curso0)
	cursos = append(cursos, curso1)
	cursos = append(cursos, curso2)
	cursos = append(cursos, curso3)

	//mantendo os valores do parse em um arquivo externo
	tmp := template.Must(template.New("template.html").ParseFiles("template.html")) //posso passar vários arquivos dentro de um slice
	err := tmp.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
