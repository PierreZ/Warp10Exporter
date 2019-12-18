/*
Package warp10exporter is an helper to create GTS for Warp10
Example:
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

For a full guide visit https://github.com/PierreZ/Warp10Exporter
*/
package warp10exporter
