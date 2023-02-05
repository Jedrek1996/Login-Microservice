package util

import (
	"regexp"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

// Random input generators for user testing ---✨
func RandomFirstNameGenerator() (firstName string) {
	firstName = gofakeit.FirstName()
	return
}

func RandomLastNameGenerator() (lastName string) {
	lastName = gofakeit.LastName()
	return
}

func RandomUsernameGenerator() (username string) {
	username = gofakeit.Username()
	return
}

func RandomEmailGenerator() (email string) {
	email = gofakeit.Email()
	return
}

func RandomMobileGenerator() int32 {
	mobile := gofakeit.Phone()
	mobile = removePart(mobile)
	return convertToInt32(mobile)
}

// Specifically for mobile as the generated number is more than 8 digits ---✨
func removePart(str string) string {
	re := regexp.MustCompile(`\d{8}`)
	indexes := re.FindStringIndex(str)

	if indexes != nil {
		return str[:indexes[1]]
	}

	if len(str) < 8 {
		return str
	}

	return str[:8]
}

// Converts string to int64 to int 32 ---✨
func convertToInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int32(i)
}
