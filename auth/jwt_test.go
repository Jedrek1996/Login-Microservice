package auth

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateJWT1(t *testing.T) {

	jwtMaker, err := NewJWTMaker("token")
	if err != nil {
		log.Panic(err)
	}

	userID := "user6"
	token, err := jwtMaker.CreateJWTToken(userID, time.Duration(time.Second)*300)

	if err != nil {
		log.Panic(err)
	}
	// Auth token is printed here to be used for local testing of APIs.
	fmt.Println("=====================================")
	fmt.Printf("Authorization Token for testing: %s\n", token)
	fmt.Println("=====================================")
}

func TestVerifyJWT1(t *testing.T) {

	jwtVerifier, err := NewJWTVerifier("token")
	if err != nil {
		log.Panic(err)
	}
	isValid, err := jwtVerifier.IsValidToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InVzZXI2IiwiaXNzdWVkX2F0IjoiMjAyMy0wMy0yM1QxNDo1ODoyMS41NzI0OTQrMDg6MDAiLCJleHBpcmVkX2F0IjoiMjAyMy0wMy0yM1QxNTowMzoyMS41NzI0OTQrMDg6MDAifQ.cC27UhG5bpaF3r7A195SV59n0hspbzfDuL_zBNrGad-jRn1dOUxu0ntPs0QdsGl1My9m73187YUGDcn9FCWGQg")

	if err != nil {
		assert.Error(t, err)
	}
	// Auth token is printed here to be used for local testing of APIs.

	assert.Equal(t, true, isValid)

}
