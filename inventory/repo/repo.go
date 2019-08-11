package repo

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/miun173/inventory-dispusibda/inventory/models"
	"github.com/pkg/errors"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// CheckUsernameExist repo
func CheckUsernameExist(user models.User) (bool, error) {
	var exists bool
	q := fmt.Sprintf("SELECT COUNT(1) FROM users WHERE username = '%s'", user.Username)
	err := db.QueryRow(q).Scan(&exists)
	if err != nil {
		err = errors.WithStack(err)
		return false, err
	}

	return exists, nil
}

// CreateUser insert new user to db
func CreateUser(user *models.User) error {
	stm, err := db.Prepare("INSERT INTO users (username, firstname, lastname, password, role, token) VALUES (?, ?, ?, ?, ?, ?)")
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

// GetUserByID get user by id
func GetUserByID(user *models.User, userID int) error {
	q := "SELECT id, firstname, lastname, password, role, token FROM users WHERE id = ? "
	err := db.QueryRow(q, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Token)
	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return err
	}

	return nil
}

// GetUserByUsername find user by given username
func GetUserByUsername(user *models.User, username string) error {
	q := "SELECT id, firstname, lastname, password, role, token FROM users WHERE username= ? "
	err := db.QueryRow(q, username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Token)

	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return err
	}

	return nil
}

// GetAllUser query all users
func GetAllUser() ([]models.User, error) {
	users := make([]models.User, 0)
	q := "SELECT id, username, firstname, lastname, role, token FROM users"
	rows, err := db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "select all user error")
		log.Printf("%+v\n", err)
		return users, err
	}
	defer rows.Close()

	var u models.User
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Role, &u.Token); err != nil {
			err = errors.Wrap(err, "scan users row error")
			log.Printf("%+v\n", err)
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

// CheckBarangExists check if barang exists by its ID
func CheckBarangExists(barangID int) (bool, error) {
	var exists bool
	q := fmt.Sprintf("SELECT COUNT(1) FROM barang WHERE id = %d", barangID)
	err := db.QueryRow(q).Scan(&exists)
	if err != nil {
		err = errors.Wrap(err, "check barang exists error")
		log.Printf("%+v\n", err)
		return false, err
	}

	log.Printf("exists, %v", exists)

	return exists, nil
}
