package randstr

import (
	"crypto/rand"
	"encoding/base64"
)

var (
	// strLen is the length of the length of the random generated strings.
	strLen = 16
	// chars contains all the default characters used to generate random string with Generate().
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// Generate returns a random string that is strLen characters long.
func Generate() (string, error) {
	b := make([]byte, 1)
	maxb := 255 - (256 % len(chars))
	n := len(chars)
	str := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		bint := int(b[0])
		if bint > maxb { // prevent modulo bias
			i -= 1
			continue
		}
		str[i] = chars[bint%n]
	}
	return string(str), nil
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
