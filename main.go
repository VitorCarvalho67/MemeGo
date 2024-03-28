package main

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// artASCII
	fmt.Println("MemeGo v0.1\n\n")

	var filepath string

	fmt.Println("1. Load Image (file path jpg, png)")
	fmt.Scanln(&filepath)

	err := copyImage(filepath, "imgs")
	if err != nil {
		fmt.Println("Error copying image:", err)
	} else {
		fmt.Println("Image copied successfully.")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	var image image.Image

	fmt.Println("2. Add Text")
	image, err = addTextToImage(image, "Hello, World!")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("3. Save Image")
	err = saveImage(image, "output.jpg")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done!")

}

func copyImage(srcPath, destDir string) error {
	// Abre o arquivo de origem
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Cria o diretório de destino, se necessário
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Cria o arquivo de destino
	destPath := filepath.Join(destDir, filepath.Base(srcPath))
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copia o conteúdo do arquivo de origem para o arquivo de destino
	_, err = io.Copy(destFile, srcFile)
	return err
}

func loadImage(filename string) (image.Image, error) {

	return nil, nil
}

func addTextToImage(img image.Image, text string) (image.Image, error) {

	return nil, nil
}

func saveImage(img image.Image, filename string) error {

	return nil
}
