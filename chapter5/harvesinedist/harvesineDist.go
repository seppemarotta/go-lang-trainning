package haversinedist

import (
	"fmt"
	"math"
)

const (
	// radius of the earth in kms
	earthRaidusKm = 6371
)

type MapCoord struct {
	Lat float64
	Lon float64
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func Distance(p MapCoord, q MapCoord) (km float64) {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km = c * earthRaidusKm
	fmt.Println(km)
	return km
}

func main() {

	Distance(MapCoord{Lat: 37.983972, Lon: 23.727806}, MapCoord{Lat: 52.366667, Lon: 4.9})

}
