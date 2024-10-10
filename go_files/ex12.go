package main

import (
	"fmt"
)

type Produto struct {
	nome  string
	preco float64
}

func (p Produto) Somar(outro Produto) Produto {
	return Produto{
		nome:  p.nome + " e " + outro.nome,
		preco: p.preco + outro.preco,
	}
}
func (p Produto) String() string {
	return fmt.Sprintf("Produto: %s, Preço: R$%.2f", p.nome, p.preco)
}

func main() {
	produto1 := Produto{"Produto1", 80.00}
	produto2 := Produto{"Produto2", 10.00}

	produtoSoma := produto1.Somar(produto2)
	fmt.Println(produtoSoma.String())
}
