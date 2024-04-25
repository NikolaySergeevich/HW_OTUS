package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	strRune := []rune(str)
	if len(strRune) == 0 {
		return "", nil
	}
	if _, er := strconv.Atoi(string(strRune[0])); er == nil {
		return "", ErrInvalidString
	}
	var res strings.Builder

	for i, v := range strRune {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			if !checkNextNul(i, strRune) {
				res.WriteRune(v)
			}
			continue
		}
		if !checkNextRune(i, strRune) {
			return "", ErrInvalidString
		}
		if num == 0 {
			continue
		}
		for j := 0; j < num-1; j++ {
			res.WriteRune(strRune[i-1])
		}
	}
	return res.String(), nil
}

func checkNextRune(i int, r []rune) bool {
	if i == len(r)-1 {
		return true
	}
	if _, err := strconv.Atoi(string(r[i+1])); err == nil {
		return false
	}
	return true
}

func checkNextNul(i int, r []rune) bool {
	if i == len(r)-1 {
		return false
	}
	if num, err := strconv.Atoi(string(r[i+1])); err == nil && num == 0 {
		return true
	}
	return false
}
