package filters

import (
	"fmt"
	"strings"
)

type Filter func(string) error

func BadWordFilter(s string) error {
	if strings.Contains(s, "beer") {
		return fmt.Errorf("drugs are not allowed")
	}
	return nil
}
