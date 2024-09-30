package sCrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

/*
 * Decrypt-able algorithms
 */

func EncryptAES256(stringToEncrypt string, keyString string) (string, error) {
	// MARK: Since the key is in string, we need to convert decode it to bytes.
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	// MARK: Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("aes.NewCipher(key): %v", err)
	}

	// MARK: Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	// MARK: https://golang.org.pkg.crypto.cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("cipher.NewGCM(block): %v", err)
	}

	// MARK: Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("io.ReadFull(rand.Reader, nonce): %v", err)
	}

	// MARK: Encrypt the data using aesGCM.Seal
	// MARK: Since we don't want to save the nonce somewhere else in this case, we add it as prefix to encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return hex.EncodeToString(ciphertext), nil
}

func DecryptAES256(stringToDecrypt string, keyString string) (string, error) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(stringToDecrypt)

	// MARK: Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("aes.NewCipher(key): %v", err)
	}

	// MARK: Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("cipher.NewGCM(block): %v", err)
	}

	// MARK: Get the nonce size
	nonceSize := aesGCM.NonceSize()

	// MARK: Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// MARK: Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("aesGCM.Open(nil, nonce, ciphertext, nil): %v", err)
	}

	return string(plaintext), nil
}

/*
 * UnDecrypt-able algorithms. HashAlgorithms.
 */

func GetSHA1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func GetSHA256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func GetSHA512(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
