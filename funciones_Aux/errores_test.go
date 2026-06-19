package funciones_Aux

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func captureStderr(t *testing.T, f func()) string {
	t.Helper()
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	f()
	w.Close()
	os.Stderr = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

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

func TestImprimirError(t *testing.T) {
	output := captureStderr(t, func() {
		ImprimirError("test_comando")
	})
	require.Contains(t, output, "Error en comando test_comando")
}

func TestPrintOk(t *testing.T) {
	output := captureStdout(t, func() {
		PrintOk()
	})
	require.Equal(t, "OK\n", output)
}

func TestArchivoExiste(t *testing.T) {
	require.False(t, ArchivoExiste(""))
	require.False(t, ArchivoExiste("no_existe.txt"))
	require.False(t, ArchivoExiste("no_existe.csv"))

	tmpfile, err := os.CreateTemp("", "test*.csv")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	require.True(t, ArchivoExiste(tmpfile.Name()))
}

func TestArchivoExisteSoloCSV(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test*.go")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	require.False(t, ArchivoExiste(tmpfile.Name()))
}
