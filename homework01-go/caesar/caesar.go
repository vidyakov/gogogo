package caesar


func EncryptCaesar(plaintext string, shift int) string {
	var ciphertext string

	for _, i := range plaintext {
		if (97 <= i && i <= 119) || (65 <= i && i <= 87) {
			ciphertext = ciphertext + string(i+3)
		} else if (120 <= i && i <= 123) || (88 <= i && i <= 91) {
			ciphertext = ciphertext + string(i-23)
		} else {
			ciphertext = ciphertext + string(i)
		}
	}

	return ciphertext
}

func DecryptCaesar(ciphertext string, shift int) string {
	var plaintext string

	for _, i := range ciphertext {
		if (100 <= i && i <= 123) || (68 <= i && i <= 91) {
			plaintext = plaintext + string(i-3)
		} else if (97 <= i && i <= 99) || (65 <= i && i <= 67) {
			plaintext = plaintext + string(i+23)
		} else {
			plaintext = plaintext + string(i)
		}
	}

	return plaintext
}

