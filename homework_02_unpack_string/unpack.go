package homework02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcStr string) (string, error) {
	if srcStr == "" {
		return "", nil
	}

	runes := []rune(srcStr)

	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var (
		result       strings.Builder
		isEscapeFlag bool
	)

	for i := 0; i < len(runes); i++ {
		letter := runes[i]

		if letter == '\\' && !isEscapeFlag {
			if !unicode.IsDigit(runes[i+1]) && runes[i+1] != '\\' {
				return "", ErrInvalidString
			}

			isEscapeFlag = true

			continue
		}

		if unicode.IsDigit(letter) && !isEscapeFlag {
			count, _ := strconv.Atoi(string(letter))
			result.WriteString(strings.Repeat(string(runes[i-1]), count))

			continue
		}

		isEscapeFlag = false
		if len(runes) > i+1 && unicode.IsDigit(runes[i+1]) && !isEscapeFlag {
			if len(runes) > i+2 && unicode.IsDigit(runes[i+2]) {
				return "", ErrInvalidString
			}

			continue
		}

		result.WriteRune(letter)
	}

	return result.String(), nil
}
