## Processo para instalar GOLANG:

LINK: https://go.dev/dl/

rodar no terminal o comando abaixo - para verificar se foi instalado na maquina (se for windows, e o terminal estiver aberto antes da instalação, fecha e abre e verifique novamente).

> go version

O GO ele tem algo chamado "Go Workspace" - espaço de trabalho Go - basicamente uma pasta "go" que fica localizada na **raiz do úsuario**, e contém todos os programas feito com o GOLANG estaram armazenados nela.

*   É necessário estar na **raiz do seu usuarios** , e criar a pasta 

# (Windows)

> mkdir go

*   **Configurar a pasta go**, 

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src

dentro da pasta "go", criar a pasta "pkg" - (que ficará nossos pacotes compartilhados)

> mkdir pkg

em seguida, a pasta "src" - (que ficará o "source code" - onde iremos escrever os codigos)

> mkdir src

em seguida, a pasta "bin" - (onde vai conter os binários - os compilados dos codigos GO)

> mkdir bin

# (Linux)

(o comando abaixo contém a versão 1.8.3, deve ser alterado pela versão necessária)

> sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz

* Adicionar o caminho /usr/local/go/bin no PATH do sistema

comando abaixo para abrir o GEDIT (editor de texto) direto no arquivo /etc/profile

> sudo gedit /etc/profile

Ao entrar no arquivo, adicionar a seguinte linha

> export PATH=$PATH:/usr/local/go/bin

Salve a alteração e feche o editor. Em seguida use o comando:

> source /etc/profile

E para verificar:

> go version