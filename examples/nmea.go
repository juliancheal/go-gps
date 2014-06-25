package main

import (
	"bytes"
	"fmt"
	"github.com/juliancheal/go-gps"
	"github.com/tarm/goserial"
	"log"
)

func main() {
	conf := &serial.Config{
		Name: "/dev/tty.SLAB_USBtoUART",
		Baud: 4800,
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

			gga := gps.ParseNMEA(buffer.String())
			fmt.Println(gga)
			buffer.Reset()
		} else {
			buffer.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
