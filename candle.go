package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	format    = "https://api.bitfinex.com/v2/candles/trade:%v:%v/hist?start=%d&end=%d&limit=%d&sort=%d"
	timeFrame = "1D"
	symbol    = "tBTCUSD"
	limit     = 10
	start     = 1531180800000 // 2018-07-10
	end       = 1531958400000 // 2018-07-19
	sort      = 1             // old to new
)

// Candle is 1 bar chart
type Candle struct {
	MTS    int64
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
}

// NewCandleFromRaw provides initialize of Candle
func NewCandleFromRaw(raw []float64) *Candle {
	return &Candle{
		MTS:    int64(raw[0]), // ms timestamp
		Open:   raw[1],
		Close:  raw[2],
		High:   raw[3],
		Low:    raw[4],
		Volume: raw[5],
	}
}

// GetCandles get chart of candles
func GetCandles() ([]Candle, error) {
	url := fmt.Sprintf(format, timeFrame, symbol, start, end, limit, sort)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var values [][]float64
	err = json.Unmarshal(body, &values)
	if err != nil {
		return nil, err
	}

	candles := make([]Candle, 0)
	for _, v := range values {
		candle := NewCandleFromRaw(v)
		candles = append(candles, *candle)
	}

	return candles, nil
}
