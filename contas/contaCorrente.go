package contas

type contaCorrente struct {
	Titular string
	NumeroAgencia,
	NumeroConta int
	Saldo float64
}

func Constructor(titular string, numeroAgencia int, numeroConta int, saldo float64) *contaCorrente {
	return &contaCorrente{
		Titular:       titular,
		NumeroAgencia: numeroAgencia,
		NumeroConta:   numeroConta,
		Saldo:         saldo,
	}
}

func (conta *contaCorrente) GetSaldo() float64 {
	return conta.Saldo
}

func (conta *contaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= conta.Saldo {
		conta.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (conta *contaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		conta.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", conta.Saldo
	} else {
		return "Valor do depósito menor que zero", conta.Saldo
	}
}
