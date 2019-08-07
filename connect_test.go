package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConnection(t *testing.T) {

	port0 := "/dev/nonexistentport"
	port1 := "/tmp/bifrostmaster"

	_, err := NewConnection(port0, 115200)
	require.Error(t, err)

	_, err = NewConnection(port1, 115200)
	require.NoError(t, err)
}
