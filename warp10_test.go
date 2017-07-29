package Warp10Helper

import (
	"bytes"
	"testing"
	"time"
)

func TestGTSCreation(t *testing.T) {

	ts := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	classname := "test"
	labels := Labels{
		"ip": "1.2.3.4",
	}

	gts := CreateGTS(classname).WithLabels(labels).AddDatapoint(ts, "42")

	singleGTS := "1257894000000000// test{ip=1.2.3.4} \"42\"\n"

	var b bytes.Buffer
	gts.printGTS(&b)
	if b.String() != singleGTS {
		t.Fatalf("Expected '%v', got '%v'", singleGTS, b.String())
	}

	batch := NewBatch()
	batch.AddGTS(gts)

	var buf bytes.Buffer
	for _, gts := range *batch {
		gts.printGTS(&buf)
	}

	singleBatch := `1257894000000000// test{ip=1.2.3.4} "42"
`

	if buf.String() != singleBatch {
		t.Fatalf("Expected:\n%v\ngot\n%v", singleBatch, b.String())
	}
}
