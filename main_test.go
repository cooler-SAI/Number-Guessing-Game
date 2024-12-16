package main

import "testing"

func TestDetermineChances(t *testing.T) {
	tests := []struct {
		name      string
		choice    string
		want      int
		expectErr bool
	}{
		{"Easy level", "1", 10, false},
		{"Medium level", "2", 5, false},
		{"Hard level", "3", 3, false},
		{"Invalid choice", "4", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := determineChances(tt.choice)
			if (err != nil) != tt.expectErr {
				t.Errorf("determineChances() error = %v, wantErr %v", err, tt.expectErr)
				return
			}
			if got != tt.want {
				t.Errorf("determineChances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckGuess(t *testing.T) {
	tests := []struct {
		name     string
		guess    int
		number   int
		expected string
	}{
		{"Correct guess", 50, 50, "correct"},
		{"Guess is less", 30, 50, "greater"},
		{"Guess is greater", 70, 50, "less"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkGuess(tt.guess, tt.number)
			if result != tt.expected {
				t.Errorf("checkGuess() = %v, want %v", result, tt.expected)
			}
		})
	}
}
