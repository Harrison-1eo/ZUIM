package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// StreamCipher struct to hold the key and position
type StreamCipher struct {
	key      []uint32
	position int
}

// NewStreamCipher constructor to initialize the StreamCipher with a hexadecimal key
func NewStreamCipher(hexKey string) (*StreamCipher, error) {
	if len(hexKey) != 64 {
		return nil, errors.New("the key must be a 64-character hexadecimal string")
	}

	key := make([]uint32, 0, 32)
	for i := 0; i < 64; i += 2 {
		parsedKey, err := hex.DecodeString(hexKey[i : i+2])
		if err != nil {
			return nil, err
		}
		key = append(key, uint32(parsedKey[0]))
	}

	return &StreamCipher{key: key, position: 0}, nil
}

// Encrypt method to encrypt/decrypt the plainText
func (sc *StreamCipher) Encrypt(plainText string) (string, error) {
	plainTextBytes := []byte(plainText)
	cipherTextBytes := make([]byte, len(plainTextBytes))
	unitSize := 1

	for i := 0; i < len(plainTextBytes); i += unitSize {
		keyUnit := sc.key[sc.position]
		plainTextUnit := uint32(plainTextBytes[i])

		cipherTextUnit := plainTextUnit ^ keyUnit

		cipherTextBytes[i] = byte(cipherTextUnit)

		sc.position = (sc.position + 1) % len(sc.key)
	}
	println("Hex format: ", hex.EncodeToString(cipherTextBytes))
	println("String format: ", string(cipherTextBytes))
	println("Base64 encode format: ", base64.StdEncoding.EncodeToString(cipherTextBytes))
	return base64.StdEncoding.EncodeToString(cipherTextBytes), nil
}

func (sc *StreamCipher) Decrypt(encryptedText string) (string, error) {
	encryptedTextBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	plainTextBytes := make([]byte, len(encryptedTextBytes))
	unitSize := 1

	for i := 0; i < len(encryptedTextBytes); i += unitSize {
		keyUnit := sc.key[sc.position]
		encryptedTextUnit := uint32(encryptedTextBytes[i])

		plainTextUnit := encryptedTextUnit ^ keyUnit

		plainTextBytes[i] = byte(plainTextUnit)

		sc.position = (sc.position + 1) % len(sc.key)
	}

	return string(plainTextBytes), nil
}

func main() {
	hexKey := "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	cipher, err := NewStreamCipher(hexKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	plainText := "Hello, world!"
	encryptedText, err := cipher.Encrypt(plainText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cipher2, err := NewStreamCipher(hexKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	decryptedText, err := cipher2.Decrypt(encryptedText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Encrypted Text:", encryptedText)
	fmt.Println("Decrypted Text:", decryptedText)
}
