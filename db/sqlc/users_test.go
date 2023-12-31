package sqlc

import (
	"context"
	"database/sql"
	"simple_shop/db/util"
	"simple_shop/db/util/random"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(random.RandomHashedPassword())
	require.NoError(t, err)

	username_tmp := random.RandomUsername()
	email := username_tmp + "@gmail.com"

	arg := CreateUserParams{
		Username:       username_tmp,
		FullName:       random.RandomFullName(),
		HashedPassword: hashedPassword,
		Email:          email,
	}

	username, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, username)

	require.Equal(t, arg.Username, username.Username)
	require.Equal(t, arg.FullName, username.FullName)
	require.Equal(t, arg.HashedPassword, username.HashedPassword)
	require.Equal(t, arg.Email, email)

	return username
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestDeleteUser(t *testing.T) {
	username_1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), username_1.Username)
	require.NoError(t, err)

	username_2, err := testQueries.GetUser(context.Background(), username_1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, username_2)
}

func TestGetUser(t *testing.T) {
	username_1 := createRandomUser(t)
	username_2, err := testQueries.GetUser(context.Background(), username_1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, username_2)

	require.Equal(t, username_1.Username, username_2.Username)
	require.Equal(t, username_1.FullName, username_2.FullName)
	require.Equal(t, username_1.HashedPassword, username_2.HashedPassword)
	require.Equal(t, username_1.Email, username_2.Email)
	require.Equal(t, username_1.Level, username_2.Level)

	require.WithinDuration(t, username_1.CreatedAt, username_2.CreatedAt, time.Second)
	require.WithinDuration(t, username_1.PasswordChangedAt, username_2.PasswordChangedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  10,
		Offset: 5,
	}

	listUsernames, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, listUsernames)

	for _, username := range listUsernames {
		require.NotEmpty(t, username)
	}
}
func TestUpdateHashedPasswordOfUser(t *testing.T) {
	username_1 := createRandomUser(t)

	arg := UpdateHashedPasswordOfUserParams{
		Username:          username_1.Username,
		HashedPassword:    "sercet",
		PasswordChangedAt: time.Now(),
	}

	username_2, err := testQueries.UpdateHashedPasswordOfUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, username_2)

	require.Equal(t, arg.Username, username_2.Username)
	require.Equal(t, username_1.FullName, username_2.FullName)
	require.Equal(t, arg.HashedPassword, username_2.HashedPassword)
	require.Equal(t, username_1.Email, username_2.Email)
	require.Equal(t, username_1.Level, username_2.Level)

	require.WithinDuration(t, username_1.CreatedAt, username_2.CreatedAt, time.Second)
	require.WithinDuration(t, arg.PasswordChangedAt, username_2.PasswordChangedAt, time.Second)
}
