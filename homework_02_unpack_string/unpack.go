package homework02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var (
		unpack   string
		prevRune rune
	)

	for _, r := range runes {
		if unicode.IsDigit(r) && unicode.IsDigit(prevRune) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) {
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}

			if count > 0 {
				unpack += strings.Repeat(string(prevRune), count-1)
			} else {
				unpack = strings.Replace(unpack, string(prevRune), "", 1)
			}

			prevRune = r

			continue
		}

		unpack += string(r)
		prevRune = r
	}

	return unpack, nil
}
