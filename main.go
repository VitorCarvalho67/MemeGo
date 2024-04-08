package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func main() {
	// artASCII

	color := "\033[38;5;52m"
	reset := "\033[0m"

	fmt.Println(color + "\n\n" + `
	__       __                                     ______            
	|  \     /  \                                   /      \           
	| $$\   /  $$  ______   ______ ____    ______  |  $$$$$$\  ______  
	| $$$\ /  $$$ /      \ |      \    \  /      \ | $$ __\$$ /      \ 
	| $$$$\  $$$$|  $$$$$$\| $$$$$$\$$$$\|  $$$$$$\| $$|    \|  $$$$$$\
	| $$\$$ $$ $$| $$    $$| $$ | $$ | $$| $$    $$| $$ \$$$$| $$  | $$
	| $$ \$$$| $$| $$$$$$$$| $$ | $$ | $$| $$$$$$$$| $$__| $$| $$__/ $$
	| $$  \$ | $$ \$$     \| $$ | $$ | $$ \$$     \ \$$    $$ \$$    $$
	 \$$      \$$  \$$$$$$$ \$$  \$$  \$$  \$$$$$$$  \$$$$$$   \$$$$$$
	` + "\n\n" + reset)

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

	fmt.Println("Arial (1)")
	fmt.Println("Comic Sans (2)")
	fmt.Println("Courier New (3)")

	var fontType int
	fmt.Scanln(&fontType)

	img, err = addTextToImage(img, textImage, fontType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n4. Saving Image...")
	err = saveImage(img, "results/output.jpg")
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

func loadFont(fontPath string) (font.Face, error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return face, err
}

func addLabel(img *image.RGBA, face font.Face, label string, x, y int, textColor color.Color) {
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: face,
		Dot:  point,
	}

	drawer.DrawString(label)
}

func addTextToImage(img image.Image, text string, fontfamily int) (image.Image, error) {

	var face font.Face
	var err error

	switch fontfamily {
	case 1:
		face, err = loadFont("fonts/arial.ttf")
		if err != nil {
			return nil, fmt.Errorf("failed to load Arial font: %v", err)
		}
	case 2:
		face, err = loadFont("fonts/ComicNeue-Regular.ttf")
		if err != nil {
			return nil, fmt.Errorf("failed to load Comic Sans font: %v", err)
		}
	case 3:
		face, err = loadFont("fonts/CourierPrime-Regular.ttf")
		if err != nil {
			return nil, fmt.Errorf("failed to load Courier New font: %v", err)
		}
	}

	if face == nil {
		return nil, fmt.Errorf("failed to load font")
	}

	// Crie uma nova imagem com base na imagem original com uma barra branca no topo

	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{}, draw.Src)

	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(newImg, image.Rect(0, 0, newImg.Bounds().Dx(), 50), &image.Uniform{white}, image.Point{}, draw.Src)

	textColor := color.RGBA{0, 0, 0, 255}
	addLabel(newImg, face, text, 10, 30, textColor)

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
