package Warp10Exporter

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

// Batch is allowing you to push multiples GTS in a single push
type Batch map[string]*GTS

var ErrGTSNotFound = errors.New("GTS not found in Batch")

// NewBatch is creating a batch
func NewBatch() *Batch {
	batch := make(Batch)
	return &batch
}

// AddGTS is adding a GTS to a batch
func (batch *Batch) Register(gts *GTS) {
	(*batch)[gts.GetIdentifier()] = gts
}

// GetIdentifier is returning an identifier for a GTS
// The identifier is useful to handle a map of GTS
func (gts *GTS) GetIdentifier() string {

	var md5Hash = md5.New()
	io.WriteString(md5Hash, gts.Classname+"{"+gts.getLabels()+"}")
	return fmt.Sprintf("%x", md5.Sum(nil))
}

// Print is priting a batch of GTS
func (batch *Batch) Print(b *bytes.Buffer) {
	i := 0
	for _, gts := range *batch {
		gts.Print(b)
		if i != len(*batch)-1 {
			b.WriteString("\n")
		}
		i++
	}
}

// Push is pushing a GTS batch to a warp10 endpoint
func (batch *Batch) Push(warp10Endpoint string, warp10Token string) int {
	var b bytes.Buffer
	for _, gts := range *batch {
		gts.Print(&b)
	}
	return pushGTS(&b, warp10Endpoint, warp10Token)
}
