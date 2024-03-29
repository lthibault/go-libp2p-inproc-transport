package inproc

import (
	"testing"

	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func TestMultiaddr(t *testing.T) {
	t.Parallel()
	t.Helper()

	t.Run("ToMultiaddr", func(t *testing.T) {
		t.Parallel()

		t.Skip("NOT IMPLEMENTED")
	})

	t.Run("ToNetAddr", func(t *testing.T) {
		t.Parallel()

		a, err := toInprocNetAddr(multiaddr.StringCast("/inproc/test"))
		require.NoError(t, err)
		require.IsType(t, addr{}, a)
		require.Equal(t, prefix, a.Network())
		require.Equal(t, "/test", a.String())
	})
}

func TestResolveString(t *testing.T) {
	ma, err := ResolveString("/inproc/~")
	require.NoError(t, err)

	s, err := ma.ValueForProtocol(P_INPROC)
	require.NoError(t, err)
	require.NotEmpty(t, s)
	require.NotEqual(t, "~", s, "should have expanded ~")
}
