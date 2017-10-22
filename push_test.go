package Warp10Exporter

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGTSPush(t *testing.T) {
	internalServerError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}))
	defer internalServerError.Close()

	singleGTSValidatorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != singleGTSSingleDatapointString {
			t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, string(body))
		}
	}))
	defer singleGTSValidatorServer.Close()

	err := singleGTSSingleDatapoint.Push(internalServerError.URL, "abcd")
	if !strings.Contains(err.Error(), "Internal Server Error") {
		t.Errorf("Expected 'Internal Server Error', got '%v'", err)
	}

	err = singleGTSSingleDatapoint.Push(singleGTSValidatorServer.URL, "abcd")
	if err != nil {
		t.Errorf("Expected 'nil', got '%v'", err)
	}

	err = singleGTSSingleDatapoint.Push("256.256.256.256:9091", "abcd")
	expected := errors.New("parse 256.256.256.256:9091/api/v0/update: first path segment in URL cannot contain colon")
	if err.Error() != expected.Error() {
		t.Errorf("Expected '%v', got '%v'", expected, err)
	}

	err = singleGTSSingleDatapoint.Push("", "abcd")
	expected = errors.New("Post /api/v0/update: unsupported protocol scheme \"\"")
	if err.Error() != expected.Error() {
		t.Errorf("Expected '%v', got '%v'", expected, err)
	}

	gts := NewGTS("dsa")
	gts.Push("", "abcd")
}

func TestBatchPush(t *testing.T) {

	batch := NewBatch()
	gts := NewGTS("test").WithLabels(labels)
	batch.Register(gts)
	gts.AddDatapoint(ts, 42)

	singleGTSValidatorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != singleGTSSingleDatapointString {
			t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, string(body))
		}
	}))
	defer singleGTSValidatorServer.Close()

	err := batch.Push(singleGTSValidatorServer.URL, "abcd")
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}
func TestEmptyBatchPush(t *testing.T) {
	batch := NewBatch()
	err := batch.Push("", "abcd")
	if err != ErrEmptyBatch {
		t.Errorf("Expected '%v', got '%v'", ErrEmptyBatch, err)
	}
}

func TestEmptyGTSPush(t *testing.T) {
	gts := NewGTS("test").WithLabels(labels)
	err := gts.Push("", "abcd")
	if err != ErrEmptyGTS {
		t.Errorf("Expected '%v', got '%v'", ErrEmptyGTS, err)
	}
}
