package contas 

const quantoAdepositar = 200
const quantoAsacar = 100

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numDaConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "\nSaque realizado com sucesso"
	} else {
		return "\nSaldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeDeposito float64) (string, int) {
	podeSacar := valorDeDeposito > 0 && valorDeDeposito <= c.saldo
	if podeSacar {
		c.saldo += valorDeDeposito
		return "\nFoi depositado com sucesso, um valor de", quantoAdepositar
	} else {
		return "\nErro não deu para depositar", quantoAdepositar
	}
	
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
