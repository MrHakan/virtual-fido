package util

import (
	"testing"

	"github.com/bulwarkid/virtual-fido/test"
)

func TestRequestBuffer(t *testing.T) {
	buffer := MakeRequestBuffer()
	makeRequest := func(val []byte) func([]byte) {
		return func(response []byte) {
			test.AssertEqual(t, response[0], val[0], "Request does not equal response")
		}
	}
	// Exercise waiting requests without assuming an order for the internal map.
	for value := byte(1); value <= 3; value++ {
		if buffer.Request(uint32(value), makeRequest([]byte{value})) {
			t.Fatal("request unexpectedly completed before response")
		}
		buffer.Respond([]byte{value})
	}
	// Responses queued before requests must be delivered in arrival order.
	for value := byte(4); value <= 6; value++ {
		buffer.Respond([]byte{value})
	}
	for value := byte(4); value <= 6; value++ {
		if !buffer.Request(uint32(value), makeRequest([]byte{value})) {
			t.Fatal("queued response was not delivered")
		}
	}
}
