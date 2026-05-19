package diccionario

import (
	"fmt"
)

type estadoEntrada uint8

const (
	vacio estadoEntrada = iota
	ocupado
	borrado
)

const (
	CAPACIDADINICIAL         = 17
	FACTORCARGAMINIMO        = 0.25
	FACTORCARGAMAXIMA        = 0.7
	ERRORKEY          string = "La clave no pertenece al diccionario"
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoEntrada
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	borrados int
}

type iteradorDiccionario[K comparable, V any] struct {
	dicc   *hashCerrado[K, V]
	actual int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{tabla: make([]celdaHash[K, V], CAPACIDADINICIAL)}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (d *hashCerrado[K, V]) buscarClave(clave K) (int, bool) {
	indice := hash(clave, len(d.tabla))
	original := indice
	primeraBorrada := -1

	for {
		celda := d.tabla[indice]

		if celda.estado == vacio {
			if primeraBorrada != -1 {
				return primeraBorrada, false
			}
			return indice, false
		}
		if celda.estado == ocupado && celda.clave == clave {
			return indice, true
		}
		if celda.estado == borrado && primeraBorrada == -1 {
			primeraBorrada = indice
		}

		indice = (indice + 1) % len(d.tabla)
		if indice == original {
			if primeraBorrada != -1 {
				return primeraBorrada, false
			}
			return -1, false
		}
	}
}

func hash[K comparable](clave K, capacidad int) int {
	bytes := convertirABytes(clave)

	const prime = 16777619
	hash := 0
	for _, b := range bytes {
		hash = (hash*prime + int(b)) % capacidad
	}

	if hash < 0 {
		hash += capacidad
	}

	return hash
}

func (d *hashCerrado[K, V]) Guardar(clave K, dato V) {
	indice, encontrado := d.buscarClave(clave)
	if encontrado {
		d.tabla[indice].dato = dato
		return
	}

	d.tabla[indice] = celdaHash[K, V]{clave: clave, dato: dato, estado: ocupado}
	d.cantidad++

	if float64(d.cantidad)/float64(len(d.tabla)) > FACTORCARGAMAXIMA {
		d.redimensionar(len(d.tabla) * 2)
	}
}

func (d *hashCerrado[K, V]) Obtener(clave K) V {
	indice, encontrado := d.buscarClave(clave)
	if !encontrado {
		panic(ERRORKEY)
	}
	return d.tabla[indice].dato
}

func (d *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, encontrado := d.buscarClave(clave)
	return encontrado
}

func (d *hashCerrado[K, V]) Borrar(clave K) V {
	indice, encontrado := d.buscarClave(clave)
	if !encontrado {
		panic(ERRORKEY)
	}

	valor := d.tabla[indice].dato
	d.tabla[indice].estado = borrado
	d.cantidad--
	d.borrados++

	if float64(d.cantidad)/float64(len(d.tabla)) < FACTORCARGAMINIMO &&
		len(d.tabla)/2 >= CAPACIDADINICIAL {
		d.redimensionar(len(d.tabla) / 2)
	}

	return valor
}

func (d *hashCerrado[K, V]) Cantidad() int {
	return d.cantidad
}

func (d *hashCerrado[K, V]) Iterar(visitar func(K, V) bool) {
	for _, e := range d.tabla {
		if e.estado == ocupado {
			if !visitar(e.clave, e.dato) {
				break
			}
		}
	}
}
func (d *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorDiccionario[K, V]{dicc: d}
}

func (it *iteradorDiccionario[K, V]) HaySiguiente() bool {
	for it.actual < len(it.dicc.tabla) && it.dicc.tabla[it.actual].estado != ocupado {
		it.actual++
	}
	return it.actual < len(it.dicc.tabla)
}

func (it *iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return it.dicc.tabla[it.actual].clave, it.dicc.tabla[it.actual].dato
}

func (it *iteradorDiccionario[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	for it.actual < len(it.dicc.tabla) && it.dicc.tabla[it.actual].estado != ocupado {
		it.actual++
	}
	it.actual++
}

func (d *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {
	nuevaTabla := make([]celdaHash[K, V], nuevaCapacidad)

	for _, celda := range d.tabla {
		if celda.estado == ocupado {
			indice := hash(celda.clave, nuevaCapacidad)
			for {
				if nuevaTabla[indice].estado == vacio {
					nuevaTabla[indice] = celdaHash[K, V]{clave: celda.clave, dato: celda.dato, estado: ocupado}
					break
				}
				indice = (indice + 1) % nuevaCapacidad
			}
		}
	}

	d.tabla = nuevaTabla
	d.borrados = 0
}
