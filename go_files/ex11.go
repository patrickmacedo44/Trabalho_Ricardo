package main

import (
	"fmt"
)

type Funcionario interface {
	calcularSalario() float64
}

type FuncionarioHorista struct {
	nome      string
	horasTrabalhadas float64
	precoPorHora    float64
}
func (f FuncionarioHorista) calcularSalario() float64 {
	return f.horasTrabalhadas * f.precoPorHora
}

type FuncionarioAssalariado struct {
	nome   string
	salarioMensal float64
}

func (f FuncionarioAssalariado) calcularSalario() float64 {
	return f.salarioMensal
}

func main() {

	funcionario1 := FuncionarioHorista{
		nome:           "Fernando",
		horasTrabalhadas: 200,
		precoPorHora:    10.0,
	}

	funcionario2 := FuncionarioAssalariado{
		nome:        "Lucas",
		salarioMensal: 3000.0,
	}

	fmt.Printf("Salário de %s (Horista): R$ %.2f\n", funcionario1.nome, funcionario1.calcularSalario())
	fmt.Printf("Salário de %s (Assalariado): R$ %.2f\n", funcionario2.nome, funcionario2.calcularSalario())
}
