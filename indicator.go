package main

import "math"

// Tr is true range
func Tr(current, prev Candle) float64 {
	diff1 := current.High - current.Low
	diff2 := math.Abs(current.High - prev.Close)
	diff3 := math.Abs(prev.Close - current.Low)
	return math.Max(math.Max(diff1, diff2), diff3)
}

// Atr is average true range
func Atr(candles []Candle) {

}
