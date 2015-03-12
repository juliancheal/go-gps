package main

import (
  "fmt"
  // "github.com/juliancheal/go-gps"
  "../gps"
)

func main() {
  for {
    gga, _ := gps.ParseNmea("/dev/tty.SLAB_USBtoUART", 4800)
    lat := gga.Latitude
    lon := gga.Longitude
    if lat != "" {
      fmt.Println(lat)
      fmt.Println(lon)
    }
  }
}
