package device

import "testing"

func TestLoadExactHex(t *testing.T) {
	var s = "123"
	var buf []byte
	err := loadExactHex(buf, s)
	if err != nil {
		t.Errorf("error occurred")
	}
}

// go test noise-types_test.go noise-types.go
