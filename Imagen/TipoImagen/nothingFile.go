package Imagen

import (
	"fmt"
)

type Nothing struct {
	next (Executor)
}

func (nt *Nothing) method(imagenStruct *ImagenStructDTO) {
	if imagenStruct.Tipo == "nothing" {
		fmt.Errorf("Formato erroneo")
	}
}

func (nt *Nothing) setNext(next Executor) {
	nt.next = next
}
