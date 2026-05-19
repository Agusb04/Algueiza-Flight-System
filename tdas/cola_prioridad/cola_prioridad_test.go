package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	TDAheap "algueiza-flight-system/tdas/cola_prioridad"
	"testing"
)

const (
	ERROR_VER_TOPE   = "Debería hacer pánico al intentar ver el tope de una cola vacía"
	ERROR_DESENCOLAR = "Debería entrar en pánico al intentar Desencolar una cola vacía"
	ERROR_ENCOLAR    = "La cola no deberia dar error al encolar"
	ERROR_ESTA_VACIA = "La cola no debería estar vacía después de encolar."
	VER_PRIMERO      = "El primer elemento debe considir con el apilado"
	COLA_VACIA       = "La cola debe estar vacia"
	ARR_ORD          = "El arreglo debe quedar ordenado"
	VOLUMEN_ALTO     = 10000
)

// FUNCION COMPARACION
func cmpEnteros(a, b int) int {
	return a - b
}

// 1
func TestColaVacia(t *testing.T) {
	cola := TDAheap.CrearHeap(cmpEnteros)

	require.True(t, cola.EstaVacia(), COLA_VACIA)
	require.Equal(t, 0, cola.Cantidad())
	require.Panics(t, func() { cola.VerMax() }, ERROR_VER_TOPE)
	require.Panics(t, func() { cola.Desencolar() }, ERROR_DESENCOLAR)
}

// 2
func TestEncolarUnElemento(t *testing.T) {
	cola := TDAheap.CrearHeap(cmpEnteros)
	cola.Encolar(1)
	require.Equal(t, 1, cola.Cantidad())
	require.Equal(t, 1, cola.VerMax())

	cola.Desencolar()
	require.True(t, cola.EstaVacia(), COLA_VACIA)
	require.Equal(t, 0, cola.Cantidad())
	require.Panics(t, func() { cola.VerMax() }, ERROR_VER_TOPE)
	require.Panics(t, func() { cola.Desencolar() }, ERROR_DESENCOLAR)
}

// 3
func TestEncolarVariosElementos(t *testing.T) {
	cola := TDAheap.CrearHeap(cmpEnteros)

	elementos := []int{5, 1, 8, 3, 10, 2}
	elementosOrdenados := []int{10, 8, 5, 3, 2, 1}

	for _, elem := range elementos {
		cola.Encolar(elem)
	}

	require.Equal(t, len(elementos), cola.Cantidad())
	require.Equal(t, 10, cola.VerMax())

	for _, esperado := range elementosOrdenados {
		require.False(t, cola.EstaVacia())
		require.Equal(t, esperado, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())
}

// 4
func TestEncolarElementosRepetidos(t *testing.T) {
	cola := TDAheap.CrearHeap(cmpEnteros)
	elementos := []int{5, 3, 5, 2, 3, 5}

	for _, e := range elementos {
		cola.Encolar(e)
	}

	require.Equal(t, len(elementos), cola.Cantidad())
	require.Equal(t, 5, cola.VerMax())

	for i := 0; i < 3; i++ {
		require.Equal(t, 5, cola.Desencolar())
	}
}

// 5
func TestCrearHeapArr(t *testing.T) {
	arr := []int{7, 1, 4, 8, 3}
	cola := TDAheap.CrearHeapArr(arr, cmpEnteros)

	require.Equal(t, len(arr), cola.Cantidad())
	require.Equal(t, 8, cola.VerMax())

	anterior := cola.Desencolar()
	for !cola.EstaVacia() {
		actual := cola.Desencolar()
		require.LessOrEqual(t, actual, anterior)
		anterior = actual
	}
}

// 6
func TestVolumen(t *testing.T) {
	cola := TDAheap.CrearHeap(cmpEnteros)

	require.True(t, cola.EstaVacia(), COLA_VACIA)
	require.Panics(t, func() { cola.VerMax() }, ERROR_VER_TOPE)
	require.Panics(t, func() { cola.Desencolar() }, ERROR_DESENCOLAR)

	for i := 0; i < VOLUMEN_ALTO; i++ {
		cola.Encolar(i)
	}
	require.Equal(t, VOLUMEN_ALTO, cola.Cantidad())
	require.Equal(t, VOLUMEN_ALTO-1, cola.VerMax())

	anterior := cola.Desencolar()
	for !cola.EstaVacia() {
		actual := cola.Desencolar()
		require.LessOrEqual(t, actual, anterior)
		anterior = actual
	}
}

// 7
func TestHeapSortOrdena(t *testing.T) {
	elementos := []int{4, 1, 7, 3, 2, 6}
	TDAheap.HeapSort(elementos, cmpEnteros)

	for i := 1; i < len(elementos); i++ {
		require.LessOrEqual(t, elementos[i-1], elementos[i], ARR_ORD)
	}
}
