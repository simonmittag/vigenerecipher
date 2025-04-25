# 🚧 UNDER CONSTRUCTION 🚧

# Vigenère Cipher

Polyalphabetic substitution is an improvement over the simpler Caesar cipher. Each letter is shifted based on the
relative position of letters in a corresponding keyword.

## Getting Started

### Prerequisites

This project requires **Go** to be installed. To set up Go, you can use the following command on macOS:

```bash
brew install go
```

### Installation

To install the Caesar cipher command-line tool, use:

```bash
go install github.com/simonmittag/vigenerecipher/cmd/vigenere
```

This will make the `vigenere` command available globally.

## Usage

The `vigenere` command-line tool supports encrypting and decrypting text files using an input files and a keyword. Below are the available options and examples to get you started:

### Options

```bash
🇫🇷vigenere 0.0.1
Usage: vigenere [options]
Options:
  -d    Decrypt the input file and store results in output file
  -e    Encrypt the input file and store results in output file
  -h    Print usage information
  -i string
        Path to the input text file (required)
  -o string
        Path to the output text file (required)
  -p string
        A password for both encryption or decrytion
```

### Examples

#### Encrypt a File

To encrypt the content of `mary.txt` using a keyword of `hannibal` and save the result to `mary_encrypted.txt`:

```bash
vigenere -e -p hannibal -i mary.txt -o mary_encrypted.txt
```
