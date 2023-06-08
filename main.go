package main

import (
	"fmt"

	"github.com/renanlido/bank/contas"
)

func main() {
	conta := contas.Constructor("Renan", 123, 123456, 1000)

	conta.Sacar(200)

	fmt.Println(conta.Saldo)
}
