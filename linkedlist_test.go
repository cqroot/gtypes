package gtypes_test

import (
	"testing"

	"github.com/cqroot/gtypes"
	"github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	l := gtypes.NewLinkedList[string]()
	require.Equal(t, 0, l.Size())

	l.Add("a")
	require.Equal(t, 1, l.Size())
	require.Equal(t, 0, l.IndexOf("a"))

	l.Add("b")
	require.Equal(t, 2, l.Size())
	require.Equal(t, 1, l.IndexOf("b"))

	l.Remove(0)
	require.Equal(t, 1, l.Size())
	require.Equal(t, 0, l.IndexOf("b"))
}
