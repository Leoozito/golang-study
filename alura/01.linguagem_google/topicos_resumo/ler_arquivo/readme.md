## Ler arquivos em GO

No arquivo seguinte mostra bibliotecas apresentadas no curso, que ajudam a ler arquivos de forma diferentes

### Bibliotecas utilizadas

> Ioutil
Lê o arquivo inteiro 

> Bufio
Lê até aonde definirmos que é para ler

```go
	arquivo, err := os.Open("nomes.txt")

	leitor := bufio.NewReader(arquivo)

	resp, err := leitor.ReadString('s')

	fmt.Println(resp)
```
> Irá ler até a 1° letra "s" que conter no arquivo. 