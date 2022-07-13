package services

import (
	"backbootcamp/models"
	"backbootcamp/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Create user with complete information
func Test_InsertUserWhenValidUserShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)

	user := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Username:  "alamriosx",
		Password:  "undecodedpassword",
	}

	resultUser := &models.User{
		Id:        1,
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Username:  "alamriosx",
		Password:  "undecodedpassword",
	}

	s := NewUserService(r)

	r.EXPECT().
		AddUser(user).
		Return(resultUser, nil)

	got, err := s.InsertUser(user)

	assert.Equal(t, got, resultUser)
	assert.NoError(t, err)
}

// Create user with incomplete information
func Test_InsertUserWhenIncompleteUserShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	t.Run("firstname missing", func(t *testing.T) {
		user := &models.User{
			LastName: "Ríos Altamirano",
			Email:    "alamriosx@gmail.com",
			Username: "alamriosx",
			Password: "undecodedpassword",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "fistname missing")
	})

	t.Run("lastname missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			Email:     "alamriosx@gmail.com",
			Username:  "alamriosx",
			Password:  "undecodedpassword",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "lastname missing")
	})

	t.Run("email missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Username:  "alamriosx",
			Password:  "undecodedpassword",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "email missing")
	})

	t.Run("username missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Email:     "alamriosx@gmail.com",
			Password:  "undecodedpassword",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "username missing")
	})

	t.Run("password missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Email:     "alamriosx@gmail.com",
			Username:  "alamriosx",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "password missing")
	})
}

// user with no info should error
func Test_InsertUserWhenNoInfoShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	user := &models.User{}
	got, err := s.InsertUser(user)
	assert.Nil(t, got)
	assert.Error(t, err, "firstname missing")
}

// user with incorrect password format should error
func Test_InsertUserWhenIncorrectPasswordFormatShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	user := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Username:  "alamriosx",
		Password:  "123",
	}
	got, err := s.InsertUser(user)
	assert.Nil(t, got)
	assert.Error(t, err, "invalid password")
}

// TODO: user password without cypher should error

func Test_GetAllUsersShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	expected := []models.User{
		{
			Id:        1,
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Email:     "alamriosx@gmail.com",
			Username:  "alamriosx",
			Password:  "undecodedpassword",
		},
	}

	s := NewUserService(r)
	r.EXPECT().
		GetAllUsers().
		Return(expected, nil)

	got, err := s.GetAllUsers()
	assert.Equal(t, got, expected)
	assert.NoError(t, err)
}

// TODO: single read
// read with existing id should success
// read with no-existing id should error

// TODO: update user
// user with complete info should success
// user with incomplete info should error
// user with no-info should error
// user with incorrect password format should error
// user password without cypher should error

// TODO: delete user
// delete with existing id should success
// delete with no-existing id should error
