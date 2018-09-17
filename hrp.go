package hrp

import (
	"fmt"
	"regexp"
	"strings"
)

// Convert ..
func Convert(s string) (string, string, error) {
	if len(strings.TrimSpace(s)) == 0 {
		return "", "", fmt.Errorf("got the void")
	}

	vType := s[len(s)-1:]
	value := s[:len(s)-1]

	r := regexp.MustCompile("[a-zA-Z]")
	if !r.MatchString(vType) {
		vType = ""
		value = s[:len(s)]
	}

	if value == "" {
		return "", "", fmt.Errorf("incorrect format")
	}

	return value, vType, nil
}
