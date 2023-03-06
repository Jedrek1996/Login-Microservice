package util

import (
	"regexp"
	"strings"
)

// ValidateFirstName validates the first name input
func ValidateFirstName(name string) bool {
	// Check if name has minimum 3 letters
	if len(name) < 3 {
		return false
	}
	// Check if name has only letters
	match, _ := regexp.MatchString("^[a-zA-Z]*$", name)
	return match
}

// ValidateLastName validates the last name input
func ValidateLastName(name string) bool {
	// Check if name has minimum 3 letters
	if len(name) < 3 {
		return false
	}
	// Check if name has only letters
	match, _ := regexp.MatchString("^[a-zA-Z]*$", name)
	return match
}

// ValidateUsername validates the username input
func ValidateUsername(username string) bool {
	// Check if username has minimum 6 letters
	if len(username) < 6 {
		return false
	}
	// Check if username has only letters and digits
	match, _ := regexp.MatchString("^[a-zA-Z0-9]*$", username)
	return match
}

// ValidatePassword validates the password input
func ValidatePassword(password string) bool {
	// Check if password has minimum 8 characters
	if len(password) < 8 {
		return false
	}
	return true
}

// ValidateEmail validates the email input
func ValidateEmail(email string) bool {
	// Check if email is in a valid format
	match, _ := regexp.MatchString("^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$", email)
	return match
}

// ValidateMobile validates the mobile input for Singapore
func ValidateMobile(mobile string) bool {
	// Check if mobile is in a valid format for Singapore number
	match, _ := regexp.MatchString("^(\\+65|65)?[689]\\d{7}$", strings.Replace(mobile, " ", "", -1))
	return match
}
