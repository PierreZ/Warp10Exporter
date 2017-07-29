package Warp10Helper

// Batch is allowing you to push multiples GTS in a single push
type Batch map[string]GTS

// NewBatch is creating a batch
func NewBatch() *Batch {
	batch := make(Batch)
	return &batch
}

// AddGTS is adding a GTS to a batch
func (batch *Batch) AddGTS(gts *GTS) {
	(*batch)[gts.getIdentifier()] = *gts
}
