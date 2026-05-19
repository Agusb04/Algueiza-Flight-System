package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := new(nodoCola[T])
	nuevoNodo.dato = dato

	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
		c.ultimo = nuevoNodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {

	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	dato := c.primero.dato

	if c.primero == c.ultimo {
		c.primero = nil
		c.ultimo = nil

	} else {
		c.primero = c.primero.prox
	}

	return dato
}
