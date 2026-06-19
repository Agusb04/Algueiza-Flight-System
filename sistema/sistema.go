package sistema

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	Heap "algueiza-flight-system/tdas/cola_prioridad"
	DICCIONARIO "algueiza-flight-system/tdas/diccionario"

	"algueiza-flight-system/funciones_Aux"
	"algueiza-flight-system/interfaz"
)

const (
	FORMATO       = "2006-01-02T15:04:05"
	CMD_PRIORIDAD = "prioridad_vuelos"
	CMD_TABLERO   = "ver_tablero"
	CMD_GUARDAR   = "agregar_archivo"
)

type sistema struct {
	vuelosPorFechas DICCIONARIO.DiccionarioOrdenado[IdentificadorVuelo, struct{}]
	vuelosPorCodigo DICCIONARIO.Diccionario[string, Vuelo]
	conexiones      DICCIONARIO.Diccionario[string, DICCIONARIO.Diccionario[string, DICCIONARIO.DiccionarioOrdenado[string, *Vuelo]]]
}

type Vuelo struct {
	prioridad      int
	airTime        int
	departureDelay int
	codigo         string
	aerolinea      string
	origen         string
	destino        string
	matricula      string
	fechaHora      string
	cancelado      bool
}

type IdentificadorVuelo struct {
	fecha  string
	codigo string
}

// --- Funciones constructoras y públicas

func CrearSistemaVuelos() interfaz.SistemaVuelos {
	return &sistema{
		vuelosPorFechas: DICCIONARIO.CrearABB[IdentificadorVuelo, struct{}](compararIdentificadorVuelo),
		vuelosPorCodigo: DICCIONARIO.CrearHash[string, Vuelo](),
		conexiones:      DICCIONARIO.CrearHash[string, DICCIONARIO.Diccionario[string, DICCIONARIO.DiccionarioOrdenado[string, *Vuelo]]](),
	}
}

func (sis *sistema) CargarArchivo(ruta string) {
	vuelosArchivo := DICCIONARIO.CrearHash[string, []string]()

	archivo, err := os.Open(ruta)
	if err != nil {
		funciones_Aux.ImprimirError(CMD_GUARDAR)
		return
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := strings.Split(scanner.Text(), ",")
		for i := range linea {
			linea[i] = strings.TrimSpace(linea[i])
		}
		if !esFormatoValido(linea) {
			continue
		}
		codigo := linea[0]

		vuelosArchivo.Guardar(codigo, linea)
	}

	seAgregaronNuevos := false

	iter := vuelosArchivo.Iterador()
	for iter.HaySiguiente() {
		codigo, linea := iter.VerActual()
		nVuelo, ok := construirVuelo(linea)
		if !ok {
			iter.Siguiente()
			continue
		}

		if sis.vuelosPorCodigo.Pertenece(codigo) {
			vueloAnt := sis.vuelosPorCodigo.Obtener(codigo)
			claveAnt := IdentificadorVuelo{fecha: vueloAnt.fechaHora, codigo: vueloAnt.codigo}

			sis.vuelosPorCodigo.Borrar(codigo)
			sis.vuelosPorFechas.Borrar(claveAnt)
			sis.borrarConexion(&vueloAnt)
		}

		sis.vuelosPorCodigo.Guardar(codigo, nVuelo)
		claveNuevo := IdentificadorVuelo{fecha: nVuelo.fechaHora, codigo: codigo}
		sis.vuelosPorFechas.Guardar(claveNuevo, struct{}{})
		sis.agregarConexion(&nVuelo)

		seAgregaronNuevos = true

		iter.Siguiente()
	}

	if !seAgregaronNuevos {
		funciones_Aux.ImprimirError(CMD_GUARDAR)
	}
}

