package Warp10Exporter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var prefix = "warp10Exporter"

// ChangePrefix is changing the prefix for metrics files.
func ChangePrefix(newprefix string) {
	prefix = newprefix
}

// FlushOnDisk is flushing the metrics into a file
// compatible with the Warp10 Input format.
// You can then use Beamium to handle the push.
func (gts *GTS) FlushOnDisk(path string) error {
	var b bytes.Buffer
	gts.Print(&b)

	return writeFile(&b, path)
}

// FlushOnDisk is flushing the metrics into a file
// compatible with the Warp10 Input format.
// You can then use Beamium to handle the push.
func (batch *Batch) FlushOnDisk(path string) error {
	var b bytes.Buffer
	batch.Print(&b)

	return writeFile(&b, path)
}

func writeFile(b *bytes.Buffer, path string) error {

	// Checking that folder exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
	}

	// Generating filename
	filename := fmt.Sprintf("%s/%s-%d", path, prefix, time.Now().Unix())

	// Writing the file
	err := ioutil.WriteFile(filepath.Clean(filename+".tmp"), b.Bytes(), 0644)
	if err != nil {
		return err
	}

	// rename the file
	err = os.Rename(filename+".tmp", filename+".metrics")
	if err != nil {
		return err
	}

	return nil
}
