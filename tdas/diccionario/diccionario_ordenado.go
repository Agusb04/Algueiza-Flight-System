package diccionario

type DiccionarioOrdenado[K comparable, V any] interface {
	Guardar(K, V)
	Obtener(K) V
	Pertenece(K) bool
	Borrar(K) V
	Cantidad() int
	IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
	IterarRango(desde *K, hasta *K, visitar func(K, V) bool)
	Iterador() IterDiccionario[K, V]
}
