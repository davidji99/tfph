package tfph

import (
	"fmt"
	"strings"
)

func ParseCompositeID(id string, numOfSplits int, separator ...string) ([]string, error) {
	sep := ":"
	if len(separator) > 0 {
		sep = separator[0]
	}

	parts := strings.SplitN(id, sep, numOfSplits)

	if len(parts) != numOfSplits {
		return nil, fmt.Errorf("error: composite ID requires %d parts separated by a [%[2]s] (x%[2]sy)",
			numOfSplits, sep)
	}

	return parts, nil
}

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
