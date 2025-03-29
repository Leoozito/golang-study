#### Struct (Estrutura):

Uma estrutura de dados que permite agrupar diferentes tipos de dados sob um nome.
Um tipo que conterá vários outros tipos dentro dele.

Os elementos dentro da estrutura, ganham automaticamente um valor inicial

```go
type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

func main() {
    fmt.Println(ContaCorrente{})
    // {  0 0 0}
}
```

### Atribuir valores

```go
// 1° modo de atribuir valor á estrutura
func assignValue1() {
    conta := ContaCorrente{
        "Leonardo",
        1234,
        12345678,
        2300.00,
    }

    fmt.Println(conta)
    // {Leonardo, 1234, 12345678, 2300.00}
}

// 2° modo de atribuir valor á estrutura
func assignValue2() {
    var conta *ContaCorrente
    conta = new(ContaCorrente)
    
    conta.titular = "Leonardo"
    conta.numeroAgencia = 1234
    conta.numeroConta = 12345678
    conta.saldo = 2300.00

    fmt.Println(*conta)
    // {Leonardo, 1234, 12345678, 2300.00}
}
```

-   O asterisco (*) indica que a variável conta é um **ponteiro** para o tipo ContaCorrente.
