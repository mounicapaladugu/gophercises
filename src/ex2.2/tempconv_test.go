package tempconv

import (
	"math"
	"testing"
)

func testTempConv(t *testing.T){
	tests := []struct {
		f Fahrenheit
		c Celsius
	}{
		{68,20,293.15},
		{32,0,273.15},
	}
	eps := 0.0000001
	for _,test := range tests {
		if math.Abs(float64(CToF(test.c)-test.f)) > eps {
			t.Errorf("CToF(%s): got %s, want %s", test.c, CToF(test.c), test.f)
		}
	}
}