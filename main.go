package main

import (
	"bufio"
	"os"
	"strings"

	"algueiza-flight-system/comandos"
	"algueiza-flight-system/sistema"
)

func main() {

	sistema := sistema.CrearSistemaVuelos()
	texto := bufio.NewScanner(os.Stdin)

	for texto.Scan() {
		linea := strings.TrimSpace(texto.Text())
		if linea == "" {
			continue
		}

		palabras := strings.Fields(linea)
		if len(palabras) == 0 {
			continue
		}

		comando := palabras[0]
		argumentos := palabras[1:]

		comandos.ManejoComandos(sistema, comando, argumentos)
	}
}
