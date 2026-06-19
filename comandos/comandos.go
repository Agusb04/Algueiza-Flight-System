package comandos

import (
	"algueiza-flight-system/funciones_Aux"
	"algueiza-flight-system/interfaz"
)

func ManejoComandos(sistema interfaz.SistemaVuelos, comando string, argumentos []string) {

	defer func() {
		if r := recover(); r != nil {
			funciones_Aux.ImprimirError(comando)
		}
	}()
	switch comando {
	case "agregar_archivo":

		if len(argumentos) != 1 {
			funciones_Aux.ImprimirError(comando)
			return
		}
		ruta := argumentos[0]
		if !funciones_Aux.ArchivoExiste(ruta) {
			funciones_Aux.ImprimirError(comando)
			return
		}
		sistema.CargarArchivo(ruta)
		funciones_Aux.PrintOk()

	case "ver_tablero":

		if len(argumentos) != 4 {
			funciones_Aux.ImprimirError(comando)
			return
		}
		K, modo, desde, hasta := argumentos[0], argumentos[1], argumentos[2], argumentos[3]
		cantidad := funciones_Aux.ParseNumeroPositivo(K, comando)
		if cantidad < 0 {
			return
		}
		if modo != "asc" && modo != "desc" {
			funciones_Aux.ImprimirError(comando)
			return
		}
		fechadesde, err1 := funciones_Aux.ParseFecha(desde)
		fechahasta, err2 := funciones_Aux.ParseFecha(hasta)
		if err1 || err2 {
			funciones_Aux.ImprimirError(comando)
			return
		}
		if fechahasta.Before(fechadesde) {
			funciones_Aux.ImprimirError(comando)
			return
		}

		sistema.VerTablero(cantidad, modo, desde, hasta)
		funciones_Aux.PrintOk()

	case "info_vuelo":

		if len(argumentos) != 1 {
			funciones_Aux.ImprimirError(comando)
			return
		}

		codigo := argumentos[0]
		if funciones_Aux.ParseNumeroPositivo(codigo, comando) == -1 {
			return
		}

		if !sistema.InfoVuelo(codigo) {
			funciones_Aux.ImprimirError(comando)
			return
		}
		funciones_Aux.PrintOk()

	case "prioridad_vuelos":

		if len(argumentos) != 1 {
			funciones_Aux.ImprimirError(comando)
			return
		}

		numeroVuelos := funciones_Aux.ParseNumeroPositivo(argumentos[0], comando)
		if numeroVuelos == -1 {
			return
		}

		sistema.PrioridadVuelos(numeroVuelos)
		funciones_Aux.PrintOk()

	case "siguiente_vuelo":
		if len(argumentos) != 3 {
			funciones_Aux.ImprimirError(comando)
			return
		}

		origen, destino, fecha := argumentos[0], argumentos[1], argumentos[2]
		_, err := funciones_Aux.ParseFecha(fecha)
		if err {
			funciones_Aux.ImprimirError(comando)
			return
		}
		sistema.SiguienteVuelo(origen, destino, fecha)

		funciones_Aux.PrintOk()

	case "borrar":
		if len(argumentos) != 2 {
			funciones_Aux.ImprimirError(comando)
			return
		}

		desde, hasta := argumentos[0], argumentos[1]
		fechadesde, err1 := funciones_Aux.ParseFecha(desde)
		fechahasta, err2 := funciones_Aux.ParseFecha(hasta)
		if err1 || err2 {
			funciones_Aux.ImprimirError(comando)
			return
		}
		if fechadesde.After(fechahasta) {
			funciones_Aux.ImprimirError(comando)
			return
		}
		sistema.Borrar(desde, hasta)
		funciones_Aux.PrintOk()

	default:
		funciones_Aux.ImprimirError(comando)
	}

}
