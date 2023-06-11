// Description
package main

import (
	"fmt"
	"log"
	"math"
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
	fmt.Printf("\n4. Sum digits function results:\n")
	number := 543
	fmt.Printf("The sum of digits of %d is %d\n", number, sumDigits(number))
	fmt.Printf("Is %d a multiple of 3 -> %t\n", number, isMultipleOf3(number))

	// roots function
	fmt.Printf("\n5. Roots function results:\n")
	a := 1.0
	b := -5.0
	c := 6.0
	roots := roots(a, b, c)
	fmt.Printf("The roots of %.1fx²+%.1fx+%.1f are %.1f and %.1f\n", a, b, c, roots[0], roots[1])

	// palindrome function
	fmt.Printf("\n6. Palindrome function results:\n")
	text := "la ruta nos aporto otro paso natural"
	fmt.Printf("Is '%s' a palindrome word? -> %t\n", text, palindrome(text))
	fmt.Printf("Is '%s' a palindrome word (using switch)? -> %t\n", text, palindromeSwitch(text))

	/* 	// bitmap exercise
	   	fmt.Printf("\n7. Bitmap exercise results:\n")
	   	positions := 30
	   	bitmap := newBitmap(positions)
	   	fmt.Printf("bitMap with, at least, %d positions is:\n", positions)
	   	printBitmap(bitmap)
	   	position := 10
	   	fmt.Printf("Position %d of the bitmap is %b\n", position, readBit(bitmap, position))
	   	changePosition := 0
	   	fmt.Printf("Going to enable position %d. The bitmap is now:\n", changePosition)
	   	newBitMap := enableBit(bitmap, changePosition)
	   	printBitmap(newBitMap) */

	/* 	changePosition = 12
	   	fmt.Printf("Going to enable position %d. The bitmap is now:\n", changePosition)
	   	newBitMap = enableBit(newBitMap, changePosition)
	   	printBitmap(newBitMap)

	   	changePosition = 20
	   	fmt.Printf("Going to enable position %d. The bitmap is now:\n", changePosition)
	   	newBitMap = enableBit(newBitMap, changePosition)
	   	printBitmap(newBitMap) */

	/* 	fmt.Printf("Going to enable all the bits. The bitmap is now:\n")
	   	newBitMap = enableAllBits(newBitMap)
	   	printBitmap(newBitMap) */

	// CSV encoder/decoder
	fmt.Printf("\n CSV Decoder:\n")
	if csvfile, err := openCSV("csv/music.csv"); err != nil {
		log.Fatal(err)
	} else {
		scanner := scan(csvfile)
		for data, err := scanner(); err == nil; data, err = scanner() {
			for _, field := range data {
				switch field.(type) { // Análisis del tipo de cada campo
				case int:
					fmt.Printf("%d(int)", field)
				case float64:
					fmt.Printf("%v(float)", field)
				case string:
					fmt.Printf("%v(string)", field)
				}
			}
			fmt.Println()
		}
	}
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

// roots return the roots of ax²+bx+c
// Enunciado: Crea una función que reciba tres parámetros a, b y c que son los coeficientes de una ecuación de segundo grado: ax^2 + bx + c = 0, y que imprima por pantalla las raíces o soluciones de la ecuación
func roots(a, b, c float64) []float64 {
	x1 := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / 2 * a
	x2 := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / 2 * a
	// math.Cmplx
	return []float64{x1, x2}
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

/* // newBitmap debe devolver un string lo suficientemente largo para alojar length bits
func newBitmap(length int) string {
	neededLength := 0
	if length%8 > 0 {
		neededLength = length/8 + 1
	} else {
		neededLength = length / 8
	}
	var zeroByte byte
	bitmap := strings.Repeat(string(zeroByte), neededLength)
	return bitmap
}

// readBit devuelve el valor del bit en la posición i del bitmap almacenado en el string entregado en primer lugar
func readBit(bitmap string, i int) byte {
	despl := 8 - (i % 8) - 1
	thisByte := (bitmap[i/8] >> despl) & 1
	return thisByte
}

// printBitmap imprime por pantalla los contenidos de todos los bits del bitmap entregado como argumento
func printBitmap(bitmap string) {
	fmt.Printf("-> Printing bitmap:\n")
	for i := 0; i < len(bitmap); i++ {
		subBitMap := string(bitmap[i])
		for j := 0; j < 8; j++ {
			fmt.Printf("%b ", readBit(subBitMap, j))
		}
		fmt.Printf("\n")
	}
}

// enableBit devuelve un bitmap que es una copia del entregado como primer parámetro y que, además, tiene activado el bit i-ésimo
func enableBit(bitmap string, i int) string {
	var str string
	for j := 0; j < len(bitmap); j++ {
		if j == i/8 {
			despl := 8 - (i % 8) - 1
			fmt.Printf("desplazamioento es %d\n", despl)
			bitToChange := (bitmap[i/8] >> despl)
			printBitmap(string(bitToChange))
			bitChanged := bitToChange | 1
			printBitmap(string(bitChanged))
			bitRestored := bitChanged << despl
			printBitmap(string(bitRestored))
			str = str + string(bitRestored)
		} else {
			str = str + string(bitmap[j])
		}
	}
	return str
}

func enableAllBits(bitmap string) string {
	var str string
	for j := 0; j < len(bitmap); j++ {
		str = str + string(bitmap[j]|127)
	}
	return str
} */
