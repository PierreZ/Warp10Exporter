# Warp10Helper [![Build Status](https://travis-ci.org/PierreZ/Warp10Helper.svg?branch=master)](https://travis-ci.org/PierreZ/Warp10Helper) [![GoDoc](https://godoc.org/github.com/PierreZ/Warp10Helper?status.svg)](https://godoc.org/github.com/PierreZ/Warp10Helper)
Structured warp10 Input Format for Go.


## Example

TODO
```go 
package main
import (
  warp "github.com/PierreZ/Warp10Helper"
)
func main() {
  gts := warp.CreateGTS("metrics.test").WithLabels(warp.Labels{
    "ip": "1.2.3.4",
  }).AddDatapoint(time.Now(), "42")
  warp.Push(gts, "localhost:8080", "WRITE_TOKEN")

  // You can also create batchs
  batch := warp.NewBatch()
  batch.AddGTS(gts)

  warp.PushBatch(batch, "localhost:8080", "WRITE_TOKEN")
}
```
