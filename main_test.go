package main

import "testing"

func TestValidateInput(t *testing.T) {
	t.Run("test_empty_password", func(t *testing.T) {
		err := validateInput("", "")
		if err == nil {
			t.Errorf("Expected error but received none")
		}
	})
	t.Run("test_invalid_alg", func(t *testing.T) {
		err := validateInput("", "SOMERANDOMALG")
		if err == nil {
			t.Errorf("Expected error but received none")
		}
	})
	t.Run("test_valid_input", func(t *testing.T) {
		err := validateInput("securestring", "SHA3")
		if err != nil {
			t.Errorf("Expected no error but received %s", err.Error())
		}
	})
}
