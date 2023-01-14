package main

import "golang.org/x/crypto/bcrypt"

// Hashes password ---✨
func HashString(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Check if password input matches hashed password ---✨
func CheckStringHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
