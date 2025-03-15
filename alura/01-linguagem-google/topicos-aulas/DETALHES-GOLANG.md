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

**Loop**

O GOLANG é uma linguagem de programação que **NÃO POSSUI WHILE** , no caso de loop, se utiliza o **FOR**