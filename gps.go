package main

import (
	"bytes"
	"fmt"
	serial "github.com/tarm/goserial"
	"log"
	"strings"
)

type Nmea struct {
	LastNmea        string
	Time            string
	Latitude        string
	LatRef          string
	Longitude       string
	LongRef         string
	Quality         string
	NumSat          string
	Hdop            string
	Altitude        string
	AltUnit         string
	HeightGeoid     string
	HeightGeoidUnit string
	LastDgps        string
	Dgps            string
}

func latlngToDecimal(coord string, dir string, lat bool) string {
	// decimal := nil
	// if (lat && dir.upcase == "S") || dir.upcase == "W" {
	//   negative := true
	// }
	//
	// if coord =~ /^-?([0-9]*?)([0-9]{2,2}\.[0-9]*)$/ {
	//   deg = $1.to_i // degrees
	//    min = $2.to_f // minutes & seconds
	//
	//    // Calculate
	//    decimal = deg + (min / 60)
	//
	//    if negative {
	//      decimal *= -1
	//    }
	// }
	return ""
}

// $GPGGA,221440.069,3033.2807,N,08126.6636,W,1,05,1.7,-20.0,M,-31.7,M,,0000*72
func parseNMEA(raw string) Nmea {

	line := strings.Split(raw, ",")
	t := strings.Split(line[0], "")

	if 0 < len(t) {
		temp := t[2:5]
		switch strings.Join(temp, "") {
		case "GGA":
			gga := Nmea{
				LastNmea:        line[0],
				Time:            line[1],
				Latitude:        line[2],
				LatRef:          line[3],
				Longitude:       line[4],
				LongRef:         line[5],
				Quality:         line[6],
				NumSat:          line[7],
				Hdop:            line[8],
				Altitude:        line[9],
				AltUnit:         line[10],
				HeightGeoid:     line[11],
				HeightGeoidUnit: line[12],
				LastDgps:        line[13],
				Dgps:            line[14],
			}
			fmt.Println("Latitude", gga.Latitude, gga.LatRef)
			fmt.Println("Longitude", gga.Longitude, gga.LongRef)

			gga.Latitude = latlngToDecimal(gga.Latitude, gga.LatRef, true)
			gga.Longitude = latlngToDecimal(gga.Longitude, gga.LongRef, true)

			return gga
		}
	}
	return Nmea{}
}

func main() {
	conf := new(serial.Config)
	conf.Name = "/dev/tty.SLAB_USBtoUART"
	conf.Baud = 4800

	sc, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.NewBuffer([]byte{})
	for {
		buf := make([]byte, 1)
		_, err := sc.Read(buf)
		if string(buf[0]) == "$" {
			// fmt.Println(buffer.String())
			parseNMEA(buffer.String())
			buffer.Reset()
		} else {
			buffer.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}