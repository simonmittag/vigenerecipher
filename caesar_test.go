package vigenerecipher

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkShiftWithOffset(b *testing.B) {
	cc := NewCaesarCipher(0)
	baseString := strings.Repeat("A", 1024)
	shift := 3

	var tests []struct {
		name  string
		input string
	}

	for i := 0; i < 16; i++ {
		size := 1 << i // 2^i
		tests = append(tests, struct {
			name  string
			input string
		}{
			name:  fmt.Sprintf("%d chars", size*len(baseString)), // Calculate total characters
			input: strings.Repeat(baseString, size),              // Repeat baseString to get the size
		})
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cc.ShiftWithOffset(test.input, shift)
			}
		})
	}
}

func TestShiftWithOffset(t *testing.T) {
	cc := NewCaesarCipher(0)

	tests := []struct {
		input    string
		shift    int
		expected string
	}{
		{"ABC", 3, "DEF"},
		{"XYZ", 3, "ABC"},
		{"abc", 3, "def"},
		{"xyz", 3, "abc"},
		{"ABC", -3, "XYZ"},
		{"XYZ", -3, "UVW"},
		{"abc", -3, "xyz"},
		{"xyz", -3, "uvw"},
		{"Hello, World!", 5, "Mjqqt, Btwqi!"},
		{"", 10, ""},
		{"123", 5, "123"},
	}

	for _, test := range tests {
		result := cc.ShiftWithOffset(test.input, test.shift)
		if result != test.expected {
			t.Errorf("ShiftWithOffset(%q, %d) = %q; want %q", test.input, test.shift, result, test.expected)
		}
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		name      string
		offset    int
		input     string
		decrypt   bool
		want      string
		expectErr bool
	}{
		{
			name:   "Encrypt multiline text",
			offset: 2,
			input: `Hello, World!
This is a test.
Caesar cipher!`,
			decrypt: false,
			want: `Jgnnq, Yqtnf!
Vjku ku c vguv.
Ecguct ekrjgt!`,
		},
		{
			name:   "Decrypt multiline text",
			offset: 2,
			input: `Jgnnq, Yqtnf!
Vjku ku c vguv.
Ecguct ekrjgt!`,
			decrypt: true,
			want: `Hello, World!
This is a test.
Caesar cipher!`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &bytes.Buffer{}

			cipher := NewCaesarCipher(tc.offset)
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
			if strings.TrimSpace(got) != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
