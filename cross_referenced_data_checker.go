package main

import (
	"fmt"
	"os"
)

// CrossReferencedDataChecker is a JSON schema format validator that enforces valid links to other data points
type CrossReferencedDataChecker struct{}

// IsFormat will return false if the referenced data point cannot be read
func (f CrossReferencedDataChecker) IsFormat(input interface{}) bool {
	reference, ok := input.(string)
	if ok == false {
		return false
	}

	path := "data/" + reference + ".yaml"
	if _, err := os.Stat(path); err != nil {
		fmt.Printf("⛔️ link error to %s: %s\n", path, err)
		return false
	}

	return true
}
