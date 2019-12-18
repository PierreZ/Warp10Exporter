package warp10exporter

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Batch is allowing you to push multiples GTS in a single push
type Batch map[string]*GTS

// NewBatch is creating a batch
func NewBatch() *Batch {
	batch := make(Batch)
	return &batch
}

// Register is adding a GTS to a batch
func (batch *Batch) Register(gts *GTS) {
	if gts == nil {
		return
	}
	(*batch)[gts.GetIdentifier()] = gts
}

// GetIdentifier is returning an identifier for a GTS
// The identifier is useful to handle a map of GTS
func (gts *GTS) GetIdentifier() string {

	sha := sha256.Sum256([]byte(gts.Classname + "{" + gts.getLabels() + "}"))
	return fmt.Sprintf("%x", sha)
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

// RegisterBatch is registering all GTS from a Batch to another batch.
// You need to register first all GTS before registering an Batch
func (batch *Batch) RegisterBatch(newBatch *Batch) {
	if batch == nil {
		return
	}
	for _, gts := range *newBatch {
		batch.Register(gts)
	}
}
