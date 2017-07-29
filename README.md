# Warp10Exporter [![Build Status](https://travis-ci.org/PierreZ/Warp10Exporter.svg?branch=master)](https://travis-ci.org/PierreZ/Warp10Exporter) [![GoDoc](https://godoc.org/github.com/PierreZ/Warp10Exporter?status.svg)](https://godoc.org/github.com/PierreZ/Warp10Exporter)
Structured warp10 Input Format for Go.


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
  warp.Push(gts, "http://localhost:8080", "WRITE_TOKEN")

  // You can also create batchs
  batch := warp.NewBatch()
  batch.AddGTS(gts)

  warp.PushBatch(batch, "http://localhost:8080", "WRITE_TOKEN")
}
```
