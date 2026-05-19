package lista

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la lista no tiene elementos apilados
	//false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero coloca en la primera posicion de la lista un elemento.
	// el elemento queda en la primera posicion de la lista.
	InsertarPrimero(T)

	// InsertarUltimo coloca en la ultima posicion de la lista un elemento.
	// el elemento queda en la ultima posicion de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer valor de la lista, la lista tendra un elemento menos.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el primer valor de la lista.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo muestra el ultimo valor de la lista sin eliminarlo.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista, si la lista esta vacia,
	// devolvera 0.
	Largo() int

	// Iterar aplica una funcion visitar a cada elemento de la lista hasta que termine
	// si algun elemento es falso.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo para recorrer y modificar la lista.
	Iterador() IteradorLista[T]
}

// IteradorLista define las operaciones que un iterador externo de la lista debe implementar.
type IteradorLista[T any] interface {
	// VerActual devuelve el elemento actual en el iterador.
	// En caso de que el iterador haya terminado de recorrer la lista, entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve true si el iterador tiene un siguiente elemento.
	// En caso contrario, devuelve false.
	HaySiguiente() bool

	// Siguiente mueve el iterador al siguiente elemento.
	// Si ya no hay más elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar inserta un elemento en la posición actual del iterador.
	// Inserta antes del elemento actual.
	Insertar(T)

	// Borrar elimina el elemento actual y devuelve su valor.
	// Si no hay más elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
