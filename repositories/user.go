package repositories

import (
	"backbootcamp/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Login(email, password string) (*models.User, error) {
	var user *models.User

	rows, err := r.db.Query("SELECT * FROM user u WHERE u.email=? AND u.password=?", email, password)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		aux := new(models.User)
		err := rows.Scan(&aux.Id, &aux.FirstName, &aux.LastName,
			&aux.Email, &aux.Password)
		if err != nil {
			return nil, err
		}
		user = aux
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) AddUser(u *models.User) (*models.User, error) {
	stmtIns, err := r.db.Prepare("INSERT INTO user(firstname, lastname, email, password) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res, err := stmtIns.Exec(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	idUser, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.Id = int(idUser)
	return u, nil
}

func (r *UserRepository) GetUserById(idUser int) (*models.User, error) {
	var user *models.User

	rows, err := r.db.Query("SELECT * FROM user u WHERE u.id=?", idUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		aux := new(models.User)
		err := rows.Scan(&aux.Id, &aux.FirstName, &aux.LastName,
			&aux.Email, &aux.Password)
		if err != nil {
			return nil, err
		}
		user = aux
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := r.db.Query("SELECT * FROM user u")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		aux := new(models.User)
		err := rows.Scan(&aux.Id, &aux.FirstName, &aux.LastName,
			&aux.Email, &aux.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, *aux)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(idUser int, u *models.User) (*models.User, error) {
	stmtIns, err := r.db.Prepare("UPDATE user SET firstname=?, lastname=?, email=?, password=? WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(u.FirstName, u.LastName, u.Email, u.Password, idUser)
	if err != nil {
		return nil, err
	}
	u.Id = idUser
	return u, nil
}

func (r *UserRepository) DeleteUser(idUser int) error {
	stmtIns, err := r.db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(idUser)
	if err != nil {
		return err
	}

	return nil
}
