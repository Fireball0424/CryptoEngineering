package main

import (
	"fmt"
	"sort"
)

var ciphertext = "WKHUHDUHWZRZDBVRIFRQVWUXFWLQJDVRIWZDUHGHVLJQRQHZDBLVWRPDNHLWVRVLPSOHWKDWWKHUHDUHREYLRXVOBQRGHILFLHQFLHVDQGWKHRWKHUZDBLVWRPDNHLWVRFRPSOLFDWHGWKDWWKHUHDUHQRREYLRXVGHILFLHQFLHVWKHILUVWPHWKRGLVIDUPRUHGLIILFXOW"

func calculateIC(text string) float64 {
	var length int = len(text)
	var freq [26]int
	var ic float64 = 0.0
	var i int

	for i = 0; i < length; i++ {
		freq[text[i]-'A']++
	}

	for i = 0; i < 26; i++ {
		ic += float64(freq[i] * (freq[i] - 1))
	}

	ic /= float64(length * (length - 1))
	return ic
}

type Alpha struct {
	alphabet string
	freq     float64
}

func frequencyAnalysis(text string) [26]Alpha {
	var freq [26]Alpha
	var i int
	for i = 0; i < 26; i++ {
		freq[i].alphabet = string('A' + i)
		freq[i].freq = 0.0
	}
	for i = 0; i < len(text); i++ {
		freq[text[i]-'A'].freq += 1.0
	}
	for i = 0; i < 26; i++ {
		freq[i].freq /= float64(len(text))
	}
	return freq
}

var freqRank string = "ETAOINSHRDLCUMWFGYPBVKJXQZ"

func main() {
	// calculate IC of cipher text
	var ic float64 = calculateIC(ciphertext)
	fmt.Println("IC of cipher text: ", ic, "\n")

	// try to decrypt the cipher text

	// try 1 : Frequency analysis
	var freq [26]Alpha = frequencyAnalysis(ciphertext)
	sort.Slice(freq[:], func(i, j int) bool {
		return freq[i].freq > freq[j].freq
	})

	// mapping back to the original alphabet by frequency analysis
	var mapping [26]int
	for i := 0; i < 26; i++ {
		mapping[int(freq[i].alphabet[0]-'A')] = int(freqRank[i] - 'A')
	}

	var plaintext string
	for i := 0; i < len(ciphertext); i++ {
		plaintext += string('A' + rune(mapping[int(ciphertext[i]-'A')]))
	}
	fmt.Println("Frequency analysis mapping: ", plaintext, "\n")

	// oberserve that “W → T” , “H → E” and “K → H” might be correct mapping , which is identical to Caesar encryption with shift 3

	// try 2 : Caesar shift decryption
	plaintext = ""
	for i := 0; i < len(ciphertext); i++ {
		plaintext += string('A' + rune((int(ciphertext[i]-'A')+23)%26))
	}
	fmt.Println("Caesar shift: ", plaintext)
}
