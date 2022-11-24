package Imagen

import (
	"image/png"
	"log"
	"os"
)

type PngFile struct {
	next Executor
}

func (pF *PngFile) method(imagenStruct *ImagenStructDTO) {
	if imagenStruct.Tipo == ".png" {
		img, err := png.Decode(imagenStruct.Archivo)
		if err != nil {
			log.Fatal(os.Stderr, "%s: %v\n", "no se pudo leer la imagen", err)
		}
		imagenStruct.SetImagen(img)
	} else {
		pF.next.method(imagenStruct)
	}
}

func (pF *PngFile) setNext(next Executor) {
	pF.next = next
}
