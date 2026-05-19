package funciones_Aux

import (
	"strconv"
	"time"
)

const FORMATO = "2006-01-02T15:04:05"

func ParseInt(cadena string) int {
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		return -1
	}
	return numero
}

func ParseBool(cadena string) (bool, bool) {
	valores := map[string]bool{
		"1": true,
		"0": false,
	}

	valor, ok := valores[cadena]
	if !ok {
		return false, true
	}
	return valor, false
}

func ParseFecha(fecha string) (time.Time, bool) {
	t, err := time.Parse(FORMATO, fecha)
	if err != nil {
		return time.Time{}, true
	}
	return t, false
}

func ParseNumeroPositivo(input string, comando string) int {
	n, err := strconv.Atoi(input)
	if err != nil || n <= 0 {
		ImprimirError(comando)
		return -1
	}
	return n
}
