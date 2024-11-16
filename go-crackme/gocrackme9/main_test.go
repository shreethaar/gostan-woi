
package main

import "testing"

func TestValidateSerial(t *testing.T) {
    tests := []struct {
        serial  string
        isValid bool
    }{
        {"A1C3F7J1K4", true},  // Example valid serial
        {"a1c3f7j1k4", true},  // Lowercase version (converted to uppercase in function)
        {"A1C3F7J1K5", false}, // Invalid serial (last character mismatch)
        {"SHORT", false},      // Too short
        {"TOOLONGSERIAL123", false}, // Too long
        {"1234567890", false}, // Completely invalid characters
    }

    for _, test := range tests {
        got := validateSerial(test.serial)
        if got != test.isValid {
            t.Errorf("validateSerial(%q) = %v; want %v", test.serial, got, test.isValid)
        }
    }
}
