package randstr

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"unsafe"
)

const (
	// strLen is the length of the length of the random generated strings.
	strLen = 16
	// chars contains all the default characters used to generate random string with Generate().
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// Generate returns a random string that is 16 characters long.
func Generate() (string, error) {
	return GenerateLen(strLen)
}

// GenerateLen returns a random string of the specified length.
func GenerateLen(length int) (string, error) {
	if length < 0 {
		return "", errors.New("invalid length: length must be greater than or equal to 0")
	}
	if length == 0 {
		return "", nil
	}
	n := len(chars)
	str := make([]byte, length)
	b := make([]byte, length+(length/2))
	maxb := 255 - (256 % n)
	i := 0
	for {
		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		for _, byt := range b {
			bint := int(byt)
			if bint > maxb { // prevent modulo bias
				continue
			}
			str[i] = chars[bint%n]
			i += 1
			if i == length {
				return *(*string)(unsafe.Pointer(&str)), nil
			}
		}
	}
}

// Base64Encoded returns a random base64 encoded string.
func Base64Encoded() (string, error) {
	byt, err := generateBytes()
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(byt)
	return encoded, nil
}

func generateBytes() ([]byte, error) {
	buffer := make([]byte, strLen)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
