### 1. Inicio de tudo

Iniciar uma aplicação em go

> go mod init github.com/url-do-repositorio

Criará nosso arquivo **go.mod** - que guardará as dependencias do app e a nomenclatura do app.

Em seguida baixar um framework que auxilie para criação de rotas em um servidor local, no caso tem o framework GIN

Em seguida criar um arquivo **main.go** - onde ficará a área principal da nossa aplicação.
Ainda no main, digitando "pkgm" criamos a o pacote e a função main, importamos o gin e utilizamos o framework criando uma função básica:

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", teste)
	r.Run("8000")
}

func teste(c *gin.Engine) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
```

---

### 2. Estruturando

Claro que não podemos deixar tudo no main, pois conforme nossa aplicação cresça, temos que ter nosso codigo organizado, para isso. (A modularização traz muitas vantagens, como a reutilização do código, a facilidade de manutenção e a melhor legibilidade.)

Devemos criar uma pasta /controller/controller.go - para guardar nossas receber e retornar valores, pasta /routes/routes.go - para guardar nossas rotas, a pasta /models/models.go para criar as estruturas dos tipos de dados que retornaremos.

Ficará estruturado:

```go
package models

type User struct {
	Id string `json:"name"`
	Name  string  `json:"name"`
}
```

```go
package controllers

import (
	"github.com/repositorio/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(200, models.User)
}
```

```go
package routes

import (
	"github.com/repositorio/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/users", controllers.GetUser)
	r.Run()
}
```

```go
package main

import (
	"github.com/repositorio/models"
	"github.com/repositorio/routes"
)

func main() {
	routes.Initialize()
}
```

#### 2.1. Caracteristica

Para receber dados pelo parametros da requisição da nossa api, utilizamos este formato - com dois pontos e o nome de definição do parâmetro:

```go
package main
	r.GET("/users/:nome", controllers.GetUserByName)

package controller
	func GetUserByName(c *gin.Context) {
		name := c.Params.ByName("name")

		c.JSON(200, gin.H{		
			"name": name,
		})
	}
```

---

### 3. Docker e Banco

As aplicações reais utilizam banco de dados, e para trabalharmos com o banco desejado, rodamos o docker que criará um ecossistema para acessarmos nosso banco.

#### 3.1. Comunicação do banco a aplicação GO

Para facilitar a comunicação do banco com nossa aplicação GO, nós podemos utilizar um ORM, que é o GORM.

##### 3.1.1. Detalhes ( GORM )

Com o GORM temos a facilidade de criar uma tabela por meio do GO. 
O GORM também nos disponibiliza a criação de colunas default - id, created-at e etc. Uma forma de conseguirmos declarar o Gorm **_“embedando”_** alguns códigos:

```go
type User struct {
	gorm.Model
	Name string
}
// gorm.Model definition, EQUALS:
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
```

Guarda uma série de informações só de colocar o "gorm.Model"

> Mais informações: https://gorm.io/docs

--- 

Com o **AutoMigrate( )** do GORM conseguimos criar uma tabela com base na struct que definirmos e como o nome ja diz, fazer uma auto migração.
Os modelos são criados no banco de dados utilizando como base os dados acessados a partir da referência de memória de uma ou mais structs que são passadas como parâmetro para o método db.AutoMigrate.

```go
import (
	"github.com/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	connection := "host=host user=user password=password dbname=dbname port=port"

	DB, err := gorm.Open(postgres.Open(connection))
	DB.AutoMigrate(&model.Aluno{})
}
```

---

