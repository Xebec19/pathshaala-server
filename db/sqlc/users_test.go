package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Xebec19/pathshaala-server/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomTeacherAccount(t)
}

type sqlNullString struct {
	String string
	Valid  bool
}

func createRandomTeacherAccount(t *testing.T) User {
	lastName := sqlNullString{utils.RandomName(), true}
	arg := CreateUsersParams{
		FirstName: utils.RandomName(),
		LastName:  sql.NullString(lastName),
		Email:     utils.RandomEmail(),
		Password:  utils.RandomString(16),
		Role:      NullUserRole{"teacher", true},
	}
	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	return user
}
