package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	dic "algueiza-flight-system/tdas/diccionario"
	"testing"
)

const (
	ErrPertenece   = "La clave deberia pertenecer"
	ErrNoPertenece = "La clave no deberia pertenecer"
	ErrBorrado     = "La clave borrada deberia devolver su valor asociado correctamente"
	ErrDicVacio    = "El diccionario deberia estar vacia"
)

func cmpInt(a, b int) int {
	return a - b
}

func TestABBVacio(t *testing.T) {
	abb := dic.CrearABB[int, string](cmpInt)

	require.Equal(t, 0, abb.Cantidad(), "El diccionario debería estar vacío al inicio.")

	require.Panics(t, func() {
		abb.Obtener(1)
	}, "Debería generar un pánico porque la clave no existe en el diccionario.")
	require.False(t, abb.Pertenece(1), "La clave 1 no debería pertenecer al diccionario.")
	abb.Guardar(1, "Uno")
	require.Equal(t, 1, abb.Cantidad(), "El diccionario debería contener 1 elemento después de insertar.")
	require.True(t, abb.Pertenece(1), "La clave 1 debería pertenecer al diccionario.")
	require.Equal(t, "Uno", abb.Obtener(1), "El valor de la clave 1 debería ser 'Uno'.")
	borrado := abb.Borrar(1)
	require.Equal(t, "Uno", borrado, "El valor borrado debería ser 'Uno'.")
	require.Equal(t, 0, abb.Cantidad(), "El diccionario debería estar vacío después de borrar el elemento.")
	require.False(t, abb.Pertenece(1), "La clave 1 no debería pertenecer al diccionario después de borrarla.")
}

func TestAgregarElementosABB(t *testing.T) {

	abb := dic.CrearABB[int, string](cmpInt)

	abb.Guardar(1, "Uno")
	abb.Guardar(25, "Dos")
	abb.Guardar(17, "Tres")
	abb.Guardar(13, "Cuatro")
	abb.Guardar(11, "Cinco")

	require.Equal(t, 5, abb.Cantidad(), "El diccionario deberia contener 5 elementos")

	require.Equal(t, true, abb.Pertenece(1), ErrPertenece)
	require.Equal(t, true, abb.Pertenece(25), ErrPertenece)
	require.Equal(t, true, abb.Pertenece(17), ErrPertenece)
	require.Equal(t, true, abb.Pertenece(13), ErrPertenece)
	require.Equal(t, true, abb.Pertenece(11), ErrPertenece)
}

func TestBorrarABB(t *testing.T) {
	abb := dic.CrearABB[int, string](cmpInt)

	abb.Guardar(12, "Veinte")
	abb.Guardar(15, "Cuarenta")
	abb.Guardar(22, "Cincuenta")

	require.Equal(t, 3, abb.Cantidad(), "El diccionario deberia contener 3 elementos")

	require.Equal(t, "Veinte", abb.Borrar(12), "El valor borrado debería ser 'Veinte'.")
	require.Equal(t, "Cuarenta", abb.Borrar(15), "El valor borrado debería ser 'Cuarenta'.")
	require.Equal(t, "Cincuenta", abb.Borrar(22), "El valor borrado debería ser 'Cincuenta'.")

	require.Equal(t, false, abb.Pertenece(12), ErrNoPertenece)
	require.Equal(t, false, abb.Pertenece(15), ErrNoPertenece)
	require.Equal(t, false, abb.Pertenece(22), ErrNoPertenece)

}

func TestAgregarYBorrarABB(t *testing.T) {

	abb := dic.CrearABB[int, float64](cmpInt)

	abb.Guardar(20, 12.3)
	abb.Guardar(13, 12.5)
	abb.Guardar(7, 12.7)
	abb.Guardar(6, 13.2)
	abb.Guardar(4, 22.2)
	abb.Guardar(12, 32.1)

	require.Equal(t, 6, abb.Cantidad(), "El diccionario deberia contener 6 elementos")

	require.Equal(t, 12.3, abb.Borrar(20), ErrBorrado)
	require.Equal(t, 12.5, abb.Borrar(13), ErrBorrado)
	require.Equal(t, 12.7, abb.Borrar(7), ErrBorrado)
	require.Equal(t, 13.2, abb.Borrar(6), ErrBorrado)
	require.Equal(t, 22.2, abb.Borrar(4), ErrBorrado)
	require.Equal(t, 32.1, abb.Borrar(12), ErrBorrado)

	require.Equal(t, 0, abb.Cantidad(), ErrDicVacio)

}

