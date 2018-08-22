package main

import (
	"fmt"
)

func main() {
	candles, err := GetCandles()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v\n", candles)
}
