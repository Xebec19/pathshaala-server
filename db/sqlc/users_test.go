package db

import (
	"testing"

	"github.com/Xebec19/pathshaala-server/utils"
)

func TestCreateAccount(t *testing.T) {

}

func createRandomTeacherAccount(t *testing.T) User {
	arg := CreateUsersParams{
		FirstName: utils.RandomName(),
		LastName:  utils.RandomName(),
		Email:     utils.RandomEmail(),
		Password:  utils.RandomString(16),
		Role:      "teacher",
	}
}
