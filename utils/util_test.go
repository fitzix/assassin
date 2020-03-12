package utils

import "testing"

func TestNextID(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Log(NextID())
	}
}
