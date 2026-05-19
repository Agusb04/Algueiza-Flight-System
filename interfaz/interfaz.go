package interfaz

type SistemaVuelos interface {
	CargarArchivo(ruta string)

	VerTablero(cant int, modo string, fechaDesde, fechaHasta string)

	InfoVuelo(codigo string)

	PrioridadVuelos(cant int)

	SiguienteVuelo(origen, destino string, fecha string)

	Borrar(desde, hasta string)
}
