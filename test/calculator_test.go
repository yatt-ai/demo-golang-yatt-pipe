package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yatt-ai/calculator/calculator"
)

func TestAdd(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 5, 3, 8},
		{"Negative numbers", -2, -3, -5},
		{"Mixed numbers", -5, 10, 5},
		{"Zeros", 0, 0, 0},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Addition result should match expected value")
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 10, 5, 5},
		{"Negative numbers", -5, -3, -2},
		{"Mixed numbers", 5, 10, -5},
		{"Zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.Subtract(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Subtraction result should match expected value")
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 5, 3, 15},
		{"Negative numbers", -2, -3, 6},
		{"Mixed numbers", -5, 10, -50},
		{"Zero multiplication", 5, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.Multiply(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Multiplication result should match expected value")
		})
	}
}

func TestDivide(t *testing.T) {
	// Test cases for successful division
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
		hasError bool
	}{
		{"Positive numbers", 10, 2, 5, false},
		{"Negative numbers", -6, -3, 2, false},
		{"Mixed numbers", -10, 5, -2, false},
		{"Division by zero", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculator.Divide(tt.a, tt.b)
			
			if tt.hasError {
				assert.Error(t, err, "Expected an error for division by zero")
			} else {
				assert.NoError(t, err, "Did not expect an error for valid division")
				assert.Equal(t, tt.expected, result, "Division result should match expected value")
			}
		})
	}
} 