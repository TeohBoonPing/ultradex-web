package utils

import (
	"fmt"
	"strconv"
)

// StringToPtr converts a string to a pointer to that string.
func StringToPtr(s string) *string {
	return &s
}

// StringToFloat64Ptr converts a string to a pointer to a float64.
func StringToFloat64Ptr(s string) (*float64, error) {
	if s == "" {
		return nil, nil
	}

	parsed, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert string to float64: %w", err)
	}

	return &parsed, nil
}

// StringToUintPtr converts a string to a pointer to a uint.
func StringToUintPtr(s string) (*uint, error) {
	if s == "" {
		return nil, nil
	}

	parsed, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert string to uint: %w", err)
	}

	uintVal := uint(parsed)
	return &uintVal, nil
}
