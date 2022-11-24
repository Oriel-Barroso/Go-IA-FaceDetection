package Imagen

import (
	"gorm.io/gorm"
	"image"
	"os"
)

type ImagenStructDTO struct {
	gorm.Model
	Archivo    *os.File    `json:"archivo"`
	Tipo       string      `json:"tipo"`
	ImagenFile image.Image `json:"imagenFile"`
}

func (t *ImagenStructDTO) Imagen() image.Image {
	return t.ImagenFile
}

func (t *ImagenStructDTO) SetImagen(imagen image.Image) {
	t.ImagenFile = imagen
}
