package main

import "testing"

func TestName(t *testing.T) {
	name := name()
	expected := "Leander"
	if name != expected {
		t.Error("Wrong name")
	}
}

func TestNameFail(t *testing.T) {
	name := name()
	if name != "Leander" {
		t.Error("Wrong name")
	}
}
