package main

import (
	"fmt"
	"math"
)

func pembulatan(n float64) float64 {
	// nType := reflect.TypeOf(n)
	// if k := nType.Kind(); k != reflect.Float64 {
	// 	err := "Input must be a number"
	// 	return (err)
	// } else {
	rounded := math.Floor(n*100) / 100
	return (rounded)
	// }
}

func main() {
	fmt.Println("1. Buatlah sebuah function yang menerima parameter N yang akan dibulatkan ke persepuluhan terdekat")
	fmt.Println("Rounded to the nearest tenth :", pembulatan(4.37))
	fmt.Println("Rounded to the nearest tenth :", pembulatan(4.32))
	fmt.Println("Rounded to the nearest tenth :", pembulatan(4.324))
}
