package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Kind string

const (
	KindAsterisk = Kind("ASTERISK")
	KindDigit    = Kind("DIGIT")
	KindCommon   = Kind("COMMON")
)

var ErrInvalidString = errors.New("invalid string")

func KindOfRune(r rune) Kind {
	if unicode.IsDigit(r) {
		return KindDigit
	}
	if r == '\\' {
		return KindAsterisk
	}
	return KindCommon
}

func Unpack(input string) (string, error) {
	// Обработка исключений
	if input == "" {
		return input, nil
	}

	runedInput := []rune(input)
	switch KindOfRune(runedInput[0]) {
	case KindDigit:
		return "", ErrInvalidString
	case KindAsterisk, KindCommon:
	}

	if len(runedInput) == 1 {
		if KindOfRune(runedInput[0]) == KindAsterisk {
			return "", ErrInvalidString
		}
		return input, nil
	}

	if KindOfRune(runedInput[len(runedInput)-1]) == KindAsterisk {
		return "", ErrInvalidString
	}

	var result strings.Builder
	var candidate rune
	var count int
	var err error
	i := 0

	for i < len(runedInput) {
		current := runedInput[i]

		switch KindOfRune(current) {
		case KindCommon:
			candidate = current
		case KindAsterisk:
			next := runedInput[i+1]

			if KindOfRune(next) == KindCommon {
				return "", ErrInvalidString
			}
			candidate = next
			i++
		case KindDigit:
			return "", ErrInvalidString
		}

		i++
		if i > len(runedInput)-1 {
			result.WriteRune(runedInput[len(runedInput)-1])
			break
		}

		current = runedInput[i]

		if KindOfRune(current) == KindDigit {
			count, err = strconv.Atoi(string(current))
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(string(candidate), count))
			i++
		} else {
			result.WriteRune(candidate)
		}
	}

	return result.String(), nil
}
