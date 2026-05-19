package cola_test

import (
	TDACola "algueiza-flight-system/tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ERRCOLAVACIA      string = "Una cola vacia deberia devolver true"
	ERRDESENCOLAR     string = "Al desencolar deberia devolver el primer elemento encolado"
	ERRVERPRIMERO     string = "Deberia devolver el primer dato encolado"
	ERRINVARIANTEFIFO string = "Al desencolar deberia mantener la invariante FIFO"

	VOLUMEN  int = 100000
	VOLUMEN2 int = 11
	VOLUMEN3 int = 20
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia(), ERRCOLAVACIA)

	cola.Encolar(10)
	cola.Encolar(20)
	require.Equal(t, 10, cola.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, 20, cola.Desencolar(), ERRDESENCOLAR)

	require.True(t, cola.EstaVacia(), ERRCOLAVACIA)
}

func TestColaVaciaComoNueva(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia(), ERRCOLAVACIA)
	for i := range VOLUMEN2 {
		cola.Encolar(i)
	}

	for range VOLUMEN2 {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia(), ERRCOLAVACIA)

}

func TestInvarianteFIFO(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	for i := range VOLUMEN3 {
		cola.Encolar(i)
	}

	for range VOLUMEN3 {
		require.Equal(t, cola.VerPrimero(), cola.Desencolar())
	}
}

func TestVolumen(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	for i := range VOLUMEN {
		cola.Encolar(i)
	}

	for i := range VOLUMEN {
		require.Equal(t, i, cola.VerPrimero(), ERRVERPRIMERO)
		require.Equal(t, i, cola.Desencolar(), ERRDESENCOLAR)
	}

}

func TestVerPrimero(t *testing.T) {
	colaStrings := TDACola.CrearColaEnlazada[string]()

	colaStrings.Encolar("c")
	colaStrings.Encolar("d")
	colaStrings.Encolar("e")
	require.Equal(t, "c", colaStrings.VerPrimero(), ERRVERPRIMERO)
	colaStrings.Desencolar()
	require.Equal(t, "d", colaStrings.VerPrimero(), ERRVERPRIMERO)
	colaStrings.Desencolar()
	require.Equal(t, "e", colaStrings.VerPrimero(), ERRVERPRIMERO)
	colaStrings.Desencolar()
	require.Panics(t, func() { colaStrings.VerPrimero() }, "Una cola vacia deberia arrojar un panic")
}

func TestDistintosTipos(t *testing.T) {
	colaStrings := TDACola.CrearColaEnlazada[string]()

	colaStrings.Encolar("a")
	colaStrings.Encolar("b")
	colaStrings.Encolar("c")
	colaStrings.Encolar("d")
	colaStrings.Encolar("e")

	require.Equal(t, "a", colaStrings.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, "b", colaStrings.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, "c", colaStrings.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, "d", colaStrings.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, "e", colaStrings.Desencolar(), ERRDESENCOLAR)

	colaBools := TDACola.CrearColaEnlazada[bool]()

	colaBools.Encolar(true)
	colaBools.Encolar(true)
	colaBools.Encolar(false)
	colaBools.Encolar(true)
	colaBools.Encolar(false)

	require.Equal(t, true, colaBools.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, true, colaBools.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, false, colaBools.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, true, colaBools.Desencolar(), ERRDESENCOLAR)
	require.Equal(t, false, colaBools.Desencolar(), ERRDESENCOLAR)

}
