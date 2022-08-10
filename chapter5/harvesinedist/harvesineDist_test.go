package haversinedist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHarvesine(t *testing.T) {
	type args struct {
		a MapCoord
		b MapCoord
	}

	tests := map[string]struct {
		args args
		want float64
	}{
		// Athens haversineDist.MapCoord{Lat:37.983972,Lon: 23.727806}
		// Amsterdam haversineDist.MapCoord{Lat:52.366667,Lon: 4.9}
		// Berlin haversineDist.MapCoord{Lat: 52.516667,Lon: 13.388889}
		"Athens-Amsterdam": {args: args{a: MapCoord{Lat: 37.983972, Lon: 23.727806}, b: MapCoord{Lat: 52.366667, Lon: 4.9}}, want: 2163.2310285824487},
		"Athens-Berlin":    {args: args{a: MapCoord{Lat: 37.983972, Lon: 23.727806}, b: MapCoord{Lat: 52.516667, Lon: 13.388889}}, want: 1803.1087879059257},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Distance(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
