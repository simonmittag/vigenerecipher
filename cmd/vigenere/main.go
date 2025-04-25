package main

import (
	"flag"
	"fmt"
	"github.com/simonmittag/vigenerecipher"
	"os"
)

const Version = "0.1.0"

func main() {

	var mode string
	inputFile := flag.String("i", "", "Path to the input text file (required)")
	outputFile := flag.String("o", "", "Path to the output text file (required)")
	password := flag.String("p", "", "Password to use for encryption/decryption (required)")
	encrypt := flag.Bool("e", false, "Encrypt the input file and store results in output file")
	decrypt := flag.Bool("d", false, "Decrypt the input file and store results in output file")
	help := flag.Bool("h", false, "Print usage information")

	flag.Parse()

	if *help || (flag.NFlag() == 0) {
		printUsage()
		return
	}

	if !*encrypt && !*decrypt {
		fmt.Println("Error: You must specify either -e (encrypt) or -d (decrypt)")
		printUsage()
		os.Exit(1)
	}
	if *inputFile == "" {
		fmt.Println("Error: Input file (-i flag) is required.")
		printUsage()
		os.Exit(1)
	}
	if *outputFile == "" {
		fmt.Println("Error: Output file (-o flag) is required.")
		printUsage()
		os.Exit(1)
	}

	if *encrypt {
		mode = "encrypt"
	} else if *decrypt {
		mode = "decrypt"
	}

	input, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error: Could not open input file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error: Could not create output file: %v\n", err)
		os.Exit(1)
	}
	defer output.Close()

	cipher := vigenerecipher.NewVigenereCipher(*password)

	switch mode {
	case "encrypt":
		if err := cipher.Shift(input, output, false); err != nil {
			fmt.Printf("Error during encryption: %v\n", err)
			os.Exit(1)
		}
	case "decrypt":
		if err := cipher.Shift(input, output, true); err != nil {
			fmt.Printf("Error during decryption: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Output written to %s\n", *outputFile)
}

func printUsage() {
	fmt.Println("🇫🇷vigenere " + Version)
	fmt.Println("Usage: vigenere [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
