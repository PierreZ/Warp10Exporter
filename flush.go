package Warp10Exporter

import (
	"bytes"
	"io/ioutil"
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
	// Generating filename
	filename := path + "-" + string(time.Now().Nanosecond()) + ".warp10"
	return ioutil.WriteFile(filename, b.Bytes(), 0644)
}