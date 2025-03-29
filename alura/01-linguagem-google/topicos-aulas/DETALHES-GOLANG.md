Para declarar variavel, basta apenas usar **:=** , a tipagem é automatica, de acordo com o valor inserido

```go
nome := "Leonardo"
```
---

**No GO não precisa de parênteses em condições**

(usando de exemplo javascript para comparar)

```js
if (nome == "Leonardo") {

}
```

```go
if nome == "Leonardo" {

}
```

e é só aceito expressões que retorne _boolean_, exemplo

no javascript se é informado a seguinte condição no if, é denominado true ou false para a expressão

```js
if (nome) {
    console.log("Nome tem valor")
}
```

já no go, não pode, tem que ser definido a expressão:

```go
if nome == "" {
    fmt.Println("Nome tem valor")
}
```

---

É comum no go utilizar estilo : _camelCase_

```go
func exibeMenu() {
    
}
```

---

**Erro no codigo**

O "-1" significa para o sistema operacional que no codigo ocorreu um erro inesperado

```go
os.Exit(-1)
```

---

**LOOP**

O GOLANG é uma linguagem de programação que **NÃO POSSUI WHILE** :x: , no caso de loop, se utiliza o **FOR** :white_check_mark:

O Golang tem uma forma simples  de percorer uma lista com for:

Usando o **RANGE NO FOR**, um operador de iteração do Go, conseguimos acesso a cada item do array, ou do slice:

```go
slice := []string{"Lais", "Leonardo", "Vitoria"}

for i, items := range slice {
    fmt.Println("Posição ", i,
        " do meu slice e essa posição tem o nome", items)
}
```

---

**ARRAYS E SLICES**

No Golang os arrays ficam com o tamanho fixo que pré definimos, então normalmente no GO:

-   não trabalhamos com Arrays :x: 
-   e sim com Slices :white_check_mark:

```go
var lista [4]string
slice := []string{"Lais", "Leonardo", "Vitoria"}

fmt.Println("O meu slice tem", len(slice), "itens")
// print: O meu slice tem 3 itens
fmt.Println("O meu slice tem capacidade para", cap(slice), "itens")
// print: O meu slice tem capacidade para 3 itens"

nomes = append(nomes, "Aparecida")

fmt.Println("O meu slice tem", len(nomes), "itens")
// print: O meu slice tem 4 itens"
fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")    
// print: O meu slice tem capacidade para 6 itens"
```

Slices é como se fosse uma abstração do array, só que encima do array. Resumindo basicamente o slice em GO, é um array, os dois tem a mesma caracteristicas mas possuem funcionalidades e finalidades diferentes

**Outro detalhe:** Quando adicionamos um novo item, o slice dobra de tamanho. 

#### :sparkles: O slice nada mais é do que o Go cuidando do array para nós, evitando com que nos preocupemos com o tamanho e capacidade do array :sparkles:

---

##### Tirar espaçamentos no GO:

```go
import (
	"strings"
)

nome := "Leonardo    "

nome = strings.TrimSpace(nome)
```

---

##### Converter todos tipos para string:

```go
import (
    "strconv"    
)

arquivo := true
strconv.FormatBool(arquivo)
```