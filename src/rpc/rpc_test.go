package rpc_test

import (
	"testing"

	"github.com/brunobmello25/educationalsp/src/rpc"
	"github.com/stretchr/testify/assert"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	assert.Equal(t, expected, actual)
}
