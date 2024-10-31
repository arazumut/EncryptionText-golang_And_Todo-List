package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)



func generateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func encrypt(plainText, key string) (string, error) {
	plainTextBytes := []byte(plainText)
	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv, err := generateRandomBytes(aes.BlockSize)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	padding := aes.BlockSize - len(plainTextBytes)%aes.BlockSize
	paddedText := append(plainTextBytes, make([]byte, padding)...)

	cipherText := make([]byte, len(paddedText))
	mode.CryptBlocks(cipherText, paddedText)

	return hex.EncodeToString(append(iv, cipherText...)), nil
}

func decrypt(cipherText, key string) (string, error) {
	cipherTextBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	plainTextBytes := make([]byte, len(cipherTextBytes))
	mode.CryptBlocks(plainTextBytes, cipherTextBytes)

	padding := plainTextBytes[len(plainTextBytes)-1]
	plainTextBytes = plainTextBytes[:len(plainTextBytes)-int(padding)]

	return string(plainTextBytes), nil
}

//! Produced by Kamil Umut Araz instagram : K.umutarazz  github : arazumut

func main() {
	key := "examplekey1234567890123456"
	plainText := "Hello, World!"

	encryptedText, err := encrypt(plainText, key)
	if err != nil {
		fmt.Println("Encryption failed:", err)
		return
	}
	fmt.Println("Encrypted Text:", encryptedText)

	decryptedText, err := decrypt(encryptedText, key)
	if err != nil {
		fmt.Println("Decryption failed:", err)
		return
	}
	fmt.Println("Decrypted Text:", decryptedText)
}
