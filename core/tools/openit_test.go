package tools

import "testing"

func TestOpen(t *testing.T) {
	if err := Open("http://127.0.0.1:8008"); err != nil {
		t.Error(err)
	}
}

func TestOpen2(t *testing.T) {
	if err := Open("/Users/"); err != nil {
		t.Error(err)
	}
}
