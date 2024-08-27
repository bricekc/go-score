// déclaration du package pour ce fichier
package main

// importation du package fmt
import (
	"errors"
	"fmt"
)

func fizzBuzz(a int) {
	for i := 1; i <= a; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func isPrime(a int) bool {
	if a == 2 {
		return true
	} else if a%2 == 0 {
		return false
	}
	for i := 3; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func primeNumbers(a int) []int {
	var intSlice []int
	for i := 2; i <= a; i++ {
		if isPrime(i) {
			intSlice = append(intSlice, i)
		}
	}
	return intSlice
}

func calculator(a float64, b float64, caractere string) (float64, error) {
	switch {
	case b == 0 && caractere == "/":
		return 0, errors.New("division par 0 impossible")
	case caractere == "+":
		return a + b, nil
	case caractere == "-":
		return a - b, nil
	case caractere == "*":
		return a * b, nil
	case caractere == "/":
		return a / b, nil
	default:
		return 0, errors.New("l'opération n'existe pas")
	}
}

func isPalindrome(a uint) bool {
	var arrive uint
	depart := a
	for depart != 0 {
		arrive = arrive * 10 + depart % 10
		depart = depart / 10
	}

	return arrive == a
}

// point d'entrée d'un programme Go
func main() {
	// invocation de la fonction Println déclarée dans le package fmt
	fmt.Println("Test fizzBuzz") // notez la possibilité de passer plusieurs paramètres à Println
	fizzBuzz(15)
	fmt.Println("Test isPrime")
	fmt.Println(isPrime(89))
	fmt.Println("Test primeNumbers")
	fmt.Println(primeNumbers(89))
	fmt.Println("Test calculator")
	fmt.Println(calculator(1, 2, "+"))
	fmt.Println(calculator(1, 2, "*"))
	fmt.Println(calculator(1, 2, "/"))
	fmt.Println(calculator(1, 2, "-"))
	fmt.Println(calculator(1, 0, "/"))
	fmt.Println(calculator(1, 2, "bl"))
	fmt.Println("Test isPalindrome")
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(123))
}
