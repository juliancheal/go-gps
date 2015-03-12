package gps

import (
	"regexp"
	"strconv"
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

var _nmea = []string{}

func latlngToDecimal(coord string, dir string, lat bool) string {
	decimal := 0.0
	negative := false

	if (lat && strings.ToUpper(dir) == "S") || strings.ToUpper(dir) == "W" {
		negative = true
	}

	r, _ := regexp.Compile("^-?([0-9]*?)([0-9]{2,2}\\.[0-9]*)$")

	result := r.FindStringSubmatch(coord)
	deg, _ := strconv.ParseFloat(result[1], 32) // degrees
	min, _ := strconv.ParseFloat(result[2], 32) // minutes & seconds

	// Calculate
	decimal = deg + (min / 60)

	if negative {
		decimal *= -1
	}

	_decimal := strconv.FormatFloat(decimal, 'g', 'g', 32)
	return _decimal
}

func parser(raw string) (nmea Nmea, err error) {
  err = nil
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
      gga.Latitude = latlngToDecimal(gga.Latitude, gga.LatRef, true)
      gga.Longitude = latlngToDecimal(gga.Longitude, gga.LongRef, true)
      return gga, err
    }
  }
  return Nmea{}, err
}

// $GPGGA,221440.069,3033.2807,N,08126.6636,W,1,05,1.7,-20.0,M,-31.7,M,,0000*72
func ParseNMEA(raw string) (nmea Nmea, err error) {
  gga, err := parser(raw)
  return gga, err
}
