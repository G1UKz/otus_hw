package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if !isValidString(input) {
		return "", ErrInvalidString
	}
	output := formString(input)
	return output, nil
}

func formString(input string) (output string) {
	var builder strings.Builder
	var runes []rune
	var rep []int
	var isPrevSlash bool
	for _, character := range input {
		repeats, err := strconv.Atoi(string(character))
		switch {
		case err != nil:
			switch {
			case character == '\\' && !isPrevSlash:
				isPrevSlash = true
			default:
				runes = append(runes, character)
				rep = append(rep, 1)
				isPrevSlash = false
			}
		case err == nil:
			switch {
			case isPrevSlash:
				runes = append(runes, character)
				rep = append(rep, 1)
				isPrevSlash = false
			default:
				rep[len(rep)-1] = repeats
			}
		}
	}
	for z := 0; z < len(runes); z++ {
		builder.WriteString(strings.Repeat(string(runes[z]), rep[z]))
	}
	return builder.String()
}

func isValidString(input string) (valid bool) {
	prevIsInt := true
	prevIsSlash := false
	for _, character := range input {
		_, err := strconv.Atoi(string(character))
		switch {
		case err != nil:
			prevIsInt = false
			switch {
			case character != '\\' && prevIsSlash:
				return false
			case character == '\\' && prevIsSlash:
				prevIsSlash = false
			case character == '\\':
				prevIsSlash = true
			}
		case err == nil:
			switch {
			case prevIsSlash:
				prevIsSlash = false
			case prevIsInt:
				return false
			default:
				prevIsInt = true
			}
		}
	}
	return true
}
