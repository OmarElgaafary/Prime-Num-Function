package main

import "fmt"
import "math"

func findPrime(max int) {
	for n := 0; n < max+1; n++ {

		if n == 2 {
			fmt.Printf("%v is a prime number\n", 2)
			continue
		} else if n%2 == 0 {
			fmt.Printf("%v is not a prime, skip to next %v\n", n, n)
			continue
		}
		isPrime := true
		for i := 3; i < int(math.Sqrt(float64(n))+1); i++ {

			if i%n == 0 {
				fmt.Printf("%v is not a prime, skip to next %v\n", i, n)
				isPrime = false
				break
			}

		}

		if !isPrime {
			continue
		} else {
			fmt.Printf("%v is a prime\n", n)
			isPrime = true
		}

	}
}

func main() {

	findPrime(100)

}
