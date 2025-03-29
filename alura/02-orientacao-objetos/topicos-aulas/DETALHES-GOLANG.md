### Inicialização nil

No GO, ao atribuir um valor nil, precisamos apontar para o compilador, o tipo do lugar a ser atribuido:

```go
package main

import (
    "fmt"
)

func main() {
    var s *string = nil
    fmt.Println(s)
}
```

---

### Ponteiro (*) - o que é?

```go
var conta *ContaCorrente

conta = new(ContaCorrente)
conta.titular = "Leonardo"

fmt.Println(*conta)
// {Leonardo, 0, 0, 0}
```

Exemplo: Um edifício com vários apartamentos, representando a memória do computador, onde cada apartamento é um tipo de dado diferente. Cada apartamento possui uma caixa de correio, que simboliza os ponteiros.

Essas caixas de correio têm o mesmo tamanho, mas cada uma aponta para um apartamento específico dentro do edifício. Assim, mesmo que as caixas sejam iguais, elas direcionam para diferentes espaços de memória, onde estão armazenados os dados.

Os ponteiros, portanto, são como essas caixas de correio: eles apontam para um endereço na memória, permitindo que você acesse e manipule os dados armazenados em diferentes tipos de estruturas. Isso é útil porque, em vez de duplicar dados, você pode referenciar um único espaço de memória, economizando recursos e facilitando a gestão dos dados.
