package vigenerecipher

import (
	"bytes"
	"io"
	"unicode"
)

type VigenereCipher struct {
	Password string
}

func NewVigenereCipher(password string) *VigenereCipher {
	return &VigenereCipher{Password: password}
}

// Shift implements encryption and decryption for the Vigenere cipher.
// It applies a different Caesar cipher to each character in the input,
// with the shift determined by the corresponding character in the password.
func (v *VigenereCipher) Shift(reader io.Reader, writer io.Writer, decrypt bool) error {
	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	contentStr := string(content)
	result := new(bytes.Buffer)

	passIndex := 0

	for _, char := range contentStr {
		// we skip everything that's not in the roman alphabet and disregard it for the shift.
		if !unicode.IsLetter(char) {
			result.WriteRune(char)
			continue
		}

		passChar := rune(v.Password[passIndex%len(v.Password)])
		passIndex++

		var shift int
		if passChar >= 'A' && passChar <= 'Z' {
			shift = int(passChar - 'A')
		} else if passChar >= 'a' && passChar <= 'z' {
			shift = int(passChar - 'a')
		} else {
			// For non-alphabetic characters, use 0 as shift
			shift = 0
		}

		if decrypt {
			shift = -shift
		}

		var base rune
		if unicode.IsUpper(char) {
			base = 'A'
		} else {
			base = 'a'
		}

		newChar := (char-base+rune(shift)+26)%26 + base
		result.WriteRune(newChar)
	}

	_, err = writer.Write(result.Bytes())
	return err
}
