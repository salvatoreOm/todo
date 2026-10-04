package handler

import (
	"fmt"
	"strconv"
)

// parseID converts a command-line argument into a todo ID.
func parseID(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid id %q: must be a positive number", arg)
	}
	return id, nil
}
