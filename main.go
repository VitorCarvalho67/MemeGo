package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// artASCII
	fmt.Println("\n\nMemeGo v0.1\n\n")

	// - meme com uma imagem e um texto
	// - meme com uma imagem e dois textos
	// - meme com duas imagens e um texto
	// - meme com duas imagens e dois textos

	var filepath string

	fmt.Println("1. Load Image (file path jpg, png)")
	fmt.Scanln(&filepath)

	err := copyImage(filepath, "imgs")
	if err != nil {
		fmt.Println("\nError copying image:", err)
	} else {
		fmt.Println("\nImage copied successfully.")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	var nameImage string

	sp := strings.Split(filepath, "/")
	nameImage = sp[len(sp)-1]

	img, err := loadImage("imgs/" + nameImage)

	var textImage string

	fmt.Println("\n2. Add Text")
	fmt.Scanln(&textImage)

	// tipo da fonte
	fmt.Println("\n3. Font Type")
	fmt.Println("1. Arial")
	fmt.Println("2. Comic Sans")
	fmt.Println("3. Courier")

	var fontType int
	fmt.Scanln(&fontType)

	img, err = addTextToImage(img, textImage)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n4. Saving Image...")
	err = saveImage(img, "output.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")
}

func copyImage(srcPath, destDir string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	destPath := filepath.Join(destDir, filepath.Base(srcPath))
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func addTextToImage(img image.Image, text string) (image.Image, error) {
	// Carregue uma fonte (neste exemplo, usamos uma fonte padrão do sistema)
	face := truetype.NewFace(fonts, &truetype.Options{Size: 24})

	// Crie uma nova imagem com base na imagem original
	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{}, draw.Src)

	// Defina as cores do texto e do contorno
	textColor := color.RGBA{255, 255, 255, 255} // Branco
	outlineColor := color.RGBA{0, 0, 0, 255}    // Preto

	// Adicione o texto superior
	addLabel(newImg, face, text, 10, 30, textColor, outlineColor)

	return newImg, nil
}

func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}
