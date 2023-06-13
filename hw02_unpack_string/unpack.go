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
		if err != nil {
			if character == '\\' {
				if isPrevSlash {
					prevRune = character
					isPrevRune = true
					isPrevSlash = false
				} else {
					isPrevSlash = true
					runes = append(runes, prevRune)
					rep = append(rep, 1)
				}
			} else if isPrevInt {
				prevRune = character
				isPrevRune = true
				isPrevInt = false
				fmt.Println("isPrevInt " + string(prevRune))
			} else if isPrevRune {
				runes = append(runes, prevRune)
				rep = append(rep, 1)
				isPrevRune = true
				fmt.Println("isPrevRune" + string(prevRune))
				prevRune = character
			}
		} else if isPrevSlash && err == nil {
			isPrevSlash = false
			isPrevRune = true
			prevRune = character
			fmt.Println("isPrevSlash " + string(prevRune) + " " + strconv.Itoa(repeats))
		} else {
			runes = append(runes, prevRune)
			rep = append(rep, repeats)
			println("Default " + string(prevRune) + string(character))
			isPrevInt = true

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
		if err != nil {
			prevIsInt = false
			if character != '\\' && prevIsSlash {
				return false
			}
			if character == '\\' {
				prevIsSlash = true
			} else if character == '\\' && prevIsSlash {
				prevIsSlash = false
			}
		} else if prevIsInt && prevIsSlash {
			prevIsSlash = false
		} else if prevIsInt {
			return false
			break
		} else {
			prevIsInt = true
		}
	}
	return true
}
