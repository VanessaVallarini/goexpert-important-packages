package main

import (
	"os"
	"strings"
	"text/template"
)

//usando funções dentro dos templates
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

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
