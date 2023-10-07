package main

import (
	"html/template" //tem outra lib chamada text/template, porém, a html/template nos blinda de ataques
	"os"
)

//trabalhando com 1 único template para várias páginas

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

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	tmp := template.Must(template.New("content.html").ParseFiles(templates...))
	err := tmp.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
