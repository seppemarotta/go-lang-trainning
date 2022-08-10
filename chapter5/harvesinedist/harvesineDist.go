package haversinedist

import (
	"fmt"
	"math"
)

const (
	// radius of the earth in kms
	earthRadKm = 6371
)

type MapCoord struct {
	Lat float64
	Lon float64
}

func degToRad(d float64) float64 {
	return d * math.Pi / 180
}

func Distance(p MapCoord, q MapCoord) float64 {
	lat1 := degToRad(p.Lat)
	lon1 := degToRad(p.Lon)
	lat2 := degToRad(q.Lat)
	lon2 := degToRad(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km := c * earthRadKm
	fmt.Println(km)
	return km
}
