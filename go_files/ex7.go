package main

import "fmt"

type Professor struct {
	nome     string
	escolas  []*Escola 
}

func (p *Professor) AdicionarEscola(escola *Escola) {
	p.escolas = append(p.escolas, escola)
}

type Escola struct {
	nome      string
	professores []*Professor 
}

func (e *Escola) AdicionarProfessor(professor *Professor) {
	e.professores = append(e.professores, professor)
	professor.AdicionarEscola(e)
}

func (p Professor) Informar() string {
	return fmt.Sprintf("Professor: %s", p.nome)
}

func (e Escola) Informar() string {
	return fmt.Sprintf("Escola: %s", e.nome)
}

func main() {

	professor1 := &Professor{nome: "Guilherme"}
	professor2 := &Professor{nome: "Eduardo"}

	escola1 := &Escola{nome: "Escola A"}
	escola2 := &Escola{nome: "Escola B"}

	escola1.AdicionarProfessor(professor1)
	escola1.AdicionarProfessor(professor2)
	escola2.AdicionarProfessor(professor1) 


	fmt.Println(escola1.Informar())
	for _, prof := range escola1.professores {
		fmt.Println(prof.Informar())
	}

	fmt.Println("\n" + escola2.Informar())
	for _, prof := range escola2.professores {
		fmt.Println(prof.Informar())
	}
}
