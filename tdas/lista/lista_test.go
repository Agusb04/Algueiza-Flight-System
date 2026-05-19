package lista_test

import (
	"github.com/stretchr/testify/require"
	TDAlista "algueiza-flight-system/tdas/lista"
	"testing"
)

const (
	ERROR_BORRAR_PRIMERO = "Debería hacer pánico al intentar borrar el primer elemento de una lista vacia"
	ERROR_VER_PRIMERO    = "Debería hacer pánico al intentar ver el primer elemento de una lista vacia"
	ERROR_VER_ULTIMO     = "Debería hacer pánico al intentar ver el ultimo elemento de una lista vacia"
	N                    = 10000
)

// --------------------------------------------------TEST LISTA------------------------------------------------------
func TestListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[any]()

	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() }, ERROR_BORRAR_PRIMERO)
	require.Panics(t, func() { lista.VerPrimero() }, ERROR_VER_PRIMERO)
	require.Panics(t, func() { lista.VerUltimo() }, ERROR_VER_ULTIMO)
}

func TestEnlistarElementos(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[string]()
	valores := []string{"a", "b", "c", "d"}

	for _, v := range valores {
		lista.InsertarUltimo(v)
	}

	require.Equal(t, len(valores), lista.Largo())
	require.Equal(t, "a", lista.VerPrimero())
	require.Equal(t, "d", lista.VerUltimo())
}

func TestVolumenLista(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()

	for i := 0; i < N; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, N, lista.Largo())

	for i := 0; i < N; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())

}

func TestEnlistarYSacarElementos(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() }, ERROR_BORRAR_PRIMERO)
	require.Panics(t, func() { lista.VerPrimero() }, ERROR_VER_PRIMERO)
	require.Panics(t, func() { lista.VerUltimo() }, ERROR_VER_ULTIMO)
}

func TestAlternarEnlistarySacarElementos(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[any]()

	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.BorrarPrimero())
	lista.InsertarUltimo("hola")
	lista.InsertarUltimo(3)
	require.Equal(t, "hola", lista.BorrarPrimero())
	lista.InsertarUltimo(4)
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 4, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

// --------------------------------------------------TEST ITERADOR---------------------------------------------------
func TestInsertarElementoPosicionInicial(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	iterador.Insertar(10)

	require.Equal(t, 10, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

	iterador.Insertar(20)

	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
}

func TestInsertarElementoPosicionFinal(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	require.False(t, iterador.HaySiguiente())

	iterador.Insertar(4)

	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	iter := lista.Iterador()
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 4, iter.VerActual())
}

func TestInsertarElementoPosicionMedio(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Insertar(5)

	require.Equal(t, 5, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 4, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
}

func TestRemoverElementoIteradorInicial(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())

	iterador := lista.Iterador()
	require.Equal(t, 1, iterador.VerActual())

	eliminado := iterador.Borrar()

	require.Equal(t, 1, eliminado)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())

	eliminado2 := iterador.Borrar()

	require.Equal(t, 2, eliminado2)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())

	eliminado3 := iterador.Borrar()

	require.Equal(t, 3, eliminado3)
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())

	eliminado4 := iterador.Borrar()

	require.Equal(t, 4, eliminado4)
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())

}

func TestRemoverUltimoElementoConIterador(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 3, lista.Largo())

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual() == 3 {
			break
		}
		iterador.Siguiente()
	}

	require.Equal(t, 3, iterador.VerActual())

	eliminado := iterador.Borrar()

	require.Equal(t, 3, eliminado)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerUltimo())

	require.False(t, iterador.HaySiguiente())
}

func TestRemoverElementoDelMedio(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	require.Equal(t, 5, lista.Largo())

	iterador := lista.Iterador()

	for i := 0; i < 2; i++ {
		iterador.Siguiente()
	}

	require.Equal(t, 3, iterador.VerActual())

	eliminado := iterador.Borrar()

	require.Equal(t, 3, eliminado)

	require.Equal(t, 4, lista.Largo())

	iteradorNuevo := lista.Iterador()
	encontrado := false
	for iteradorNuevo.HaySiguiente() {
		if iteradorNuevo.VerActual() == 3 {
			encontrado = true
			break
		}
		iteradorNuevo.Siguiente()
	}

	require.False(t, encontrado)
}

func TestCasosBorde(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	require.Panics(t, func() { iterador.VerActual() })
	require.Panics(t, func() { iterador.Siguiente() })
	require.Panics(t, func() { iterador.Borrar() })
	require.False(t, iterador.HaySiguiente())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestIteradorInterno(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(1)
	iterador.Siguiente()
	iterador.Insertar(2)
	iterador.Siguiente()
	iterador.Insertar(3)

	var resultado []int
	visita := func(elemento int) bool {
		resultado = append(resultado, elemento*2)
		return true
	}

	lista.Iterar(visita)

	esperado := []int{2, 4, 6}
	require.Equal(t, esperado, resultado, "Los elementos deben ser multiplicados por 2 correctamente")
}
