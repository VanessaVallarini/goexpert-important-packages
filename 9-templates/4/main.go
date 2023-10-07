package main

import (
	"html/template"
	"net/http"
)

//fazendo o 3, porém, usando web server

type Cursos []Curso

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		curso0 := Curso{"Fundamentals", 10}
		curso1 := Curso{"Go", 40}
		curso2 := Curso{"Python", 80}
		curso3 := Curso{"Java", 1000}
		cursos := Cursos{}
		cursos = append(cursos, curso0)
		cursos = append(cursos, curso1)
		cursos = append(cursos, curso2)
		cursos = append(cursos, curso3)

		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := tmp.Execute(w, cursos) //no lugar de excecutar os.Stdout uso o w (response)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
