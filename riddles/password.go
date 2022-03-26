package riddles

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	ErrPasswordLength                 = fmt.Errorf("invalid password length: length cannot be zero")
	ErrPasswordLengthLessThanSum      = fmt.Errorf("invalid password length: length < minDigits + minSpecialChars")
	ErrPasswordInvalidMinDigits       = fmt.Errorf("invalid minimal digits: cannot be lees than zero")
	ErrPasswordInvalidMinSpecialChars = fmt.Errorf("invalid minimal special chars: cannot be lees than zero")
	ErrPasswordFailRandomGenerator    = fmt.Errorf("fail to generate random number")
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digits  = "0123456789"
	special = "+-@*#"
)

// PasswordGenerator
// Allowed regular characters: A to Z or a to z
// Allowed special characters: '+', '-', '@', '*', '#'
// Allowed digits: 0-9
// Length has to be >= (number_special_chars + number_digits)
// Number of special characters has to be >= 0
// Number of digits has to be >= 0
func PasswordGenerator(length, minDigits, minSpecialChars int) (string, error) {
	if length < minDigits+minSpecialChars {
		return "", ErrPasswordLengthLessThanSum
	}

	if length == 0 {
		return "", ErrPasswordLength
	}

	if minDigits < 0 {
		return "", ErrPasswordInvalidMinDigits
	}

	if minSpecialChars < 0 {
		return "", ErrPasswordInvalidMinSpecialChars
	}

	digitsRandom, err := getRandomString(digits, minDigits)
	if err != nil {
		return "", err
	}

	specialRandom, err := getRandomString(special, minSpecialChars)
	if err != nil {
		return "", err
	}

	anyCharPassword, err := getRandomString(letters+special+digits, length-minSpecialChars-minDigits)
	if err != nil {
		return "", err
	}

	password, err := scrambler(digitsRandom + specialRandom + anyCharPassword)
	if err != nil {
		return "", err
	}
	return password, nil
}

func randomIndex(max int) (int64, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, ErrPasswordFailRandomGenerator
	}

	return index.Int64(), nil
}

func getRandomString(chars string, count int) (string, error) {
	length := len(chars)
	var gen string
	for i := 0; i < count; i++ {
		index, err := randomIndex(length)
		if err != nil {
			return "", err
		}
		gen += string(chars[index])
	}
	return gen, nil
}

func scrambler(pass string) (string, error) {
	length := len(pass)
	passRune := []rune(pass)
	for i := range passRune {
		index, err := randomIndex(length)
		if err != nil {
			return "", err
		}
		passRune[i], passRune[index] = passRune[index], passRune[i]

	}
	return string(passRune), nil
}
