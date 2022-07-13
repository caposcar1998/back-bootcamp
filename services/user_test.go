package services

import (
	"backbootcamp/models"
	"backbootcamp/repositories"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestUserService_InsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repositories.NewMockIUserRepository(ctrl)

	userWCI := &models.User{
		FirstName: "Alam Yael",
		LastName:  "Ríos Altamirano",
		Email:     "alamriosx@gmail.com",
		Username:  "alamriosx",
		Password:  "undecodedpassword",
	}

	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr string
	}{
		{name: "user with complete info should success", args: args{user: userWCI}, want: userWCI, wantErr: ""},
		// user with incomplete info should error
		// user with no info should error
		// user with incorrect password format should error
		// user password without cypher should error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: r,
			}
			got, err := s.InsertUser(tt.args.user)

			r.EXPECT().

			if tt.wantErr != "" && tt.wantErr != err.Error() {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
