package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConnection(t *testing.T) {

	port0 := "/dev/nonexistentport"
	port1 := "/tmp/bifrostport1"
	port2 := "/tmp/bifrostport2"
	testInput := []byte("from Asgard...\n")

	_, err := NewConnection(port0, 115200)
	require.Error(t, err)

	connect1, err := NewConnection(port1, 115200)
	require.NoError(t, err)

	connect2, err := NewConnection(port2, 115200)
	require.NoError(t, err)

	connect1.Write(testInput)
	message1, err := connect2.portReader.ReadBytes('\n')
	require.NoError(t, err)
	assert.Equal(t, testInput, message1)
}
