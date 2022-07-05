package encode

import (
	"encoding/base64"
	"math/rand"
	"strings"
)

func Base64Encode(keyword string) string {
	input := []byte(keyword)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func Base64Decode(encodeString string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	return string(decodeBytes), err
}

func First2Upper(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func FixLine(line string) string {
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\t", "")
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.ReplaceAll(line, "\xc2\xa0", "")
	line = strings.ReplaceAll(line, " ", "")
	return line
}

func FixMap(m map[string]string) map[string]string {
	var arr []string
	rm := make(map[string]string)
	for key, value := range m {
		if value == "" {
			continue
		}
		if IsInStrArr(arr, value) {
			continue
		}
		arr = append(arr, value)
		rm[key] = value
	}
	return rm
}

func IsInStrArr(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func IsInIntArr(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func StrRandomCut(s string, length int) string {
	sRune := []rune(s)
	if len(sRune) > length {
		i := rand.Intn(len(sRune) - length)
		return string(sRune[i : i+length])
	} else {
		return s
	}
}

func AutoWidth(s string, length int) int {
	length1 := len(s)
	length2 := len([]rune(s))

	if length1 == length2 {
		return length
	}
	return length - (length1-length2)/2
}
