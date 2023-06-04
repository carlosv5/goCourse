// Description
package main

import "fmt"

func main() {
	// Sign function
	fmt.Printf("1. Sign function results:\n")
	printSign(-5)
	printSign(5)
	printSign(0)

	// sumMul function
	fmt.Printf("\n2. Multiples function results:\n")
	max := 3000000
	fmt.Printf("The sum of multiplies of 3 and 5, but not of 15, from 0 to %d is %d\n", max, sumMul3_5_no15(max))
}

// sign function returns the sign of the argument number. -1 for negatives, 0 for zero, 1 for positives
// Enunciado: Crea una función que devuelva el signo de su argumento (-1 si es negativo, +1 si es positivo, y 0 en cualquier otro caso)
func sign(n int) int {
	switch {
	case n == 0:
		return 0
	case n < 0:
		return -1
	default:
		return 1
	}
}

// printSign prints the result of the call to sign function with a parameter number
func printSign(n int) {
	sign := sign(n)
	fmt.Printf("The sign of %d is %d\n", n, sign)
}

// sumMul3_5_no15 returns the sum of multiple numbers of 3 and 5 but not for 15 until a max number
// Enunciado: Considera la serie aritmética entre 0 y 3.000.000. Calcula la suma de todos los múltiplos de 3 (0, 3, 6, 9, 12, 15, ...) y de 5 (0, 5, 10, 15 ... ) pero que no sean múltiplos de 15
func sumMul3_5_no15(max int) int {
	sum := 0
	for i := 0; i <= max; i++ {
		if (i%3 == 0 || i%5 == 0) && (i%15 != 0) {
			sum = sum + i
		}
	}
	return sum
}
