/*
Package Warp10Helper is an helper to create GTS for Warp10
Example:
  package main

  import (
    warp "github.com/PierreZ/Warp10Helper"
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

For a full guide visit https://github.com/PierreZ/Warp10Helper
*/
package Warp10Helper