func (s *sistema) VerTablero(k int, modo string, desdeStr, hastaStr string) {
	desde := IdentificadorVuelo{fecha: desdeStr, codigo: ""}
	hasta := IdentificadorVuelo{fecha: hastaStr, codigo: "\U0010FFFF"}

	iter := s.vuelosPorFechas.IteradorRango(&desde, &hasta)

	if modo == "asc" {
		count := 0
		for iter.HaySiguiente() && count < k {
			vueloID, _ := iter.VerActual()
			s.mostrarVueloEnTablero(vueloID)
			iter.Siguiente()
			count++
		}
		return
	}

	if modo == "desc" {
		var vuelos []IdentificadorVuelo
		for iter.HaySiguiente() {
			vueloID, _ := iter.VerActual()
			vuelos = append(vuelos, vueloID)
			iter.Siguiente()
		}

		for i, j := 0, len(vuelos)-1; i < j; i, j = i+1, j-1 {
			vuelos[i], vuelos[j] = vuelos[j], vuelos[i]
		}

		limite := k
		if len(vuelos) < limite {
			limite = len(vuelos)
		}

		for i := 0; i < limite; i++ {
			s.mostrarVueloEnTablero(vuelos[i])
		}
		return
	}

	funciones_Aux.ImprimirError(CMD_TABLERO)
}

func (s *sistema) InfoVuelo(codigo string) bool {
	if !s.vuelosPorCodigo.Pertenece(codigo) {
		return false
	}
	vuelo := s.vuelosPorCodigo.Obtener(codigo)
	s.mostrarDetallesVuelo(&vuelo)
	return true
}

func (s *sistema) PrioridadVuelos(k int) {

	if k <= 0 {
		funciones_Aux.ImprimirError(CMD_PRIORIDAD)
		return
	}

	iter := s.vuelosPorCodigo.Iterador()
	todos := []Vuelo{}

	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		todos = append(todos, v)
		iter.Siguiente()
	}

	heap := Heap.CrearHeapArr(todos, comparacionVuelos)

	limite := k
	if heap.Cantidad() < k {
		limite = heap.Cantidad()
	}

	for i := 0; i < limite; i++ {
		v := heap.Desencolar()
		fmt.Printf("%d - %s\n", v.prioridad, v.codigo)
	}

}

func (s *sistema) SiguienteVuelo(origen, destino string, fecha string) {

	if !s.conexiones.Pertenece(origen) {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
		return
	}

	destinosDesdeOrigen := s.conexiones.Obtener(origen)

	if !destinosDesdeOrigen.Pertenece(destino) {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
		return
	}

	fechas := destinosDesdeOrigen.Obtener(destino)

	iter := fechas.IteradorRango(&fecha, nil)

	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()

		if !vuelo.cancelado {
			s.InfoVuelo(vuelo.codigo)
			return
		}

		iter.Siguiente()
	}

	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
}

func (s *sistema) Borrar(desde, hasta string) {
	vuelosEliminados := []*Vuelo{}
	clavesParaEliminar := []IdentificadorVuelo{}

	rangoInicio := IdentificadorVuelo{desde, ""}
	rangoFin := IdentificadorVuelo{hasta, "\U0010FFFF"}

	iterador := s.vuelosPorFechas.IteradorRango(&rangoInicio, &rangoFin)
	for iterador.HaySiguiente() {
		claveActual, _ := iterador.VerActual()
		clavesParaEliminar = append(clavesParaEliminar, claveActual)
		iterador.Siguiente()
	}

	for _, clave := range clavesParaEliminar {
		if !s.vuelosPorCodigo.Pertenece(clave.codigo) || !s.vuelosPorFechas.Pertenece(clave) {
			continue
		}

		vueloActual := s.vuelosPorCodigo.Obtener(clave.codigo)
		vuelosEliminados = append(vuelosEliminados, &vueloActual)

		s.vuelosPorCodigo.Borrar(vueloActual.codigo)
		s.vuelosPorFechas.Borrar(clave)

		if s.conexiones.Pertenece(vueloActual.origen) {
			destinosDesdeOrigen := s.conexiones.Obtener(vueloActual.origen)
			if destinosDesdeOrigen.Pertenece(vueloActual.destino) {
				fechasDeDestino := destinosDesdeOrigen.Obtener(vueloActual.destino)
				if fechasDeDestino.Pertenece(clave.fecha) {
					fechasDeDestino.Borrar(clave.fecha)
				}
			}
		}
	}

	for _, vuelo := range vuelosEliminados {
		s.mostrarDetallesVuelo(vuelo)
	}
}

