package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConnection(t *testing.T) {

	port0 := "/dev/nonexistentport"
	port1 := "/tmp/bifrostport1"
	port2 := "/tmp/bifrostport2"
	testInput := []byte("from Asgard...\n")
	message1 := make([]byte, len(testInput))

	_, err := NewConnection(port0, 115200)
	require.Error(t, err)

	connect1, err := NewConnection(port1, 115200)
	require.NoError(t, err)

	connect2, err := NewConnection(port2, 115200)
	require.NoError(t, err)

	connect1.Write(testInput)
	_, err = io.ReadFull(connect2.port, message1)
	require.NoError(t, err)
	assert.Equal(t, testInput, message1)
}
