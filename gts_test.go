package Warp10Exporter

import (
	"bytes"
	"testing"
	"time"
)

var ts = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
var classname = "test"
var labels = Labels{
	"ip": "1.2.3.4",
}

var singleGTS = NewGTS(classname).WithLabels(labels)

var singleGTSSingleDatapoint = NewGTS(classname).WithLabels(labels).AddDatapoint(ts, 42)
var singleGTSSingleDatapointString = "1257894000000000// test{ip=1.2.3.4} 42"

var singleGTSMultipleDatapoints = NewGTS(classname).WithLabels(labels).AddDatapoint(ts, 42).AddDatapoint(ts.Add(time.Duration(1)*time.Millisecond), "43")
var singleGTSMultipleDatapointsString = `1257894000000000// test{ip=1.2.3.4} 42
1257894000001000// test{ip=1.2.3.4} "43"`

var singleGTSMultipleTypes = NewGTS(classname).WithLabels(labels).AddDatapoint(ts, 42).AddDatapoint(ts.Add(time.Duration(1)*time.Second), "43").AddDatapoint(ts.Add(time.Duration(2)*time.Second), true)
var singleGTSMultipleTypesString = `1257894000000000// test{ip=1.2.3.4} 42
1257894001000000// test{ip=1.2.3.4} "43"
1257894002000000// test{ip=1.2.3.4} T`

func TestNewGTS(t *testing.T) {

	var b bytes.Buffer
	singleGTSSingleDatapoint.Print(&b)
	if b.String() != singleGTSSingleDatapointString {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", singleGTSSingleDatapointString, b.String())
	}

	b.Reset()
	singleGTSMultipleDatapoints.Print(&b)
	if b.String() != singleGTSMultipleDatapointsString {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", singleGTSMultipleDatapointsString, b.String())
	}

	b.Reset()
	singleGTSMultipleTypes.Print(&b)
	if b.String() != singleGTSMultipleTypesString {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", singleGTSMultipleTypesString, b.String())
	}
}

func TestAddLabel(t *testing.T) {
	gts := NewGTS(classname).AddDatapoint(ts, 42)
	gts.AddLabel("unicorn", "my-little-poney")

	var b bytes.Buffer
	gts.Print(&b)

	expected := "1257894000000000// test{unicorn=my-little-poney} 42"
	if b.String() != expected {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", expected, b.String())
	}

	b.Reset()

	gtsURLEncode := NewGTS(",{}=").AddDatapoint(ts, 42)
	gtsURLEncode.AddLabel(",}=", ",}=")
	gtsURLEncode.Print(&b)

	expected = "1257894000000000// %2C%7B%7D%3D{%2C%7D%3D=%2C%7D%3D} 42"
	if b.String() != expected {
		t.Fatalf("Expected \n'%v'\ngot\n'%v'", expected, b.String())
	}
}
