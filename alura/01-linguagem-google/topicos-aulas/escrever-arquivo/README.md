## Ler arquivos em GO

No arquivo *to-write-file.go* mostra o funcionamento das bibliotecas, que ajudam a escrever em arquivos.

### Bibliotecas utilizadas

> OS

No uso do pacote os, podemos consultar:

> https://pkg.go.dev/os

Onde podemos tirar a duvida sobre a biblioteca, no caso, usamos para saber que FLAGS usar em uma função "OpenFile" do pacote:

```go
import (
	"os"
    "strconv"
)

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)
```

E para escrever no arquivo, ainda usando propriedades do pacote OS:

```go
	arquivo.WriteString(site + "- online: " + strconv.FormatBool(status))
```

---

> STRCONV

Para converter tipos, para strings

```go
import (
    "strconv"    
)

arquivo := true
strconv.FormatBool(arquivo)
```

---

> TIME

No uso do pacote time, podemos consultar: 

https://pkg.go.dev/time#pkg-examples

Onde podemos tirar a duvida sobre a biblioteca, no caso, usamos para saber como formatar data

https://go.dev/src/time/format.go

no link acima, notamos que existe um monte de constante representados por números que podemos utilizar para formatar, no caso se baseando no link acima nosso codigo fica assim:

```go
arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05"))
```

> Cada número representa um tipo de formatação
