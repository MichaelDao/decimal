package decimal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecimal64Gob(t *testing.T) {
	require := require.New(t)

	gob, err := NewDecimal64FromInt64(23456).GobEncode()
	require.NoError(err)

	var d Decimal64
	require.NoError(d.GobDecode(gob))
	require.Equal(NewDecimal64FromInt64(23456), d)
}
