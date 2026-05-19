package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	CAPACIDAD_INICIAL  int = 10
	TAM_REDIMENSION    int = 2
	FACTOR_REDIMENSION int = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {

	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(dato T) {

	if pila.cantidad == len(pila.datos) {
		redimensionar(pila, (len(pila.datos) * TAM_REDIMENSION))
	}
	pila.datos[pila.cantidad] = dato
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {

	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	var valorDesapilado T = pila.VerTope()

	pila.cantidad--

	if pila.cantidad*FACTOR_REDIMENSION <= len(pila.datos) && len(pila.datos) > CAPACIDAD_INICIAL {
		redimensionar(pila, (len(pila.datos) / TAM_REDIMENSION))
	}

	return valorDesapilado
}

func redimensionar[T any](pila *pilaDinamica[T], capacidad int) {

	nuevosDatos := make([]T, capacidad)

	copy(nuevosDatos, pila.datos)

	pila.datos = nuevosDatos

}
