package util

import (
	"crypto/rand"
	"errors"
	"unicode"
)

const (
	numberBytes    = "0123456789"
	letterBytes    = "ABCDEFGHJKLMNPQRSTUVWXYZ0123456789"
	letterBytesAll = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits  = 6                    // 6 bits to represent 64 possibilities / indexes
	letterIdxMask  = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

var (
	ErrInvalidLength            = errors.New("invalid length")
	ErrFailedToMakeRandomString = errors.New("failed to make random string")
)

func GetRandomStringClean(length int) (string, error) {
	var err error
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes, err = SecureRandomBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result), nil
}

func GetRandomStringAll(length int) (string, error) {
	var err error
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes, err = SecureRandomBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytesAll) {
			result[i] = letterBytesAll[idx]
			i++
		}
	}

	return string(result), nil
}

func GetRandomNumbers(length int) (string, error) {
	var err error
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes, err = SecureRandomBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(numberBytes) {
			result[i] = numberBytes[idx]
			i++
		}
	}

	return string(result), nil
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) ([]byte, error) {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, ErrFailedToMakeRandomString
	}
	return randomBytes, nil
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
