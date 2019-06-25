package repository

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/pkg/errors"
)

// UserRepo represent repository of the user
// Expect implementation of infrastructure layer
type UserRepo interface {
	InitUser() error
	Get(id int) (models.User, error)
	GetAll() ([]models.User, error)
	Save(*models.User) error
	GetUserByUsername(username string) (models.User, error)
}

type userRepo struct {
	db *sql.DB
}

// NewUserRepo get user repo
func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) InitUser() error {
	users := [3][2]string{{"divisi", "divisi"}, {"subagumum", "pejabat"}, {"petugasBarang", "petugasBarang"}}
	for _, user := range users {
		user := models.User{
			Username:  user[0],
			FirstName: user[0],
			LastName:  user[0],
			Role:      user[1],
			Password:  "123",
		}

		exists, err := u.CheckUsernameExist(user)
		if err != nil {
			log.Fatalf("%+v\n", err)
			return err
		}

		if !exists {
			err = u.Save(&user)
			if err != nil {
				log.Fatalf("%+v\n", err)
				return err
			}
		}
	}

	return nil
}

// CheckUsernameExist repo
func (u *userRepo) CheckUsernameExist(user models.User) (bool, error) {
	var exists bool
	q := fmt.Sprintf("SELECT COUNT(1) FROM users WHERE username = '%s'", user.Username)
	err := u.db.QueryRow(q).Scan(&exists)
	if err != nil {
		err = errors.WithStack(err)
		return false, err
	}

	return exists, nil
}

// Save user to persistance storage
func (u *userRepo) Save(user *models.User) error {
	stm, err := u.db.Prepare("INSERT INTO users (username, firstname, lastname, password, role, token) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		err = errors.Wrap(err, "create user error")
		log.Printf("%+v\n", err)
		return err
	}

	secret := time.Now().String()
	token := base64.StdEncoding.EncodeToString([]byte(string(secret)))

	res, err := stm.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Role, token)
	if err != nil {
		err = errors.Wrap(err, "create user error")
		log.Printf("%+v\n", err)
		return err
	}
	defer stm.Close()

	id, err := res.LastInsertId()
	user.ID = int(id)
	return nil
}

// Get get user by id
func (u *userRepo) Get(userID int) (models.User, error) {
	var user models.User
	q := "SELECT id, firstname, lastname, password, role, token FROM users WHERE id = ? "
	err := u.db.QueryRow(q, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Token)
	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return user, err
	}

	return user, nil
}

// GetUserByUsername find user by given username
func (u *userRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	q := "SELECT id, firstname, lastname, password, role, token FROM users WHERE username= ? "
	err := u.db.QueryRow(q, username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Token)

	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return user, err
	}

	return user, nil
}

// GetAllUser query all users
func (u *userRepo) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	q := "SELECT id, username, firstname, lastname, role, token FROM users"
	rows, err := u.db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "select all user error")
		log.Printf("%+v\n", err)
		return users, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Role, &user.Token); err != nil {
			err = errors.Wrap(err, "scan users row error")
			log.Printf("%+v\n", err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
