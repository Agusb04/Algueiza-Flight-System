package comandos

import (
	"bytes"
	"io"
	"os"
	"testing"

	"algueiza-flight-system/interfaz"
	"algueiza-flight-system/sistema"

	"github.com/stretchr/testify/require"
)

func captureOutput(t *testing.T, f func()) (string, string) {
	t.Helper()
	oldStdout, oldStderr := os.Stdout, os.Stderr
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr
	f()
	wOut.Close()
	wErr.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr
	var bufOut, bufErr bytes.Buffer
	io.Copy(&bufOut, rOut)
	io.Copy(&bufErr, rErr)
	return bufOut.String(), bufErr.String()
}

func crearSistemaConVuelos(t *testing.T) interfaz.SistemaVuelos {
	t.Helper()
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")
	return s
}

func TestManejoComandosAgregarArchivo(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "agregar_archivo", []string{"../data/vuelos-algueiza-parte-01.csv"})
	})
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosAgregarArchivoSinArgumentos(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "agregar_archivo", []string{})
	})
	require.Contains(t, stderr, "Error en comando agregar_archivo")
}

func TestManejoComandosAgregarArchivoInexistente(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "agregar_archivo", []string{"no_existe.csv"})
	})
	require.Contains(t, stderr, "Error en comando agregar_archivo")
}

func TestManejoComandosInfoVuelo(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "info_vuelo", []string{"4608"})
	})
	require.Contains(t, stdout, "4608")
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosInfoVueloNoExistente(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "info_vuelo", []string{"INEXISTENTE"})
	})
	require.Contains(t, stderr, "Error en comando info_vuelo")
}

func TestManejoComandosInfoVueloSinArgumentos(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "info_vuelo", []string{})
	})
	require.Contains(t, stderr, "Error en comando info_vuelo")
}

func TestManejoComandosVerTableroAsc(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "ver_tablero", []string{"3", "asc", "2018-01-01T00:00:00", "2018-12-31T23:59:59"})
	})
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosVerTableroDesc(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "ver_tablero", []string{"3", "desc", "2018-01-01T00:00:00", "2018-12-31T23:59:59"})
	})
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosVerTableroParametrosInvalidos(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "ver_tablero", []string{"abc", "asc", "2018-01-01T00:00:00", "2018-12-31T23:59:59"})
	})
	require.Contains(t, stderr, "Error en comando ver_tablero")
}

func TestManejoComandosPrioridadVuelos(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "prioridad_vuelos", []string{"3"})
	})
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosPrioridadVuelosInvalido(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "prioridad_vuelos", []string{"0"})
	})
	require.Contains(t, stderr, "Error en comando prioridad_vuelos")
}

func TestManejoComandosSiguienteVuelo(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "siguiente_vuelo", []string{"PDX", "SEA", "2018-01-01T00:00:00"})
	})
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosSiguienteVueloFechaInvalida(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "siguiente_vuelo", []string{"PDX", "SEA", "fecha-invalida"})
	})
	require.Contains(t, stderr, "Error en comando siguiente_vuelo")
}

func TestManejoComandosBorrar(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	s.CargarArchivo("../data/vuelos-algueiza-parte-01.csv")

	stdout, stderr := captureOutput(t, func() {
		ManejoComandos(s, "borrar", []string{"2018-01-01T00:00:00", "2018-06-01T00:00:00"})
	})
	require.Contains(t, stdout, "OK")
	require.Empty(t, stderr)
}

func TestManejoComandosBorrarFechasInvalidas(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "borrar", []string{"2018-12-01T00:00:00", "2018-06-01T00:00:00"})
	})
	require.Contains(t, stderr, "Error en comando borrar")
}

func TestManejoComandosDefault(t *testing.T) {
	s := sistema.CrearSistemaVuelos()
	_, stderr := captureOutput(t, func() {
		ManejoComandos(s, "comando_invalido", []string{})
	})
	require.Contains(t, stderr, "Error en comando comando_invalido")
}
