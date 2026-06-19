package sistema

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"algueiza-flight-system/interfaz"

	"github.com/stretchr/testify/require"
)

const testCSV = `4608,OO,PDX,SEA,N812SK,08,2018-04-10T23:22:55,05,43,0
5243,OO,PIA,DEN,N903SW,04,2018-08-07T18:05:12,-08,69,0
1092,UA,FLL,EWR,N34455,01,2018-09-05T16:33:02,-04,21,0
4701,EV,EWR,CMH,N11150,12,2018-10-04T04:19:24,-10,55,0
1086,UA,DTW,IAH,N27733,15,2018-06-11T06:02:25,-03,132,1
1863,DL,DTW,MCO,N584NW,06,2018-11-03T18:40:31,07,47,0
`

func captureStdout(t *testing.T, f func()) string {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func crearSistemaConVuelos(t *testing.T) (interfaz.SistemaVuelos, string) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "test_vuelos*.csv")
	require.NoError(t, err)
	_, err = tmpfile.WriteString(testCSV)
	require.NoError(t, err)
	tmpfile.Close()

	s := CrearSistemaVuelos()
	s.CargarArchivo(tmpfile.Name())
	return s, tmpfile.Name()
}

func TestCrearSistemaVuelos(t *testing.T) {
	s := CrearSistemaVuelos()
	require.NotNil(t, s)
}

func TestCargarArchivo(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.InfoVuelo("4608")
	})
	require.Contains(t, output, "4608")
	require.Contains(t, output, "PDX")
}

func TestInfoVueloExistente(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		ok := s.InfoVuelo("4608")
		require.True(t, ok)
	})
	require.Contains(t, output, "4608 OO PDX SEA")
	require.Contains(t, output, "2018-04-10T23:22:55")
}

func TestInfoVueloNoExistente(t *testing.T) {
	s := CrearSistemaVuelos()
	ok := s.InfoVuelo("INEXISTENTE")
	require.False(t, ok)
}

func TestInfoVueloConVueloCancelado(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		ok := s.InfoVuelo("1086")
		require.True(t, ok)
	})
	require.Contains(t, output, "1086 UA DTW IAH")
	require.Contains(t, output, "1") // cancelado
}

func TestVerTableroAscendente(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.VerTablero(3, "asc", "2018-01-01T00:00:00", "2018-12-31T23:59:59")
	})

	lines := strings.Split(strings.TrimSpace(output), "\n")
	require.Equal(t, 3, len(lines))
	require.Contains(t, lines[0], "2018-04-10T23:22:55") // fecha mas temprana
}

func TestVerTableroDescendente(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.VerTablero(3, "desc", "2018-01-01T00:00:00", "2018-12-31T23:59:59")
	})

	lines := strings.Split(strings.TrimSpace(output), "\n")
	require.Equal(t, 3, len(lines))
	require.Contains(t, lines[0], "2018-11-03") // fecha mas tardia
}

func TestVerTableroModoInvalido(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	_ = captureStdout(t, func() {
		s.VerTablero(3, "invalid", "2018-01-01T00:00:00", "2018-12-31T23:59:59")
	})
}

func TestPrioridadVuelos(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.PrioridadVuelos(3)
	})
	require.Contains(t, output, "15")
	require.Contains(t, output, "12")
}

func TestPrioridadVuelosConCero(t *testing.T) {
	s := CrearSistemaVuelos()
	// no debe panic
	_ = captureStdout(t, func() {
		s.PrioridadVuelos(0)
	})
}

func TestSiguienteVueloExistente(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.SiguienteVuelo("PDX", "SEA", "2018-01-01T00:00:00")
	})
	require.Contains(t, output, "4608")
}

func TestSiguienteVueloSinConexion(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.SiguienteVuelo("PDX", "JFK", "2018-01-01T00:00:00")
	})
	require.Contains(t, output, "No hay vuelo registrado desde PDX hacia JFK")
}

func TestSiguienteVueloOrigenInexistente(t *testing.T) {
	s := CrearSistemaVuelos()
	output := captureStdout(t, func() {
		s.SiguienteVuelo("AAA", "BBB", "2018-01-01T00:00:00")
	})
	require.Contains(t, output, "No hay vuelo registrado desde AAA hacia BBB")
}

func TestSiguienteVueloConCancelado(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.SiguienteVuelo("DTW", "IAH", "2018-01-01T00:00:00")
	})
	require.Contains(t, output, "No hay vuelo registrado")
}

func TestSiguienteVueloNoCancelado(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.SiguienteVuelo("DTW", "MCO", "2018-01-01T00:00:00")
	})
	require.Contains(t, output, "1863")
}

func TestBorrarRango(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	output := captureStdout(t, func() {
		s.Borrar("2018-01-01T00:00:00", "2018-07-01T00:00:00")
	})
	require.Contains(t, output, "4608")

	ok := s.InfoVuelo("4608")
	require.False(t, ok)

	ok = s.InfoVuelo("1092")
	require.True(t, ok)
}

func TestCargarArchivoConRemplazo(t *testing.T) {
	s, path := crearSistemaConVuelos(t)
	defer os.Remove(path)

	ok := s.InfoVuelo("4608")
	require.True(t, ok)

	s.CargarArchivo(path)
	ok = s.InfoVuelo("4608")
	require.True(t, ok)
}

func TestVuelosInvalidosSaltados(t *testing.T) {
	badCSV := `INVALID,NO,NUM,COLS
9999,AA,JFK,LAX,N12345,XX,2018-01-01T00:00:00,00,00,0
`
	tmpfile, err := os.CreateTemp("", "bad_vuelos*.csv")
	require.NoError(t, err)
	_, err = tmpfile.WriteString(badCSV)
	require.NoError(t, err)
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	s := CrearSistemaVuelos()
	s.CargarArchivo(tmpfile.Name())

	ok := s.InfoVuelo("9999")
	require.False(t, ok)

	ok = s.InfoVuelo("INVALID")
	require.False(t, ok)
}
