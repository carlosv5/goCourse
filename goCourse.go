// Description
package main

import (
	"fmt"
	"strings"
)

func main() {
	// sign function
	fmt.Printf("1. Sign function results:\n")
	printSign(-5)
	printSign(5)
	printSign(0)

	// sumMul function
	fmt.Printf("\n2. Multiples function results:\n")
	max := 3000000
	fmt.Printf("The sum of multiplies of 3 and 5, but not of 15, from 0 to %d is %d\n", max, sumMul3_5_no15(max))

	// multiply function
	fmt.Printf("\n3. Multiply table function results:\n")
	multiply(5)

	// sumDigits function
	fmt.Printf("\n4. sum digits function results:\n")
	number := 543
	fmt.Printf("The sum of digits of %d is %d\n", number, sumDigits(number))
	fmt.Printf("Is %d a multiple of 3 -> %t\n", number, isMultipleOf3(number))

	// palindrome function
	fmt.Printf("\n6. palindrome function results:\n")
	text := "la ruta nos aporto otro paso natural"
	fmt.Printf("Is '%s' a palindrome word? -> %t\n", text, palindrome(text))
	fmt.Printf("Is '%s' a palindrome word (using switch)? -> %t\n", text, palindromeSwitch(text))
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

// multiply prints the multiplication table of a given number
// Enunciado: Imprime la tabla de multiplicar de cualquier número entregado como argumento
func multiply(n int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d × %d = %d\n", n, i, n*i)
	}
}

// sumDigits sums the digits of a given number until one digit
// Enunciado: Calcula la suma de los dígitos de una función, y usa esa misma función para determinar si un valor dado es múltiplo de 3, o no.
func sumDigits(n int) int {
	/* 	switch {
	   	case n < 10:
	   		return n
	   	default:
	   		return sumDigits(n%10 + sumDigits(n/10))
	   	} */
	if n < 10 {
		return n
	}
	return sumDigits(n%10 + sumDigits(n/10))
}

// isMultipleOf3 indicates if a number is multiple of 3
func isMultipleOf3(n int) bool {
	return sumDigits(n)%3 == 0
}

// palindrome indicates if a given string is a palindrome
// Enunciado: Crea una función que determine si un string es un palíndromo. ¿Serías capaz de hacerlo usando switch?
func palindrome(text string) bool {
	formattedText := strings.ReplaceAll(text, " ", "")
	for i := 0; i < len(formattedText); i++ {
		if formattedText[i] != formattedText[len(formattedText)-1-i] {
			return false
		}
	}
	return true
}

// palindrome indicates if a given string is a palindrome
// Enunciado: Crea una función que determine si un string es un palíndromo. ¿Serías capaz de hacerlo usando switch?
func palindromeSwitch(text string) bool {
	formattedText := strings.ReplaceAll(text, " ", "")
	for i := 0; i < len(formattedText); i++ {
		switch {
		case formattedText[i] == formattedText[len(formattedText)-1-i]:
			continue
		default:
			return false
		}
	}
	return true
}
