package db

import (
	"Microservice-Login/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) UserDetail {
	arg := CreateUserParams{
		FirstName: util.RandomFirstNameGenerator(),
		LastName:  util.RandomLastNameGenerator(),
		UserName:  util.RandomUsernameGenerator(),
		Email:     util.RandomEmailGenerator(),
		Mobile:    util.RandomMobileGenerator(),
	}

	userDetails, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userDetails)

	require.Equal(t, arg.FirstName, userDetails.FirstName)
	require.Equal(t, arg.LastName, userDetails.LastName)
	require.Equal(t, arg.UserName, userDetails.UserName)
	require.Equal(t, arg.Email, userDetails.Email)
	require.Equal(t, arg.Mobile, userDetails.Mobile)

	return userDetails
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
