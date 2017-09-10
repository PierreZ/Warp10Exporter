package Warp10Exporter

import (
	"bytes"
	"fmt"
	"net/http"
)

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
		return fmt.Errorf("Warp10 response status is %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	return nil
}

// Push is pushing a single GTS to a warp10 endpoint
func (gts *GTS) Push(warp10Endpoint string, warp10Token string) error {

	if len(gts.Datapoints) == 0 {
		return nil
	}

	var b bytes.Buffer

	gts.Print(&b)
	return pushGTS(&b, warp10Endpoint, warp10Token)
}