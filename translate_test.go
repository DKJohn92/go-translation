package main

import "testing"

func TestTranslate(t *testing.T) {
	tests := []struct {
		sourceText     string
		targetLanguage string
		screenName     string
		expected       string
	}{
		{"データセーフティ", "en", "", "data safety"},
		{"WOVN.io", "en", "MainActivity", "WOVN.io - MainScreen"},
		{"WOVN.io", "en", "", "WOVN.io"},
		{"Non-existing text", "en", "", "Non-existing text"},
	}

	for _, tt := range tests {
		result := translate(tt.sourceText, tt.targetLanguage, tt.screenName)
		if result != tt.expected {
			t.Errorf("translate(%q, %q, %q) = %q; want %q", tt.sourceText, tt.targetLanguage, tt.screenName, result, tt.expected)
		}
	}
}
