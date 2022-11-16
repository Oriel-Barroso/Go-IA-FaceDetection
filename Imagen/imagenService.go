package Imagen

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func openImage(nuevaImagen string) (image.Image, error) {

	file, err := os.Open("./iloveimg-resized/FotoAlegre-Emiliano-59162-modified.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "flower.jpg", err)
		return img, err
	}
	return img, nil
}
