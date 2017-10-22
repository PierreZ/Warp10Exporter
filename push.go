package Warp10Exporter

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ErrEmptyGTS is the error thrown when trying to push or write
// an empty GTS
var ErrEmptyGTS = errors.New("Empty GTS")

// ErrEmptyBatch is the error thrown when trying to push or write
// an empty batch of GTS
var ErrEmptyBatch = errors.New("Empty Batch")

func pushGTS(b *bytes.Buffer, warp10Endpoint string, warp10Token string) error {
	req, err := http.NewRequest("POST", warp10Endpoint+warpURI, b)
	if err != nil {
		return err
	}
	req.Header.Set(warpHeader, warp10Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var b []byte
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Warp10 response status is %d, body='%s'", resp.StatusCode, string(b))
	}

	defer resp.Body.Close()
	return nil
}

// Push is pushing a single GTS to a warp10 endpoint
func (gts *GTS) Push(warp10Endpoint string, warp10Token string) error {

	if len(gts.Datapoints) == 0 {
		return ErrEmptyGTS
	}

	var b bytes.Buffer

	gts.Print(&b)
	return pushGTS(&b, warp10Endpoint, warp10Token)
}

// Push is pushing a GTS batch to a warp10 endpoint
func (batch *Batch) Push(warp10Endpoint string, warp10Token string) error {

	if len(*batch) == 0 {
		return ErrEmptyBatch
	}

	var b bytes.Buffer
	i := 0
	for _, gts := range *batch {
		gts.Print(&b)
		if i != len(*batch)-1 {
			b.WriteString("\n")
		}
		i++
	}
	return pushGTS(&b, warp10Endpoint, warp10Token)
}
