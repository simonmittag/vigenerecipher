package vigenerecipher

import (
	"bytes"
	"strings"
	"testing"
)

func TestVigenereShift(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		input     string
		decrypt   bool
		want      string
		expectErr bool
	}{
		{
			name:     "Encrypt with 'A' password",
			password: "A",
			input:    "HELLO",
			decrypt:  false,
			want:     "HELLO", // With 'A' password (shift 0), the output should be the same as input
		},
		{
			name:     "Encrypt with 'B' password",
			password: "B",
			input:    "HELLO",
			decrypt:  false,
			want:     "IFMMP", // Shift by 1
		},
		{
			name:     "Encrypt with 'AAA' password",
			password: "BBB",
			input:    "ZZZ",
			decrypt:  false,
			want:     "AAA", // Shifts by 0, 1, 2 repeating, only for alphabetic characters
		},
		{
			name:     "Encrypt with 'ABC' password",
			password: "ABC",
			input:    "HELLO WORLD",
			decrypt:  false,
			want:     "HFNLP YOSND", // Shifts by 0, 1, 2 repeating, only for alphabetic characters
		},
		{
			name:     "Decrypt with 'ABC' password",
			password: "ABC",
			input:    "HFNLP YOSND",
			decrypt:  true,
			want:     "HELLO WORLD",
		},
		{
			name:     "Encrypt with 'BED' password",
			password: "BED",
			input:    "ABCDE",
			decrypt:  false,
			want:     "BFFEI",
		},
		{
			name:     "SHE SEES SEASHELLS NEAR THE SEASHORE",
			password: "BED",
			input:    "SHE SEES SEASHELLS NEAR THE SEASHORE",
			decrypt:  false,
			want:     "TLH TIHT WHBWKFPOT RHBV WII VFEVISUF",
		},
		{
			name:     "Encrypt with mixed case password",
			password: "AbC",
			input:    "Hello World!",
			decrypt:  false,
			want:     "Hfnlp Yosnd!", // Shifts by 0, 1, 2 repeating, only for alphabetic characters
		},
		{
			name:     "Decrypt with mixed case password",
			password: "AbC",
			input:    "Hfnlp Yosnd!",
			decrypt:  true,
			want:     "Hello World!",
		},
		{
			name:     "Encrypt multiline text",
			password: "KEY",
			input: `This is a test.
Vigenere cipher!`,
			decrypt: false,
			want: `Dlgc mq k xccx.
Tskcxipo ggzlcb!`,
		},
		{
			name:     "Decrypt multiline text",
			password: "KEY",
			input: `Dlgc mq k xccx.
Tskcxipo ggzlcb!`,
			decrypt: true,
			want: `This is a test.
Vigenere cipher!`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &bytes.Buffer{}

			cipher := NewVigenereCipher(tc.password)
			err := cipher.Shift(reader, writer, tc.decrypt)

			if tc.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			got := writer.String()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestVigenereRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		password string
		input    string
	}{
		{
			name:     "Simple test with 'A' password",
			password: "A",
			input:    "HELLO",
		},
		{
			name:     "Test with 'ABC' password",
			password: "ABC",
			input:    "HELLO WORLD",
		},
		{
			name:     "Test with mixed case password",
			password: "AbC",
			input:    "Hello World!",
		},
		{
			name:     "Test with longer password",
			password: "SECRETKEY",
			input:    "This is a test of the Vigenere cipher.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Vigenere cipher with the test password
			cipher := NewVigenereCipher(tt.password)

			// Encrypt the input
			inputReader := strings.NewReader(tt.input)
			encryptedBuffer := new(bytes.Buffer)
			err := cipher.Shift(inputReader, encryptedBuffer, false)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			// Get the encrypted result
			encrypted := encryptedBuffer.String()

			// Make sure the encrypted text is different from the input (unless password is all 'A's)
			if encrypted == tt.input && tt.password != "A" {
				t.Errorf("Encryption did not change the input: %s", encrypted)
			}

			// Decrypt the encrypted text
			encryptedReader := strings.NewReader(encrypted)
			decryptedBuffer := new(bytes.Buffer)
			err = cipher.Shift(encryptedReader, decryptedBuffer, true)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}

			// Get the decrypted result
			decrypted := decryptedBuffer.String()

			// Verify that the decrypted text matches the original input
			if decrypted != tt.input {
				t.Errorf("Decryption failed to recover the original input.\nExpected: %s\nGot: %s", tt.input, decrypted)
			}
		})
	}
}
