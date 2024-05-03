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

	for i := 0; i < len(strRune); i ++ {
		num, err := strconv.Atoi(string(strRune[i]))
		if err != nil {
			if strRune[i] == '\\' && checkLastRune(i, strRune) || strRune[i] == '\\' && !checkLastRune(i, strRune) && strRune[i+1] != '\\' && !checkNextRuneInt(i, strRune){
				return "", ErrInvalidString
			}
			if strRune[i] == '\\' && (strRune[i+1] == '\\' || checkNextRuneInt(i, strRune)){
				i += 1
				res.WriteRune(strRune[i])
				continue
			}
			if !checkNextNul(i, strRune) {
				res.WriteRune(strRune[i])
			}
			continue
		}
		if checkNextRuneInt(i, strRune) {
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

func checkNextRuneInt(i int, r []rune) bool {
	if checkLastRune(i, r) {
		return false
	}
	if _, err := strconv.Atoi(string(r[i+1])); err == nil {
		return true
	}
	return false
}

func checkLastRune(i int, r []rune) bool {
	return i == len(r)-1
}

func checkNextNul(i int, r []rune) bool {
	if checkLastRune(i, r) {
		return false
	}
	if num, err := strconv.Atoi(string(r[i+1])); err == nil && num == 0 {
		return true
	}
	return false
}
