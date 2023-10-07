package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}

	//rodando o template.New e tmp.Parse em um único comando e idnetificando possíveis erros - usando template.must
	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := tmp.Execute(os.Stdout, curso) //imprimindo na tela
	if err != nil {
		panic(err)
	}
}
