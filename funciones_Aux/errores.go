package funciones_Aux

import (
	"fmt"
	"os"
)

func ImprimirError(comando string) {
	fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
}

func PrintOk() {
	fmt.Println("OK")
}

func ArchivoExiste(ruta string) bool {

	if len(ruta) <= 4 || ruta[len(ruta)-4:] != ".csv" {
		return false
	}

	_, err := os.Stat(ruta)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
