package services

import (
	"backbootcamp/models"
	"backbootcamp/repositories"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------- Login tests -------------
// Success login
func Test_LoginWhenValidShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)

	params := struct {
		email    string
		password string
	}{
		email:    "alamriosx",
		password: "undecodedpassword",
	}

	user := &models.User{
		Id:        1,
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "undecodedpassword",
	}

	s := NewUserService(r)

	r.EXPECT().
		Login(params.email, params.password).
		Return(user, nil)

	got, err := s.Login(params.email, params.password)

	assert.Equal(t, got, user)
	assert.NoError(t, err)
}

// Fail login
func Test_LoginWhenInvalidShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)

	params := struct {
		email    string
		password string
	}{
		email:    "alamriosx",
		password: "undecodedpassword",
	}

	s := NewUserService(r)

	r.EXPECT().
		Login(params.email, params.password).
		Return(nil, fmt.Errorf("no user found"))

	got, err := s.Login(params.email, params.password)

	assert.Nil(t, got)
	assert.Error(t, err)
}

// ------------- CREATE tests -------------
// Create user with complete information
func Test_InsertUserWhenValidShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)

	user := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "undecodedpassword",
	}

	resultUser := &models.User{
		Id:        1,
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
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
func Test_InsertUserWhenIncompleteShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	t.Run("firstname missing", func(t *testing.T) {
		user := &models.User{
			LastName: "Ríos Altamirano",
			Email:    "alamriosx@gmail.com",
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
			Password:  "undecodedpassword",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "email missing")
	})

	t.Run("password missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Email:     "alamriosx@gmail.com",
		}
		got, err := s.InsertUser(user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "password missing")
	})
}

// Create user with no information
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

// Create user with incorrect format password
func Test_InsertUserWhenIncorrectPasswordFormatShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	user := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "123",
	}
	got, err := s.InsertUser(user)
	assert.Nil(t, got)
	assert.Error(t, err, "invalid password")
}

// TODO: Create user with password not decoded

// ------------- READ tests -------------
// Read user without id
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

// Read user with existing id
func Test_ReadUserWhenValidIdShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expected := &models.User{
		Id:        1,
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "undecodedpassword",
	}

	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)
	r.EXPECT().
		GetUserById(1).
		Return(expected, nil)

	got, err := s.GetUser(1)
	assert.Equal(t, got, expected)
	assert.NoError(t, err)
}

// Read user with no existing id
func Test_ReadUserWhenInvalidIdShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)
	r.EXPECT().
		GetUserById(1).
		Return(nil, fmt.Errorf("any employee was found with given id"))

	got, err := s.GetUser(1)
	assert.Nil(t, got)
	assert.EqualError(t, err, "any employee was found with given id")
}

// ------------- UPDATE tests -------------
// Edit user with complete information
func Test_UpdateUserWhenValidShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	user := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "undecodedpassword",
	}

	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	r.EXPECT().
		GetUserById(1).
		Return(user, nil)

	r.EXPECT().
		UpdateUser(1, user).
		Return(user, nil)

	got, err := s.UpdateUser(1, user)
	assert.Equal(t, got, user)
	assert.NoError(t, err)
}

// Edit user with incomplete information
func Test_UpdateUserWhenIncompleteShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	t.Run("firstname missing", func(t *testing.T) {
		user := &models.User{
			LastName: "Ríos Altamirano",
			Email:    "alamriosx@gmail.com",
			Password: "undecodedpassword",
		}
		got, err := s.UpdateUser(1, user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "fistname missing")
	})

	t.Run("lastname missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			Email:     "alamriosx@gmail.com",
			Password:  "undecodedpassword",
		}
		got, err := s.UpdateUser(1, user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "lastname missing")
	})

	t.Run("email missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Password:  "undecodedpassword",
		}
		got, err := s.UpdateUser(1, user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "email missing")
	})

	t.Run("password missing", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alam Yael",
			LastName:  "Ríos Altamirano",
			Email:     "alamriosx@gmail.com",
		}
		got, err := s.UpdateUser(1, user)
		assert.Nil(t, got)
		assert.EqualError(t, err, "password missing")
	})
}

// Edit user with no information
func Test_UpdateUserWhenNoInfoShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	user := &models.User{}
	got, err := s.UpdateUser(1, user)
	assert.Nil(t, got)
	assert.Error(t, err, "firstname missing")
}

// Edit user with incorrect format password
func Test_UpdateUserWhenIncorrectPasswordFormatShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)

	user := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "123",
	}
	got, err := s.UpdateUser(1, user)
	assert.Nil(t, got)
	assert.Error(t, err, "invalid password")
}

// TODO: Edit user with not decoded password

// ------------- DELETE tests -------------
// Delete user with existing id
func Test_DeleteUserWhenValidIdShouldSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expected := &models.User{
		Id:        1,
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Password:  "undecodedpassword",
	}

	r := repositories.NewMockIUserRepository(ctrl)
	s := NewUserService(r)
	r.EXPECT().
		GetUserById(1).
		Return(expected, nil)

	r.EXPECT().
		DeleteUser(1).
		Return(nil)

	err := s.DeleteUser(1)
	assert.NoError(t, err)
}

// Delete user with no existing id
func Test_DeleteUserWhenInvalidIdShouldError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := repositories.NewMockIUserRepository(ctrl)

	s := NewUserService(r)
	r.EXPECT().
		GetUserById(1).
		Return(nil, fmt.Errorf("any employee was found with given id"))

	err := s.DeleteUser(1)
	assert.EqualError(t, err, "any employee was found with given id")
}
