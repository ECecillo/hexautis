package parser

import "testing"

func TestHexToString(t *testing.T) {
	tests := []struct {
		hexStr string
		want   string
	}{
		{"48656c6c6f20576f726c64", "Hello World"},
		{"4c6f72656d20697073756d", "Lorem ipsum"},
		{"4c6f72656d20697073756d0a", "Lorem ipsum\n"},
	}
	for _, tt := range tests {
		t.Run(tt.hexStr, func(t *testing.T) {
			got, err := HexToString(tt.hexStr)
			if err != nil {
				t.Errorf("HexToString() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("HexToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToStringWithNoneHexValue(t *testing.T) {
	tests := []struct {
		hexStr string
	}{
		{"Hello"},
		{"World"},
		{"Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.hexStr, func(t *testing.T) {
			_, err := HexToString(tt.hexStr)
			if err == nil {
				t.Errorf("HexToString() error = %v, want error", err)
				return
			}
		})
	}
}

func TestHexToStringWithInvalidHexLength(t *testing.T) {
	tests := []struct {
		hexStr string
	}{
		{"656"},
		{"48656c6c6f20576f726ca"},
		{"4c6f72656d20697073756d0"},
	}
	for _, tt := range tests {
		t.Run(tt.hexStr, func(t *testing.T) {
			_, err := HexToString(tt.hexStr)
			if err == nil {
				t.Errorf("HexToString() error = %v, want error", err)
				return
			}
		})
	}
}
