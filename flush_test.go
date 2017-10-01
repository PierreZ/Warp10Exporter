package Warp10Exporter

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGTSFlushOnDisk(t *testing.T) {

	dir, err := ioutil.TempDir("", "sink")
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir) // clean up

	err = singleGTSSingleDatapoint.FlushOnDisk(dir)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		if string(content) != singleGTSSingleDatapointString {
			t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, string(content))
		}
	}
}

func TestBatchFlushOnDisk(t *testing.T) {

	dir, err := ioutil.TempDir("", "sink")
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir) // clean up

	batch := NewBatch()
	batch.Register(singleGTSSingleDatapoint)

	err = batch.FlushOnDisk(dir)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		if string(content) != singleGTSSingleDatapointString {
			t.Errorf("Expected '%v', got '%v'", singleGTSSingleDatapointString, string(content))
		}
	}
}
