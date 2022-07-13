package services

import (
	"backbootcamp/models"
	"backbootcamp/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestUserService_GetAllUsers(t *testing.T) {
	type fields struct {
		repository IUserRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.User
		wantErr bool
	}{
		// success case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
			}
			got, err := s.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	type fields struct {
		repository IUserRepository
	}
	type args struct {
		idUser int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// read with existing id should success
		// read with no-existing id should error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
			}
			got, err := s.GetUser(tt.args.idUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	type fields struct {
		repository IUserRepository
	}
	type args struct {
		idUser int
		user   *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// user with complete info should success
		// user with incomplete info should error
		// user with no-info should error
		// user with incorrect password format should error
		// user password without cypher should error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
			}
			got, err := s.UpdateUser(tt.args.idUser, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	type fields struct {
		repository IUserRepository
	}
	type args struct {
		idUser int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// delete with existing id should success
		// delete with no-existing id should error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
			}
			if err := s.DeleteUser(tt.args.idUser); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
