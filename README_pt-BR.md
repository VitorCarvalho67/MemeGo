<h4 align="center">
    <p>
        <b>English</b> |
        <a href="#">English</a>
    </p>
</h4>

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://i.imgur.com/bQfxNUL.png">
    <source media="(prefers-color-scheme: light)" srcset="https://i.imgur.com/5R19Yv9.png">
    <img alt="MemeGo image logo" src="https://i.imgur.com/bQfxNUL.png" style="width: 600px;">
  </picture>
</p>

# MemeGo

[![GitHub license](https://img.shields.io/github/license/vitorcarvalho67/MemeGo)](vitorcarvalho67/MemeGo/blob/master/LICENSE)
![GitHub stars](https://img.shields.io/github/stars/vitorcarvalho67/MemeGo)
![GitHub languages top](https://img.shields.io/github/languages/top/vitorcarvalho67/MemeGo)
![GitHub last commit](https://img.shields.io/github/last-commit/vitorcarvalho67/MemeGo)


O software MemeGo v0.1 é uma ferramenta de criação de memes desenvolvida na linguagem de programação Go. Com uma interface de linha de comando simples e intuitiva, os usuários podem carregar imagens nos formatos JPG e PNG, adicionar texto personalizado e escolher entre diferentes fontes, incluindo Arial, Comic Sans e Courier New, para dar um toque especial ao meme. O programa também permite que os usuários salvem a imagem final com o texto sobreposto em uma pasta específica, facilitando o compartilhamento em redes sociais ou com amigos. MemeGo v0.1 é uma ótima opção para quem procura uma maneira rápida e fácil de criar memes personalizados sem a necessidade de softwares de edição de imagem mais complexos.


## Como usar

clone o repositório
```bash
git clone https://github.com/VitorCarvalho67/MemeGo.git
```

Navegue até o diretório do projeto
```bash
cd MemeGo
```

Para criar um meme, execute o seguinte comando
```bash
go run main.go
```

- image: Caminho para o arquivo de imagem (formato JPG ou PNG).
- text: O texto que você deseja adicionar ao meme.
- font: O estilo da fonte (1 para Arial, 2 para Comic Sans, 3 para Courier New).

>[!IMPORTANT]
> A imagem final do meme será salva na pasta de saída especificada.


## Contribuindo
Contribuições para este projeto são bem-vindas. Siga estas etapas para contribuir:

1. Fork o repositório.
2. Crie um novo branch para sua funcionalidade ou correção de bug.
3. Faça commit das suas alterações.
4. Faça push para o branch.
5. Envie uma solicitação pull.
