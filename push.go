package Warp10Exporter

import (
	"bytes"
	"net/http"
)

func pushGTS(b *bytes.Buffer, warp10Endpoint string, warp10Token string) int {
	req, err := http.NewRequest("POST", warp10Endpoint+warpURI, b)
	req.Header.Set(warpHeader, warp10Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		return http.StatusInternalServerError
	}
	if err != nil {
		return http.StatusInternalServerError
	}
	defer resp.Body.Close()
	return http.StatusOK
}
