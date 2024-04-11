<h4 align="center">
    <p>
        <b>English</b> |
        <a href="https://github.com/VitorCarvalho67/MemeGo/blob/main/README_pt-BR.md">Рortuguês</a>
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

The MemeGo v0.1 software is a meme creation tool developed in the Go programming language. With a simple and intuitive command-line interface, users can load images in JPG and PNG formats, add custom text, and choose from different fonts, including Arial, Comic Sans, and Courier New, to give a special touch to the meme. The program also allows users to save the final image with the overlaid text in a specific folder, making it easy to share on social networks or with friends. MemeGo v0.1 is a great option for those looking for a quick and easy way to create personalized memes without the need for more complex image editing software.

## How to use

clone the repository
```bash
git clone https://github.com/VitorCarvalho67/MemeGo.git
```

Navigate to the project directory
```bash
cd MemeGo
```

To create a meme, run the following command
```bash
go run main.go
```

- image: Path to the image file (JPG or PNG format).
- text: The text you want to add to the meme.
- font: The font style (1 for Arial, 2 for Comic Sans, 3 for Courier New).

>[!IMPORTANT]
> The final meme image will be saved in the specified output folder.

## Contributing
Contributions to this project are welcome. Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes.
4. Push to the branch.
5. Submit a pull request.
