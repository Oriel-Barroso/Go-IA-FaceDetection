package Imagen

type ImagenStruct struct {
	resultado []float64
}

func (i *ImagenStruct) Resultado() []float64 {
	return i.resultado
}

func (i *ImagenStruct) SetResultado(resultado []float64) {
	i.resultado = resultado
}
