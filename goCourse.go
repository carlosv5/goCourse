// Description
package main

import "fmt"

func main() {
	// Sign function
	printSign(-5)
	printSign(5)
	printSign(0)
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
