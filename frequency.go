package vigenerecipher

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type Values map[rune]int
type ValuesFloat map[rune]float32

func (vf ValuesFloat) sumSquaredError(vf2 ValuesFloat) float64 {
	sum := 0.0
	for k, v := range vf {
		vf2Value, ok := vf2[k]
		if !ok {
			vf2Value = 0.0
		}
		sum += math.Pow(float64(v-vf2Value), 2)
	}
	return sum
}

type Frequency struct {
	Name   string
	Values Values
}

type FrequencyFloat struct {
	Name   string
	Values ValuesFloat
}

// MarshalJSON ensures rune-to-string conversion for JSON storage.
func (vf ValuesFloat) MarshalJSON() ([]byte, error) {
	// Create a temporary map for JSON encoding where keys are strings
	temp := make(map[string]float32)
	for k, v := range vf {
		temp[string(k)] = v
	}
	return json.Marshal(temp)
}

// UnmarshalJSON ensures string-to-rune conversion for JSON parsing.
func (vf *ValuesFloat) UnmarshalJSON(data []byte) error {
	// Decode into a temporary map where keys are strings
	temp := make(map[string]float32)
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Convert string keys back to rune
	*vf = make(ValuesFloat)
	for k, v := range temp {
		if len(k) != 1 {
			return fmt.Errorf("invalid key %q (expected single character for rune)", k)
		}
		(*vf)[[]rune(k)[0]] = v
	}
	return nil
}

func NewFrequency(name string) *Frequency {
	return &Frequency{
		Values: map[rune]int{
			'a': 0,
			'b': 0,
			'c': 0,
			'd': 0,
			'e': 0,
			'f': 0,
			'g': 0,
			'h': 0,
			'i': 0,
			'j': 0,
			'k': 0,
			'l': 0,
			'm': 0,
			'n': 0,
			'o': 0,
			'p': 0,
			'q': 0,
			'r': 0,
			's': 0,
			't': 0,
			'u': 0,
			'v': 0,
			'w': 0,
			'x': 0,
			'y': 0,
			'z': 0,
		},
	}
}

func (f *Frequency) Merge(f2 Frequency) {
	for k, v := range f2.Values {
		//we only merge existing keys
		if _, ok := f.Values[k]; ok {
			f.Values[k] += v
		}
	}
}

func (f *Frequency) ToFractions() ValuesFloat {
	fractions := map[rune]float32{}
	total := float32(0)
	for _, v := range f.Values {
		total += float32(v)
	}
	for k, v := range f.Values {
		fractions[k] = float32(v) / total
	}
	return fractions
}

func LoadFrequencyFloat(filePath string) (*FrequencyFloat, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Create a new Frequency instance
	var freq FrequencyFloat

	// Decode the JSON file into the Frequency struct
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&freq); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &freq, nil
}

func StoreFrequencyFloat(output *os.File, freq Frequency) error {
	eng := FrequencyFloat{
		Name:   "English",
		Values: freq.ToFractions(),
	}
	jsonData, err := json.MarshalIndent(eng, "", "  ")
	if err != nil {
		fmt.Printf("Error: Failed to encode frequency analysis to JSON: %v\n", err)
		return err
	}

	_, err = output.Write(jsonData)
	if err != nil {
		fmt.Printf("Error: Failed to write JSON output to file: %v\n", err)
		return err
	}
	return nil
}
