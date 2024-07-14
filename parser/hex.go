package parser

import (
	"errors"
	"strconv"
)

func HexToString(hexStr string) (string, error) {
	var runes []rune
	for i := 0; i < len(hexStr); i += 2 {
		if i+1 >= len(hexStr) {
			return "", errors.New("invalid hex length")
		}
		pair := hexStr[i : i+2]
		num, err := strconv.ParseUint(pair, 16, 8)
		if err != nil {
			return "", err
		}
		runes = append(runes, rune(num))
	}
	return string(runes), nil
}
