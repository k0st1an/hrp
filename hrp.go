package hrp

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	v1 = `(^\d*\w$)`      // 3d
	v2 = `(^\d*\w\s\w*$)` // 12k wps / 42 rps
)

var (
	errIncorrectFormat = fmt.Errorf("incorrect format")
)

// Convert ..
func Convert(s string) (string, string, error) {
	switch true {
	case regexp.MustCompile(v1).MatchString(s):
		return convertV1(s)
	case regexp.MustCompile(v2).MatchString(s):
		return convertV2(s)
	}
	return "", "", errIncorrectFormat
}

func convertV1(s string) (string, string, error) {
	vType := s[len(s)-1:]
	value := s[:len(s)-1]

	r := regexp.MustCompile("[a-zA-Z]")
	if !r.MatchString(vType) {
		vType = ""
		value = s[:len(s)]
	}

	if value == "" {
		return "", "", errIncorrectFormat
	}

	return value, vType, nil
}

func convertV2(s string) (string, string, error) {
	splitS := strings.Split(s, " ")
	vType := splitS[1]
	value, m, _ := convertV1(splitS[0])

	switch m {
	case "k":
		value = value + "000"
	}
	return value, vType, nil
}
