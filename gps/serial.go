package gps

import (
  "bytes"
  "log"
  "github.com/tarm/goserial"
)

func ParseNmea(device string,baud int) (nmea Nmea, err error) {
  gga := nmea
  
  conf := &serial.Config{
    Name: device,
    Baud: baud,
  }
  sc, err := serial.OpenPort(conf)
  if err != nil {
    log.Fatal(err)
  }
 
  buffer := bytes.NewBuffer([]byte{})
	for {
		buf := make([]byte, 1)
		_, err := sc.Read(buf)
		if string(buf[0]) == "$" {
			gga, _ = ParseNMEA(buffer.String())
			buffer.Reset()
			return gga, err
		} else {
			buffer.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return gga, err
}