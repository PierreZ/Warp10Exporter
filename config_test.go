package Warp10Exporter

import "testing"

func TestSetURI(t *testing.T) {
	SetURI("unicorn")
	expected := "unicorn"
	if warpURI != expected {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", expected, warpURI)
	}

	SetURI("/api/v0/update")
}

func TestSetHeader(t *testing.T) {
	SetHeader("unicorn")
	expected := "unicorn"
	if warpHeader != expected {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", expected, warpHeader)
	}

	SetHeader("X-Warp10-Token")
}