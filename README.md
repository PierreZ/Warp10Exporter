# Warp10Exporter [![Build Status](https://travis-ci.org/PierreZ/Warp10Exporter.svg?branch=master)](https://travis-ci.org/PierreZ/Warp10Exporter) [![GoDoc](https://godoc.org/github.com/PierreZ/Warp10Exporter?status.svg)](https://godoc.org/github.com/PierreZ/Warp10Exporter)
A Go framework to generate and push metrics to Warp10

## Example

```go 
package main

import (
	"time"

	warp "github.com/PierreZ/Warp10Exporter"
)

func main() {

	gts := warp.NewGTS("metrics.test").WithLabels(warp.Labels{
		"ip": "1.2.3.4",
	}).AddDatapoint(time.Now(), "42")
	// Not checking the error
	gts.Push("http://localhost:8080", "WRITE_TOKEN")

	// You can also create batchs
	batch := warp.NewBatch()
	batch.Register(gts)
	gts.AddDatapoint(ts, 42)

	err := batch.Push("http://localhost:8080", "WRITE_TOKEN")
	if err != nil {
		// You can also write metrics to a file, to use 
		// https://github.com/ovh/beamium for example
		err = batch.FlushOnDisk("/opt/beamium/sink")
		if err != nil {
			panic(err)
		}
	}
}

```

## TODO

 * Order Datapoints before writing into buffer to optimize parsing on Warp10 side
 * Add Geo on Datapoints by creating NewGeoDatapoints(...)