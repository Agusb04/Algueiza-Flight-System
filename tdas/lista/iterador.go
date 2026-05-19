package lista

const (
	FIN_ITERADOR = "El iterador termino de iterar"
)

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func nuevoNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func (i *iteradorLista[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic(FIN_ITERADOR)
	}
	return i.actual.dato
}

func (i *iteradorLista[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iteradorLista[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic(FIN_ITERADOR)
	}
	i.anterior = i.actual
	i.actual = i.actual.prox
}

func (i *iteradorLista[T]) Insertar(dato T) {

	nuevoDato := nuevoNodo(dato)
	nuevoDato.prox = i.actual

	if i.anterior == nil {
		i.lista.primero = nuevoDato
	} else {
		i.anterior.prox = nuevoDato
	}

	if i.actual == nil {
		i.lista.ultimo = nuevoDato
	}
	i.actual = nuevoDato
	i.lista.largo++
}

func (i *iteradorLista[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic(FIN_ITERADOR)
	}

	dato := i.actual.dato

	if i.anterior == nil {
		i.lista.primero = i.actual.prox
	} else {
		i.anterior.prox = i.actual.prox
	}

	if i.actual.prox == nil {
		i.lista.ultimo = i.anterior
	}

	i.actual = i.actual.prox
	i.lista.largo--

	return dato
}
