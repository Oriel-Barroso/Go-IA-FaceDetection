package Imagen

import "os"

func main() {
	type Imagen struct {
		Archivo *os.File `json:"archivo"`
		Alto    float32  `json:"alto"`
		Ancho   float32  `json:"ancho"`
	}
}
