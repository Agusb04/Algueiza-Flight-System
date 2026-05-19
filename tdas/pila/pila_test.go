package pila_test

import (
	TDAPila "algueiza-flight-system/tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ErrPilaVacia      = "La pila esta vacia"
	ErrPilaVaciaTrue  = "Una pila vacia deberia devolver verdadero"
	ErrDesapilarVacia = "Desapilar una pila vacia deberia arrojar un panic"
	ErrVerTopeVacia   = "Una pila vacia deberia arrojar un panic"
	ErrVerTope        = "El tope deberia ser el ultimo elemento agregado"
	ErrInvarianteLIFO = "Al desapilar deberia mantener la invariante LIFO"
	ErrPilaTrue       = "Una pila con elementos deberia devolver Falso"

	VALOR_VOLUMEN int = 100001
	VALOR_BOOL_1      = true
	VALOR_BOOL_2      = false
	LongitudTest  int = 10
)

func TestPilaVacia(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, ErrPilaVacia, func() { pila.VerTope() }, ErrVerTopeVacia)
	require.PanicsWithValue(t, ErrPilaVacia, func() { pila.Desapilar() }, ErrDesapilarVacia)

	// pruebo apilando y despues desapilando
	pila.Apilar(1)
	require.False(t, pila.EstaVacia(), ErrPilaTrue)

	pila.Desapilar()
	require.True(t, pila.EstaVacia(), ErrPilaVaciaTrue)

}

func TestPilaVaciaComportaComoNueva(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	// pruebo Apilando varios
	for i := 0; i < LongitudTest; i++ {
		pila.Apilar(i)
	}

	for i := 9; i >= 0; i-- {
		pila.Desapilar()
	}

	require.True(t, pila.EstaVacia(), ErrPilaVaciaTrue)
	require.PanicsWithValue(t, ErrPilaVacia, func() { pila.VerTope() }, ErrVerTopeVacia)
	require.PanicsWithValue(t, ErrPilaVacia, func() { pila.Desapilar() }, ErrDesapilarVacia)

	// pruebo apilando
	pila.Apilar(20)

	require.Equal(t, 20, pila.VerTope(), ErrVerTope)
	require.Equal(t, 20, pila.Desapilar(), "Al desapilar deberia devolver el tope")
	require.True(t, pila.EstaVacia(), ErrPilaVaciaTrue)
}

func TestInvarianteLifo(t *testing.T) {

	// pruebo con una pila de ints
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < LongitudTest; i++ {
		pila.Apilar(i)
	}

	for i := 9; i > 0; i-- {
		require.Equal(t, i, pila.Desapilar(), ErrInvarianteLIFO)
	}

}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	volumen := make([]int, VALOR_VOLUMEN)

	for i := range volumen {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), ErrVerTope)
	}

	for i := VALOR_VOLUMEN - 1; i >= 0; i-- {

		pila.Desapilar()

		if !pila.EstaVacia() {
			require.Equal(t, i-1, pila.VerTope())
		}
	}

}

func TestVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar("b")
	pila.Apilar("c")
	pila.Apilar("d")
	pila.Apilar("b")
	pila.Apilar("e")

	require.Equal(t, "e", pila.VerTope(), ErrVerTope)
	pila.Desapilar()
	require.Equal(t, "b", pila.VerTope(), ErrVerTope)
	pila.Desapilar()
	require.Equal(t, "d", pila.VerTope(), ErrVerTope)
	pila.Desapilar()
	require.Equal(t, "c", pila.VerTope(), ErrVerTope)
	pila.Desapilar()
	require.Equal(t, "b", pila.VerTope(), ErrVerTope)
}

func TestDistintosTipos(t *testing.T) {

	// pruebo con una pila de strings
	pila2 := TDAPila.CrearPilaDinamica[string]()

	pila2.Apilar("a")
	pila2.Apilar("b")
	pila2.Apilar("c")
	pila2.Apilar("d")
	pila2.Apilar("e")

	require.Equal(t, "e", pila2.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, "d", pila2.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, "c", pila2.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, "b", pila2.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, "a", pila2.Desapilar(), ErrInvarianteLIFO)

	// pruebo con una pila de bools
	pila3 := TDAPila.CrearPilaDinamica[bool]()

	pila3.Apilar(VALOR_BOOL_1)
	pila3.Apilar(VALOR_BOOL_2)
	pila3.Apilar(VALOR_BOOL_2)
	pila3.Apilar(VALOR_BOOL_1)

	require.Equal(t, VALOR_BOOL_1, pila3.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, VALOR_BOOL_2, pila3.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, VALOR_BOOL_2, pila3.Desapilar(), ErrInvarianteLIFO)
	require.Equal(t, VALOR_BOOL_1, pila3.Desapilar(), ErrInvarianteLIFO)
}
