package repositories

import "database/sql"

type UserRepository struct {
	dbClient *sql.DB
}

func NewUserRepository(dbClient *sql.DB) *UserRepository {
	return &UserRepository{
		dbClient: dbClient,
	}
}
