package tun

import (
	"testing"
)

func TestGetIFIndex(t *testing.T) {
	_, err := getIFIndex("test")
	if err == nil {
		t.Errorf("test irq is not exist, err can't be nil")
	}
}