func TestPerteneceABB(t *testing.T) {
	abb := dic.CrearABB[string, int](strings.Compare)

	abb.Guardar("Clave uno", 12)
	abb.Guardar("Clave dos", 25)
	abb.Guardar("Clave tres", 22)
	abb.Guardar("Clave cuatro", 32)

	require.Equal(t, true, abb.Pertenece("Clave uno"), ErrPertenece)
	require.Equal(t, true, abb.Pertenece("Clave dos"), ErrPertenece)
	require.Equal(t, true, abb.Pertenece("Clave tres"), ErrPertenece)
	require.Equal(t, true, abb.Pertenece("Clave cuatro"), ErrPertenece)
}

func TestActualizarClaveExistente(t *testing.T) {
	abb := dic.CrearABB[int, string](cmpInt)
	abb.Guardar(10, "Diez")
	require.Equal(t, "Diez", abb.Obtener(10))
	abb.Guardar(10, "Messi")
	require.Equal(t, "Messi", abb.Obtener(10))
	require.Equal(t, 1, abb.Cantidad())
}

func TestDistintosTiposABB(t *testing.T) {
	abb1 := dic.CrearABB[float64, string](func(a, b float64) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	abb1.Guardar(3.14, "pi")
	abb1.Guardar(2.71, "e")
	require.True(t, abb1.Pertenece(3.14))
	require.Equal(t, "pi", abb1.Obtener(3.14))
	require.Equal(t, "e", abb1.Borrar(2.71))

	abb2 := dic.CrearABB[string, bool](strings.Compare)
	abb2.Guardar("a", true)
	abb2.Guardar("b", false)
	require.True(t, abb2.Pertenece("a"))
	require.Equal(t, false, abb2.Obtener("b"))
}
func TestVolumenABB(t *testing.T) {
	const tam = 10000
	abb := dic.CrearABB[int, int](cmpInt)

	claves := rand.Perm(tam)

	for _, k := range claves {
		abb.Guardar(k, k*10)
	}
	require.Equal(t, tam, abb.Cantidad())

	for _, k := range claves {
		require.True(t, abb.Pertenece(k))
		require.Equal(t, k*10, abb.Obtener(k))
	}

	for _, k := range claves {
		require.Equal(t, k*10, abb.Borrar(k))
	}
	require.Equal(t, 0, abb.Cantidad())
}

func TestIterarRangoCorte(t *testing.T) {
	abb := dic.CrearABB[int, string](cmpInt)
	for _, k := range []int{1, 2, 3, 4, 5} {
		abb.Guardar(k, fmt.Sprintf("Valor %d", k))
	}
	var claves []int
	abb.IterarRango(nil, nil, func(clave int, _ string) bool {
		claves = append(claves, clave)
		return clave < 3
	})
	require.Equal(t, []int{1, 2, 3}, claves, "La iteración debió cortarse al llegar a 3")
}

func TestIteradorRango(t *testing.T) {
	abb := dic.CrearABB[int, string](cmpInt)
	for _, k := range []int{5, 3, 8, 1, 4, 7, 10} {
		abb.Guardar(k, fmt.Sprintf("Valor %d", k))
	}
	desde := 3
	hasta := 8
	it := abb.IteradorRango(&desde, &hasta)

	var claves []int
	for it.HaySiguiente() {
		k, _ := it.VerActual()
		claves = append(claves, k)
		it.Siguiente()
	}
	require.Equal(t, []int{3, 4, 5, 7, 8}, claves)
}
