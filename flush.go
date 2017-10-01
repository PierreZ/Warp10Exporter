package Warp10Exporter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

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

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
	}

	// Generating filename
	filename := fmt.Sprintf("%s/%d.warp10", path, time.Now().Unix())
	return ioutil.WriteFile(filepath.Clean(filename), b.Bytes(), 0644)
}
