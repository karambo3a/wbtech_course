package unpacking

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	runes := []rune(s)
	repeatNumber := 0
	var sb strings.Builder
	var lastRune rune
	var err error
	for i, r := range runes {
		switch {
		case unicode.IsDigit(r):
			if i-1 < 0 {
				return "", errors.New("string must starts with a letter or escaped number")
			}

			if runes[i-1] == '\\' {
				if lastRune != 0 {
					repeatNumber = repeat(&sb, repeatNumber, lastRune)
				}
				lastRune = r
			} else {
				var digit int
				if repeatNumber == 0 {
					repeatNumber, err = strconv.Atoi(string(r))
				} else {
					digit, err = strconv.Atoi(string(r))
					repeatNumber = repeatNumber*10 + digit
				}
				if err != nil {
					return "", fmt.Errorf("atoi error: %v", err)
				}
			}
		case r != '\\':
			if i != 0 {
				repeatNumber = repeat(&sb, repeatNumber, lastRune)
			}
			if err != nil {
				return "", err
			}
			lastRune = r

		default:
			if i == len(runes)-1 {
				return "", errors.New("trailing backslash")
			}
		}
	}
	repeat(&sb, repeatNumber, lastRune)

	return sb.String(), err
}

func repeat(sb *strings.Builder, repeatNumber int, r rune) int {
	if repeatNumber == 0 {
		sb.WriteRune(r)
	}

	for range repeatNumber {
		sb.WriteRune(r)
	}
	return 0
}
