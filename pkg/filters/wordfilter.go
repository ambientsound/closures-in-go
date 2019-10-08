package filters

import (
	"fmt"
	"strings"
)

type Filter func(string) error

func BadWordFilter(badWords []string) Filter {
	return func(s string) error {
		for _, word := range badWords {
			if strings.Contains(s, word) {
				return fmt.Errorf("filtered: %s", word)
			}
		}
		return nil
	}
}
