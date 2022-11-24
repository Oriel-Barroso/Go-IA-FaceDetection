package Imagen

import (
	"github.com/Oriel-Barroso/golangBackend/Database"
	Imagen "github.com/Oriel-Barroso/golangBackend/Imagen/TipoImagen"
	"image"
	"log"
	"os"
	"regexp"
)

func saveImage(archivo *os.File, tipo string) (ImagenStruct, error) {
	newImage := ImagenStruct{Archivo: archivo, Tipo: tipo}
	db, err := Database.OpenDatabase()
	db.Create(&newImage)
	return newImage, err
}

func openImage(nuevaImagen string) (image.Image, error) {
	arc, err := os.Open(nuevaImagen)
	if err != nil {
		log.Fatal(err)
	}
	re, _ := regexp.Compile("((\\.(?i)(jpe?g|png|gif|bmp))$)")
	tipo := re.FindString(nuevaImagen)
	if tipo == "" {
		tipo = "nothing"
	}
	img := Imagen.RunCliente(arc, tipo)
	defer arc.Close()
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "no se pudo leer la imagen", err)
		return img.Imagen(), err
	}
	return img.Imagen(), nil
}
