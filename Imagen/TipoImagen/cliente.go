package Imagen

import (
	"os"
)

func RunCliente(archivoImg *os.File, tipoImg string) ImagenStructDTO {
	jpegF := &JpegFile{}
	pngF := &PngFile{}
	nothingF := &Nothing{}
	tipo := ImagenStructDTO{
		Archivo: archivoImg,
		Tipo:    tipoImg,
	}
	pngF.setNext(nothingF)
	jpegF.setNext(pngF)
	jpegF.method(&tipo)
	return tipo
}
