package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Soal 1: Manipulasi String
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Soal 2: Menghitung Karakter
func countCharacters(input string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range input {
		charCount[char]++
	}
	return charCount
}

// Soal 3: Validasi Password
func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasLowerCase, hasUpperCase, hasDigit, hasSpecialChar bool

	for _, char := range password {
		if unicode.IsLower(char) {
			hasLowerCase = true
		} else if unicode.IsUpper(char) {
			hasUpperCase = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else if strings.ContainsAny(string(char), "@#$") {
			hasSpecialChar = true
		}
	}

	return hasLowerCase && hasUpperCase && hasDigit && hasSpecialChar
}

// Soal 4: Menghitung Bilangan Prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Soal 5: Menggabungkan Dua Array
func mergeArrays(arr1 []int, arr2 []int) []int {
	merged := append(arr1, arr2...)
	return merged
}

// Soal 6: Menghitung Nilai Rata-rata
func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// Soal 7: Mengurutkan Array
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Soal 8: Menghitung Faktorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Soal 9: Validasi Email
func validateEmail(email string) bool {
	// Sederhana: validasi format email dengan memeriksa adanya '@'
	return strings.Contains(email, "@")
}

// Soal 10: Menghitung Jumlah Kata dalam Kalimat
func countWords(sentence string) int {
	words := strings.Fields(sentence)
	return len(words)
}

func main() {
	// Contoh penggunaan fungsi-fungsi di atas
	fmt.Println("Soal 1:")
	fmt.Println(reverseString("Hello, World!"))

	fmt.Println("\nSoal 2:")
	fmt.Println(countCharacters("programming"))

	fmt.Println("\nSoal 3:")
	fmt.Println(validatePassword("P@ssw0rd"))
	fmt.Println(validatePassword("abc123"))

	fmt.Println("\nSoal 4:")
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(4))

	fmt.Println("\nSoal 5:")
	fmt.Println(mergeArrays([]int{1, 2, 3}, []int{4, 5, 6}))

	fmt.Println("\nSoal 6:")
	fmt.Println(calculateAverage([]int{1, 2, 3, 4, 5}))

	fmt.Println("\nSoal 7:")
	fmt.Println(bubbleSort([]int{5, 1, 4, 2, 8}))

	fmt.Println("\nSoal 8:")
	fmt.Println(factorial(5))

	fmt.Println("\nSoal 9:")
	fmt.Println(validateEmail("user@example.com"))
	fmt.Println(validateEmail("invalid_email"))

	fmt.Println("\nSoal 10:")
	fmt.Println(countWords("Ini adalah contoh kalimat."))
}
