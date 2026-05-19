package diccionario

type nodoABB[K comparable, V any] struct {
	izq   *nodoABB[K, V]
	der   *nodoABB[K, V]
	clave K
	dato  V
}

type aBB[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iteradorABB[K comparable, V any] struct {
	nodos []*nodoABB[K, V]
	pos   int
}

func CrearABB[K comparable, V any](comparar func(K, K) int) DiccionarioOrdenado[K, V] {
	return &aBB[K, V]{cmp: comparar}
}

func (a *aBB[K, V]) Guardar(clave K, dato V) {
	a.raiz = a.guardar(a.raiz, clave, dato)
}

func (abb aBB[K, V]) Iterar(visitar func(K, V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (a *aBB[K, V]) guardar(nodo *nodoABB[K, V], clave K, dato V) *nodoABB[K, V] {
	if nodo == nil {
		a.cantidad++
		return &nodoABB[K, V]{clave: clave, dato: dato}
	}
	cmp := a.cmp(clave, nodo.clave)
	if cmp == 0 {
		nodo.dato = dato
	} else if cmp < 0 {
		nodo.izq = a.guardar(nodo.izq, clave, dato)
	} else {
		nodo.der = a.guardar(nodo.der, clave, dato)
	}
	return nodo
}

func (a *aBB[K, V]) Borrar(clave K) V {
	var borrado V
	a.raiz, borrado = a.borrar(a.raiz, clave)
	return borrado
}

func (a *aBB[K, V]) borrar(nodo *nodoABB[K, V], clave K) (*nodoABB[K, V], V) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	cmp := a.cmp(clave, nodo.clave)
	if cmp < 0 {
		var borrado V
		nodo.izq, borrado = a.borrar(nodo.izq, clave)
		return nodo, borrado
	} else if cmp > 0 {
		var borrado V
		nodo.der, borrado = a.borrar(nodo.der, clave)
		return nodo, borrado
	}

	if nodo.izq == nil {
		a.cantidad--
		return nodo.der, nodo.dato
	} else if nodo.der == nil {
		a.cantidad--
		return nodo.izq, nodo.dato
	}

	datoOriginal := nodo.dato
	sucesor := nodo.der
	for sucesor.izq != nil {
		sucesor = sucesor.izq
	}
	nodo.clave, nodo.dato = sucesor.clave, sucesor.dato
	nodo.der, _ = a.borrar(nodo.der, sucesor.clave)
	return nodo, datoOriginal
}

func (a aBB[K, V]) buscar(nodo *nodoABB[K, V], clave K) *nodoABB[K, V] {
	if nodo == nil {
		return nil
	}
	cmp := a.cmp(clave, nodo.clave)
	if cmp == 0 {
		return nodo
	} else if cmp < 0 {
		return a.buscar(nodo.izq, clave)
	} else {
		return a.buscar(nodo.der, clave)
	}
}

func (ABB aBB[K, V]) Pertenece(clave K) bool {
	return ABB.buscar(ABB.raiz, clave) != nil
}

func (ABB aBB[K, V]) Obtener(clave K) V {
	nodo := ABB.buscar(ABB.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (ABB aBB[K, V]) Cantidad() int {
	return ABB.cantidad
}

func (ABB aBB[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	var recorrer func(nodo *nodoABB[K, V]) bool
	recorrer = func(nodo *nodoABB[K, V]) bool {
		if nodo == nil {
			return true
		}

		if desde != nil && ABB.cmp(nodo.clave, *desde) < 0 {
			return recorrer(nodo.der)
		}

		if hasta != nil && ABB.cmp(nodo.clave, *hasta) > 0 {
			return recorrer(nodo.izq)
		}

		if !recorrer(nodo.izq) {
			return false
		}

		if !visitar(nodo.clave, nodo.dato) {
			return false
		}

		return recorrer(nodo.der)
	}
	recorrer(ABB.raiz)
}

func (ABB aBB[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	nodos := []*nodoABB[K, V]{}
	var recorrer func(nodo *nodoABB[K, V])
	recorrer = func(nodo *nodoABB[K, V]) {
		if nodo == nil {
			return
		}
		if desde != nil && ABB.cmp(nodo.clave, *desde) < 0 {
			recorrer(nodo.der)
			return
		}
		if hasta != nil && ABB.cmp(nodo.clave, *hasta) > 0 {
			recorrer(nodo.izq)
			return
		}
		recorrer(nodo.izq)
		nodos = append(nodos, nodo)
		recorrer(nodo.der)
	}
	recorrer(ABB.raiz)
	return &iteradorABB[K, V]{nodos: nodos, pos: 0}

}

func (ABB aBB[K, V]) Iterador() IterDiccionario[K, V] {
	return ABB.IteradorRango(nil, nil)
}

func (it *iteradorABB[K, V]) HaySiguiente() bool {
	return it.pos < len(it.nodos)
}

func (it *iteradorABB[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.nodos[it.pos]
	return nodo.clave, nodo.dato
}

func (it *iteradorABB[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	it.pos++
}
