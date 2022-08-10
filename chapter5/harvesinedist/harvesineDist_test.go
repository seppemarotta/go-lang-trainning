package harvesinedist

import (
	"chapter5/harvesinedist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHarvesine(t *testing.T) {
	type args struct {
		a harvesinedist.MapCoord
		b harvesinedist.MapCoord
	}

	tests := map[string]struct {
		args args
		want float64
	}{
		// Athens haversineDist.MapCoord{lat:37.983972,lon: 23.727806}
		// Amsterdam haversineDist.MapCoord{lat:52.366667,lon: 4.9}
		// Berlin haversineDist.MapCoord{lat: 52.516667,lon: 13.388889}
		"Athens-Amsterdam": {args: args{a: harvesinedist.MapCoord{Lat: 37.983972, Lon: 23.727806}, b: harvesinedist.MapCoord{Lat: 52.366667, Lon: 4.9}}, want: 1234.0},
		"Athens-Berlin":    {args: args{a: harvesinedist.MapCoord{Lat: 37.983972, Lon: 23.727806}, b: harvesinedist.MapCoord{Lat: 52.366667, Lon: 4.9}}, want: 5678.0},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := harvesinedist.Distance(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
