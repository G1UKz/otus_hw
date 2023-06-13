package hw02unpackstring

import (
	"errors"
	"fmt"
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
	var prevRune rune
	var isPrevRune, isPrevSlash bool
	isPrevInt := true
	var runes []rune
	var rep []int
	for i, character := range input {
		repeats, err := strconv.Atoi(string(character))
		switch {
		case err != nil:
			switch {
			case character == '\\' && isPrevSlash:
				prevRune = character
				isPrevRune = true
				isPrevSlash = false
			case character == '\\' && !isPrevSlash:
				isPrevSlash = true
				runes = append(runes, prevRune)
				rep = append(rep, 1)
			case isPrevInt:
				prevRune = character
				isPrevRune = true
				isPrevInt = false
				fmt.Println("isPrevInt " + string(prevRune))
			case isPrevRune:
				runes = append(runes, prevRune)
				rep = append(rep, 1)
				isPrevRune = true
				fmt.Println("isPrevRune" + string(prevRune))
				prevRune = character
			}
		case err == nil:
			switch {
			case isPrevSlash:
				isPrevSlash = false
				isPrevRune = true
				prevRune = character
				fmt.Println("isPrevSlash " + string(prevRune) + " " + strconv.Itoa(repeats))
			default:
				runes = append(runes, prevRune)
				rep = append(rep, repeats)
				println("Default " + string(prevRune) + string(character))
				isPrevInt = true
			}
		}
		if len(input)-1 == i && !isPrevInt {
			runes = append(runes, prevRune)
			rep = append(rep, 1)
		}
	}
	fmt.Println(len(runes))
	println(len(rep))
	for i := 0; i < len(runes); i++ {
		builder.WriteString(strings.Repeat(string(runes[i]), rep[i]))
		fmt.Printf("Building string, i: %s,key: %s , value: %s \n", strconv.Itoa(i), string(runes[i]), strconv.Itoa(rep[i]))
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
			case prevIsInt && prevIsSlash:
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
