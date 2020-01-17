package vigenere

import (
	"strings"
)

func EncryptVigenere(plaintext string, keyword string) string {
	var ciphertext string
	abcd := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	ABCD := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for len(plaintext) > len(keyword) {
		for  _, value := range keyword {
			if len(plaintext) > len(keyword) {
				keyword = strings.ToLower(keyword) + strings.ToLower(string(value))
			}
		}
	}

	for index, value := range plaintext {
		if (value >= 97 && value <= 122) {
			PlaintextLetter := strings.Index(abcd, string(plaintext[index]))
			KeywordLetter := strings.Index(abcd, string(keyword[index]))
			EncryptedLetter := abcd[PlaintextLetter + KeywordLetter]
			ciphertext += string(EncryptedLetter)
		} else if (value >= 65 && value <= 90) {
			PlaintextLetter := strings.Index(ABCD, string(plaintext[index]))
			KeywordLetter := strings.Index(abcd, string(keyword[index]))
			EncryptedLetter := ABCD[PlaintextLetter + KeywordLetter]
			ciphertext += string(EncryptedLetter)
		} else {
			ciphertext += string(value)
		}
	}
	return ciphertext
}

func DecryptVigenere(ciphertext string, keyword string) string {
	var plaintext string
	abcd := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	ABCD := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for len(ciphertext) > len(keyword) {
		for  _, value := range keyword {
			if len(ciphertext) > len(keyword) {
				keyword = strings.ToLower(keyword) + strings.ToLower(string(value))
			}
		}
	}

	for index, value := range ciphertext {
		if (value >= 97 && value <= 122) {
			PlaintextLetter := strings.Index(abcd, string(ciphertext[index]))
			KeywordLetter := strings.Index(abcd, string(keyword[index]))
			if (PlaintextLetter < KeywordLetter) {
				PlaintextLetter += 52
			}

			EncryptedLetter := abcd[PlaintextLetter - KeywordLetter]
			plaintext += string(EncryptedLetter)
		} else if (value >= 65 && value <= 90) {
			PlaintextLetter := strings.Index(ABCD, string(ciphertext[index]))
			KeywordLetter := strings.Index(abcd, string(keyword[index]))
			if (PlaintextLetter < KeywordLetter) {
				PlaintextLetter += 52
			}
			EncryptedLetter := ABCD[PlaintextLetter - KeywordLetter]
			plaintext += string(EncryptedLetter)
		} else {
			plaintext += string(value)
		}
	}

	return plaintext
}
