package Warp10Helper

import (
	"bytes"
	"net/http"
)

func pushGTS(b *bytes.Buffer, warp10Endpoint string, warp10Token string) error {
	req, err := http.NewRequest("POST", warp10Endpoint+warpURI, b)
	req.Header.Set(warpHeader, warp10Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		return err
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Push is pushing a single GTS to a warp10 endpoint
func Push(gts *GTS, warp10Endpoint string, warp10Token string) error {
	var b bytes.Buffer
	gts.printGTS(&b)
	return pushGTS(&b, warp10Endpoint, warp10Token)
}

// PushBatch is pushing a GTS batch to a warp10 endpoint
func PushBatch(batch *Batch, warp10Endpoint string, warp10Token string) error {
	var b bytes.Buffer
	for _, gts := range *batch {
		gts.printGTS(&b)
	}

	return pushGTS(&b, warp10Endpoint, warp10Token)
}
