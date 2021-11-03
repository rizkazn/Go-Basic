package main

import (
	"fmt"
)

type deretBilang struct {
	N int
}

// deret bilangan prima
func (d deretBilang) prima() []int {
	primes := make([]int, 0, d.N)
	for i := 2; i < d.N; i++ {
		isPrime := true
		for _, j := range primes {
			if i%j == 0 {
				isPrime = false
				continue
			}
		}
		if isPrime == true {
			primes = append(primes, i)
		}
	}
	return (primes)
}

// deret bilangan ganjil
func (d deretBilang) ganjil() []int {
	odds := make([]int, 0, d.N)
	for i := 0; i < d.N; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		} else {
			continue
		}
	}
	return (odds)
}

// deret bilangan genap
func (d deretBilang) genap() []int {
	evens := make([]int, 0, d.N)
	for i := 0; i < d.N; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		} else {
			continue
		}
	}
	return (evens)
}

// deret bilangan fibonacci
func (d deretBilang) fibonacci() []int {
	fibos := make([]int, d.N+1, d.N+2)
	if d.N < 2 {
		fibos = fibos[0:2]
	}
	fibos[0] = 0
	fibos[1] = 1
	for i := 2; i <= d.N; i++ {
		fibos[i] = fibos[i-2] + fibos[i-1]
	}
	return (fibos)
}

func main() {
	var deret = deretBilang{40}

	fmt.Println("2. Buatlah sebuah struct berisi method yang mengolah object struct N sebagai limit untuk mencetak deret bilangan prima, ganjil/genap, dan fibonacci")
	fmt.Println("Deret Bilangan Prima :", deret.prima())
	fmt.Println("Deret Bilangan Ganjil :", deret.ganjil())
	fmt.Println("Deret Bilangan Genap :", deret.genap())
	fmt.Println("Deret Fibonacci :", deret.fibonacci())
}
