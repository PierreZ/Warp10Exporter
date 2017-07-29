# Warp10Exporter [![Build Status](https://travis-ci.org/PierreZ/Warp10Exporter.svg?branch=master)](https://travis-ci.org/PierreZ/Warp10Exporter) [![GoDoc](https://godoc.org/github.com/PierreZ/Warp10Exporter?status.svg)](https://godoc.org/github.com/PierreZ/Warp10Exporter)
A Go framework to generate and push metrics to Warp10

## Example

```go 
package main

import (
  warp "github.com/PierreZ/Warp10Exporter"
)

func main() {
    
  gts := warp.CreateGTS("metrics.test").WithLabels(warp.Labels{
    "ip": "1.2.3.4",
  }).AddDatapoint(time.Now(), "42")
  warp.PushGTS(gts, "http://localhost:8080", "WRITE_TOKEN")

  // You can also create batchs
  batch := warp.NewBatch()
  batch.AddGTS(gts)

  warp.PushBatch(batch, "http://localhost:8080", "WRITE_TOKEN")
}
```

## TODO

 * Add possibility to flush in files for [Beamium](https://github.com/runabove/beamium)
 * Order Datapoints before writing into buffer to optimize parsing on Warp10 side