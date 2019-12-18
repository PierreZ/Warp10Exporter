package warp10exporter

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewBatch(t *testing.T) {

	var buf bytes.Buffer
	batch := NewBatch()
	batch.Register(singleGTSSingleDatapoint)

	batch.Print(&buf)

	if buf.String() != singleGTSSingleDatapointString {
		t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, buf.String())
	}

	buf.Reset()
	batch = NewBatch()

	gts1 := NewGTS("dsa").WithLabels(Labels{
		"ip": "1.2.3.4",
	}).AddDatapoint(ts, "42")
	batch.Register(gts1)

	gts2 := NewGTS("metrics.test2").WithLabels(Labels{
		"ip": "1.2.3.4",
	}).AddDatapoint(ts, 42)
	batch.Register(gts2)

	if len(*batch) != 2 {
		t.Errorf("Expected '%v' got '%v'", 2, len(*batch))
	}

	batch.Print(&buf)

	expected := `1257894000000000// dsa{ip=1.2.3.4} "42"
1257894000000000// metrics.test2{ip=1.2.3.4} 42`

	if !strings.Contains(buf.String(), `1257894000000000// dsa{ip=1.2.3.4} "42"`) && strings.Contains(buf.String(), `1257894000000000// metrics.test2{ip=1.2.3.4} 42"`) {
		t.Errorf("Expected \n'%v'\ngot\n'%v'", expected, buf.String())
	}
}

func TestBatchCreationAddingDatapointsAfter(t *testing.T) {

	var buf bytes.Buffer
	batch := NewBatch()
	batch.Register(singleGTS)
	singleGTS.AddDatapoint(ts, 42)

	batch.Print(&buf)

	if buf.String() != singleGTSSingleDatapointString {
		t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, buf.String())
	}
}

func TestIdentifier(t *testing.T) {
	gts1 := NewGTS("test").WithLabels(Labels{
		"ip": "1.2.3.4",
	})

	gts2 := NewGTS("test2").WithLabels(Labels{
		"ip": "1.2.3.4",
	})

	ident1 := gts1.GetIdentifier()
	ident2 := gts2.GetIdentifier()

	if ident1 == ident2 {
		t.Errorf("ident1='%v', ident2='%v'", ident1, ident2)
	}

}

func TestUberBatch(t *testing.T) {

	var buf bytes.Buffer
	batch1 := NewBatch()
	batch2 := NewBatch()
	batch2.Register(singleGTSSingleDatapoint)

	batch1.RegisterBatch(batch2)
	batch1.Print(&buf)

	if buf.String() != singleGTSSingleDatapointString {
		t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, buf.String())
	}
}
