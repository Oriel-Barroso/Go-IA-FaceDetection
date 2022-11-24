package Imagen

import (
	"image/jpeg"
	"log"
	"os"
)

type JpegFile struct {
	next Executor
}

func (jF *JpegFile) method(imagenStruct *ImagenStructDTO) {
	if imagenStruct.Tipo == ".jepg" {
		img, err := jpeg.Decode(imagenStruct.Archivo)
		if err != nil {
			log.Fatal(os.Stderr, "%s: %v\n", "no se pudo leer la imagen", err)
		}
		imagenStruct.SetImagen(img)
	} else {
		jF.next.method(imagenStruct)
	}
}

func (jF *JpegFile) setNext(next Executor) {
	jF.next = next
}
