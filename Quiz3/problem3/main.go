package main

import (
	"fmt"
	"os"
)

var freqMap = map[rune]float64{
	'A': 0.082, 'B': 0.015, 'C': 0.028, 'D': 0.043, 'E': 0.13,
	'F': 0.022, 'G': 0.02, 'H': 0.061, 'I': 0.07, 'J': 0.0015,
	'K': 0.0077, 'L': 0.04, 'M': 0.024, 'N': 0.067, 'O': 0.075,
	'P': 0.019, 'Q': 0.00095, 'R': 0.06, 'S': 0.063, 'T': 0.091,
	'U': 0.028, 'V': 0.0098, 'W': 0.024, 'X': 0.0015, 'Y': 0.02,
	'Z': 0.00074,
}

func Decryption(text string, key string) string {
	var result string
	var keyLength int = 4
	for i := 0; i < len(text); i++ {
		var keyIndex int = i % keyLength
		var alphabetIndex int = (int(text[i]-'A') - int(key[keyIndex]-'A') + 26) % 26
		result += string(rune(alphabetIndex + 'A'))
	}
	return result
}

func chiSquare(text string) float64 {
	var freq [26]float64
	var result float64 = 0.0
	for i := 0; i < len(text); i++ {
		freq[text[i]-'A'] += 1.0
	}
	for i := 0; i < 26; i++ {
		freq[i] = freq[i] / float64(len(text))
		result += (freq[i] - freqMap[rune(i+'A')]) * (freq[i] - freqMap[rune(i+'A')]) / freqMap[rune(i+'A')]
	}
	return result
}

func calculateIC(text string) float64 {
	var freq [26]float64
	var result float64 = 0.0

	for i := 0; i < len(text); i++ {
		freq[text[i]-'A'] += 1.0
	}

	for i := 0; i < 26; i++ {
		result += freq[i] * (freq[i] - 1.0)
	}
	result = result / (float64(len(text)) * (float64(len(text)) - 1.0))
	return result
}

func main() {
	// input
	inputText, err := os.ReadFile("problem3Ciphertext.txt")
	if err != nil {
		panic(err)
	}

	// subproblem 1 - calculate possible length of key

	for keyLength := 2; keyLength < 8; keyLength++ {
		var text [8]string

		for i := 0; i < len(inputText); i++ {
			text[i%keyLength] += string(inputText[i])
		}

		var freq [8]float64
		for i := 0; i < keyLength; i++ {
			freq[i] = calculateIC(text[i])
		}

		fmt.Printf("Key length %d: ", keyLength)
		for i := 0; i < keyLength; i++ {
			fmt.Printf("%f ", freq[i])
		}
		fmt.Printf("\n")
	}

	// subproblem 2 - calculate possible key by Chi-square
	// we already know the key length = 4
	var curChiSquareValue float64 = 1000000.0
	var bestKey string = ""
	var bestDecryptedText string = ""

	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			fmt.Printf("Trying key: %c%c%c%c\n", rune(a+'A'), rune(b+'A'), '$', '$')
			for c := 0; c < 26; c++ {
				for d := 0; d < 26; d++ {
					var key string = string(rune(a+'A')) + string(rune(b+'A')) + string(rune(c+'A')) + string(rune(d+'A'))
					var decryptedText string = Decryption(string(inputText), key)
					var chiSquareValue float64 = chiSquare(decryptedText)
					if chiSquareValue < curChiSquareValue {
						curChiSquareValue = chiSquareValue
						bestKey = key
						bestDecryptedText = decryptedText
					}
				}
			}
		}
	}

	// subproblem 3 - output best decrypted text
	fmt.Printf("Best key: %s\n", bestKey)
	fmt.Printf("Decrypted text: %s\n", bestDecryptedText)
}