// --- Métodos privados del sistema

func (s *sistema) mostrarDetallesVuelo(v *Vuelo) {
	cancelado := 0
	if v.cancelado {
		cancelado = 1
	}

	fmt.Printf("%s %s %s %s %s %d %s %d %d %d\n",
		v.codigo, v.aerolinea, v.origen, v.destino, v.matricula,
		v.prioridad, v.fechaHora, v.departureDelay, v.airTime, cancelado)
}

func (s *sistema) agregarConexion(vuelo *Vuelo) {
	puntoPartida := vuelo.origen
	puntoDestino := vuelo.destino
	momentoVuelo := vuelo.fechaHora

	if !s.conexiones.Pertenece(puntoPartida) {
		s.conexiones.Guardar(puntoPartida, DICCIONARIO.CrearHash[string, DICCIONARIO.DiccionarioOrdenado[string, *Vuelo]]())
	}

	destinosMap := s.conexiones.Obtener(puntoPartida)

	if !destinosMap.Pertenece(puntoDestino) {
		destinosMap.Guardar(puntoDestino, DICCIONARIO.CrearABB[string, *Vuelo](comparacionTiempo))
	}

	vuelosPorFecha := destinosMap.Obtener(puntoDestino)
	vuelosPorFecha.Guardar(momentoVuelo, vuelo)
}

func (s *sistema) borrarConexion(vuelo *Vuelo) {
	puntoPartida := vuelo.origen
	puntoDestino := vuelo.destino
	momentoVuelo := vuelo.fechaHora

	if !s.conexiones.Pertenece(puntoPartida) {
		return
	}

	destinosMap := s.conexiones.Obtener(puntoPartida)
	if !destinosMap.Pertenece(puntoDestino) {
		return
	}

	vuelosPorFecha := destinosMap.Obtener(puntoDestino)
	vuelosPorFecha.Borrar(momentoVuelo)
}

func (s *sistema) mostrarVueloEnTablero(id IdentificadorVuelo) {
	fmt.Printf("%s - %s\n", id.fecha, id.codigo)
}

// --- Funciones auxiliares y helpers

func comparacionVuelos(a, b Vuelo) int {
	if a.prioridad != b.prioridad {
		return a.prioridad - b.prioridad
	}
	return strings.Compare(b.codigo, a.codigo)
}

func compararIdentificadorVuelo(a, b IdentificadorVuelo) int {
	if a.fecha > b.fecha {
		return 1
	} else if a.fecha < b.fecha {
		return -1
	}
	return strings.Compare(a.codigo, b.codigo)
}

func comparacionTiempo(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func esFormatoValido(l []string) bool {
	return len(l) == 10
}

func construirVuelo(campos []string) (Vuelo, bool) {

	prioridad, airTime := funciones_Aux.ParseInt(campos[5]), funciones_Aux.ParseInt(campos[8])
	departureDelay, errDelay := strconv.Atoi(campos[7])
	_, errFecha := funciones_Aux.ParseFecha(campos[6])
	fecha := campos[6]

	cancelado, errBool := funciones_Aux.ParseBool(campos[9])
	switch {
	case prioridad == -1 || errDelay != nil || airTime == -1:
		return Vuelo{}, false
	case errBool:
		return Vuelo{}, false
	case errFecha:
		return Vuelo{}, false
	}
	return Vuelo{
		codigo:         campos[0],
		prioridad:      prioridad,
		departureDelay: departureDelay,
		airTime:        airTime,
		fechaHora:      fecha,
		aerolinea:      campos[1],
		origen:         campos[2],
		destino:        campos[3],
		matricula:      campos[4],
		cancelado:      cancelado,
	}, true
}

