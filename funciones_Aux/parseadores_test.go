package funciones_Aux

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseInt(t *testing.T) {
	require.Equal(t, 42, ParseInt("42"))
	require.Equal(t, -1, ParseInt("foo"))
	require.Equal(t, -1, ParseInt(""))
	require.Equal(t, 0, ParseInt("0"))
}

func TestParseBool(t *testing.T) {
	val, err := ParseBool("1")
	require.True(t, val)
	require.False(t, err)

	val, err = ParseBool("0")
	require.False(t, val)
	require.False(t, err)

	val, err = ParseBool("true")
	require.False(t, val)
	require.True(t, err)

	val, err = ParseBool("")
	require.False(t, val)
	require.True(t, err)
}

func TestParseFecha(t *testing.T) {
	fecha, err := ParseFecha("2018-10-10T00:00:00")
	require.False(t, err)
	expected := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	require.Equal(t, expected, fecha)

	_, err = ParseFecha("10-10-2018")
	require.True(t, err)

	_, err = ParseFecha("")
	require.True(t, err)

	_, err = ParseFecha("2018/10/10T00:00:00")
	require.True(t, err)
}

func TestParseNumeroPositivo(t *testing.T) {
	require.Equal(t, 5, ParseNumeroPositivo("5", "test"))
	require.Equal(t, -1, ParseNumeroPositivo("0", "test"))
	require.Equal(t, -1, ParseNumeroPositivo("-3", "test"))
	require.Equal(t, -1, ParseNumeroPositivo("abc", "test"))
	require.Equal(t, -1, ParseNumeroPositivo("", "test"))
}
