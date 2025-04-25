package vigenerecipher

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"math"
	"reflect"
	"testing"
)

func TestValuesFloat_sumSquaredError(t *testing.T) {
	tests := []struct {
		name     string
		vf1      ValuesFloat
		vf2      ValuesFloat
		expected float64
	}{
		{
			name: "Identical values",
			vf1: ValuesFloat{
				'a': 0.1, 'b': 0.2, 'c': 0.3,
			},
			vf2: ValuesFloat{
				'a': 0.1, 'b': 0.2, 'c': 0.3,
			},
			expected: 0.0,
		},
		{
			// test negative values
			name: "Slightly different values",
			vf1: ValuesFloat{
				'a': 0.1, 'b': 0.2, 'c': 0.4,
			},
			vf2: ValuesFloat{
				'a': 0.2, 'b': 0.1, 'c': 0.3,
			},
			expected: 0.03, // (0.1^2 + 0.1^2 + 0.1^2)
		},
		{
			name:     "Empty maps",
			vf1:      ValuesFloat{},
			vf2:      ValuesFloat{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.vf1.sumSquaredError(tt.vf2)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestNewFrequency(t *testing.T) {
	freq := NewFrequency("test")
	wantValues := Values{
		'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0,
		'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0,
		'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0,
		'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0,
	}

	if !reflect.DeepEqual(freq.Values, wantValues) {
		t.Errorf("expected Values to be %v, got %v", wantValues, freq.Values)
	}
}

func TestToFractions(t *testing.T) {
	freq := NewFrequency("test")

	// Modify some values in the map
	freq.Values['a'] = 5
	freq.Values['b'] = 15
	freq.Values['c'] = 10

	gotFractions := freq.ToFractions()

	total := float32(5 + 15 + 10) // 30
	wantFractions := ValuesFloat{
		'a': 5.0 / total,
		'b': 15.0 / total,
		'c': 10.0 / total,
		'd': 0.0, 'e': 0.0, 'f': 0.0, 'g': 0.0,
		'h': 0.0, 'i': 0.0, 'j': 0.0, 'k': 0.0, 'l': 0.0, 'm': 0.0,
		'n': 0.0, 'o': 0.0, 'p': 0.0, 'q': 0.0, 'r': 0.0, 's': 0.0,
		't': 0.0, 'u': 0.0, 'v': 0.0, 'w': 0.0, 'x': 0.0, 'y': 0.0, 'z': 0.0,
	}

	// Leverage cmp for comparing float32 with tolerance
	opts := cmp.Options{
		cmpopts.EquateApprox(0.00000001, 0), // Set tolerance for float comparison
	}

	// Compare maps with tolerance
	if diff := cmp.Diff(wantFractions, gotFractions, opts); diff != "" {
		t.Errorf("ValuesFloat mismatch (-want +got):\n%s", diff)
	}

}
