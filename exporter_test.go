package Warp10Exporter

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var ts = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
var classname = "test"
var labels = Labels{
	"ip": "1.2.3.4",
}

var emptyDatapoint = NewGTS("test").WithLabels(labels)
var singleDatapoint = NewGTS("test").WithLabels(labels).AddDatapoint(ts, 42)
var singleDatapointString = "1257894000000000// test{ip=1.2.3.4} 42"

var multipleDatapoints = NewGTS("test").WithLabels(labels).AddDatapoint(ts, 42).AddDatapoint(ts.Add(time.Duration(1)*time.Millisecond), "43")
var multipleDatapointsString = `1257894000000000// test{ip=1.2.3.4} 42
1257894000000000// test{ip=1.2.3.4} "43"`

func TestGTSCreation(t *testing.T) {

	var b bytes.Buffer
	singleDatapoint.Print(&b)
	if b.String() != singleDatapointString {
		t.Fatalf("Expected '%v', got '%v'", singleDatapointString, b.String())
	}

	b.Reset()
	multipleDatapoints.Print(&b)
	if b.String() != multipleDatapointsString {
		t.Fatalf("Expected '%v', got '%v'", multipleDatapointsString, b.String())
	}
}

func TestGTSPush(t *testing.T) {
	internalServerError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}))
	defer internalServerError.Close()

	singleGTSValidatorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != singleDatapointString {
			t.Errorf("Expected '%v', got '%v'", singleDatapointString, string(body))
		}
	}))
	defer singleGTSValidatorServer.Close()

	statuscode := singleDatapoint.Push(internalServerError.URL, "abcd")
	if statuscode != http.StatusInternalServerError {
		t.Errorf("Expected '%v', got '%v'", http.StatusInternalServerError, statuscode)
	}

	statuscode = singleDatapoint.Push(singleGTSValidatorServer.URL, "abcd")
	if statuscode != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, statuscode)
	}
}

func TestBatchCreation(t *testing.T) {

	var buf bytes.Buffer
	batch := NewBatch()
	batch.Register(singleDatapoint)

	batch.Print(&buf)

	if buf.String() != singleDatapointString {
		t.Errorf("Expected '%v', got '%v'", singleDatapointString, buf.String())
	}
}

func TestBatchCreationAddingDatapointsAfter(t *testing.T) {

	var buf bytes.Buffer
	batch := NewBatch()
	batch.Register(emptyDatapoint)
	emptyDatapoint.AddDatapoint(ts, 42)

	batch.Print(&buf)

	if buf.String() != singleDatapointString {
		t.Errorf("Expected '%v', got '%v'", singleDatapointString, buf.String())
	}
}

func TestBatchPush(t *testing.T) {

	batch := NewBatch()
	gts := NewGTS("test").WithLabels(labels)
	batch.Register(gts)
	gts.AddDatapoint(ts, 42)

	singleGTSValidatorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != singleDatapointString {
			t.Errorf("Expected '%v', got '%v'", singleDatapointString, string(body))
		}
	}))
	defer singleGTSValidatorServer.Close()

	statuscode := batch.Push(singleGTSValidatorServer.URL, "abcd")
	if statuscode != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, statuscode)
	}
}
