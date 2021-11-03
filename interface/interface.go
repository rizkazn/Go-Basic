package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type kubus struct {
	sisi float64
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func main() {
	var bangunRuang = &kubus{10}

	fmt.Println("3. ELEMEN KUBUS")
	fmt.Println("Luas :", bangunRuang.luas())
	fmt.Println("Keliling :", bangunRuang.keliling())
	fmt.Println("Volume :", bangunRuang.volume())
}
