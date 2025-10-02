package gtypes_test

import (
	"testing"

	"github.com/cqroot/gtypes"
	"github.com/stretchr/testify/require"
)

func TestOMap(t *testing.T) {
	m := gtypes.NewOrderedMap[string, string]()
	require.Equal(t, 0, m.Len())

	testcases := []struct {
		k string
		v string
	}{
		{k: "key 0", v: "val 0"},
		{k: "key 1", v: "val 1"},
		{k: "key 2", v: "val 2"},
		{k: "key 3", v: "val 3"},
		{k: "key 4", v: "val 4"},
		{k: "key 5", v: "val 5"},
	}

	for i, testcase := range testcases {
		require.Equal(t, false, m.Has(testcase.k))

		m.Put(testcase.k, testcase.v)
		require.Equal(t, i+1, m.Len())

		val, exists := m.Get(testcase.k)
		require.Equal(t, testcase.v, val)
		require.Equal(t, true, exists)

		require.Equal(t, true, m.Has(testcase.k))
	}

	for i, testcase := range testcases {
		require.Equal(t, len(testcases)-i, m.Len())

		require.Equal(t, true, m.Has(testcase.k))
		m.Remove(testcase.k)
		require.Equal(t, false, m.Has(testcase.k))

		require.Equal(t, len(testcases)-i-1, m.Len())
	}
}
