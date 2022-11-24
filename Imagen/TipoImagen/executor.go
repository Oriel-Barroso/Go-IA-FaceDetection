package Imagen

type Executor interface {
	method(*ImagenStructDTO)
	setNext(Executor)
}
