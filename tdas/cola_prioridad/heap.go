package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: []T{},
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	heap := &colaConPrioridad[T]{
		datos: copia,
		cant:  len(copia),
		cmp:   funcion_cmp,
	}

	for i := (heap.cant / 2) - 1; i >= 0; i-- {
		heapify(heap.datos, heap.cant, i, heap.cmp)
	}
	return heap
}

func (h *colaConPrioridad[T]) EstaVacia() bool {
	return h.cant == 0
}

func (h *colaConPrioridad[T]) Cantidad() int {
	return h.cant
}

func (h *colaConPrioridad[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *colaConPrioridad[T]) Encolar(dato T) {
	h.datos = append(h.datos, dato)
	h.cant++
	h.upheap(h.cant - 1)
}

func (h *colaConPrioridad[T]) upheap(i int) {
	for i > 0 {
		padre := (i - 1) / 2
		if h.cmp(h.datos[i], h.datos[padre]) <= 0 {
			break
		}
		h.datos[i], h.datos[padre] = h.datos[padre], h.datos[i]
		i = padre
	}
}

func (h *colaConPrioridad[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := h.datos[0]
	h.cant--
	h.datos[0] = h.datos[h.cant]
	h.datos = h.datos[:h.cant]
	h.downheap(0)
	h.redimensionar()

	return max
}

func (h *colaConPrioridad[T]) downheap(i int) {
	for {
		izq := 2*i + 1
		der := 2*i + 2
		mayor := i

		if izq < h.cant && h.cmp(h.datos[izq], h.datos[mayor]) > 0 {
			mayor = izq
		}
		if der < h.cant && h.cmp(h.datos[der], h.datos[mayor]) > 0 {
			mayor = der
		}
		if mayor == i {
			break
		}
		h.datos[i], h.datos[mayor] = h.datos[mayor], h.datos[i]
		i = mayor
	}
}

func HeapSort[T any](elementos []T, cmp func(T, T) int) {
	n := len(elementos)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(elementos, n, i, cmp)
	}

	for i := n - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		heapify(elementos, i, 0, cmp)
	}
}

func heapify[T any](elementos []T, n int, i int, cmp func(T, T) int) {

	for {
		izq := 2*i + 1
		der := 2*i + 2
		mayor := i

		if izq < n && cmp(elementos[izq], elementos[mayor]) > 0 {
			mayor = izq
		}
		if der < n && cmp(elementos[der], elementos[mayor]) > 0 {
			mayor = der
		}
		if mayor == i {
			break
		}
		elementos[i], elementos[mayor] = elementos[mayor], elementos[i]
		i = mayor
	}
}

func (h *colaConPrioridad[T]) redimensionar() {
	if cap(h.datos) > 2*h.cant {
		nuevo := make([]T, h.cant)
		copy(nuevo, h.datos)
		h.datos = nuevo
	}
}
