package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Encrypt takes in a key and plaintext and returns encrypted version
func Encrypt(key, plaintext string) (string, error) {
	block, err := newCipherBlock(key)

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

// This code is based on the standard library examples
func Decrypt(key, encryptedHex string) (string, error) {
	block, err := newCipherBlock(key)

	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	hasher.Write([]byte(key))

	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
