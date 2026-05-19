package lista

const (
	ERROR_LISTA_VACIA = "La lista esta vacia"
)

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nodoLista := &nodoLista[T]{dato: elemento, prox: l.primero}
	l.primero = nodoLista
	if l.EstaVacia() {
		l.ultimo = nodoLista
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nodoLista := &nodoLista[T]{dato: elemento, prox: nil}
	if l.EstaVacia() {
		l.primero = nodoLista
	} else {
		l.ultimo.prox = nodoLista
	}
	l.ultimo = nodoLista
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic(ERROR_LISTA_VACIA)
	}
	elementoBorrar := l.primero.dato
	l.primero = l.primero.prox
	l.largo--

	if l.primero == nil {
		l.ultimo = nil
	}

	return elementoBorrar
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic(ERROR_LISTA_VACIA)
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic(ERROR_LISTA_VACIA)
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodo := l.primero
	for nodo != nil {
		if !visitar(nodo.dato) {
			break
		}
		nodo = nodo.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{lista: l, actual: l.primero, anterior: nil}
}
