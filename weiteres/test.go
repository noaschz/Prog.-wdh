package test

import (
	//"fmt"
	"testing"
)

// --- Aufgabe 1 ---
// ReverseWords kehrt die Reihenfolge der Wörter in einem Satz um.
func ReverseWords(sentence string) string {
	// TODO
	return ""
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Ich liebe Go", "Go liebe Ich"},
		{"Hallo Welt", "Welt Hallo"},
	}

	for _, test := range tests {
		result := ReverseWords(test.input)
		if result != test.expected {
			t.Errorf("ReverseWords(%q) = %q, erwartet %q", test.input, result, test.expected)
		}
	}
}