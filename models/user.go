package models

import (
	"database/sql"
	"log"
	"phonebook/db"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	PasswordHash string
	InUse        int
	InsertedAt   time.Time
	InsertedBy   int
	UpdatedAt    sql.NullTime
	UpdatedBy    sql.NullInt64
}

// CreateUser inserts a new user into the DB
func CreateUser(username, password string, insertedBy int) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	passwordHash := string(hashedBytes)

	_, err = db.DB.Exec(`INSERT INTO users (username, password_hash, inserted_by) VALUES (?, ?, ?)`, username, passwordHash, insertedBy)
	if err != nil {
		log.Println("Error inserting user:", err)
	}
	return err
}

func ListUsers() ([]User, error) {
	rows, err := db.DB.Query(`SELECT id, username, password_hash, in_use, inserted_at, inserted_by, updated_at, updated_by FROM users order by username`)
	if err != nil {
		log.Println("Error listing users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.InUse, &u.InsertedAt, &u.InsertedBy, &u.UpdatedAt, &u.UpdatedBy)
		if err != nil {
			log.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUser(id int) (*User, error) {
	var u User
	err := db.DB.QueryRow(`SELECT id, username, password_hash, in_use, inserted_at, inserted_by, updated_at, updated_by FROM users WHERE id = ? AND in_use = 1`, id).
		Scan(&u.ID, &u.Username, &u.PasswordHash, &u.InUse, &u.InsertedAt, &u.InsertedBy, &u.UpdatedAt, &u.UpdatedBy)
	return &u, err
}

func UpdateUser(id int, username, password string, updatedBy int) error {
	if password != "" {
		// hash new password
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		passwordHash := string(hashedBytes)
		_, err = db.DB.Exec(`UPDATE users SET username=?, password_hash=?, updated_at=CURRENT_TIMESTAMP, updated_by=? WHERE id=?`,
			username, passwordHash, updatedBy, id)
		return err
	}

	_, err := db.DB.Exec(`UPDATE users SET username=?,updated_at=CURRENT_TIMESTAMP, updated_by=? WHERE id=?`,
		username, updatedBy, id)
	return err
}

func DeleteUser(id int, updatedBy int) error {
	_, err := db.DB.Exec(`UPDATE users SET in_use=0, updated_at=CURRENT_TIMESTAMP, updated_by=? WHERE id=?`, updatedBy, id)
	return err
}

func GetUserByUsername(username string) (User, error) {
	var u User
	err := db.DB.QueryRow(`SELECT id, username, password_hash, in_use, inserted_at, inserted_by, updated_at, updated_by
		FROM users WHERE username=? AND in_use=1 LIMIT 1`, username).
		Scan(&u.ID, &u.Username, &u.PasswordHash, &u.InUse, &u.InsertedAt, &u.InsertedBy, &u.UpdatedAt, &u.UpdatedBy)
	return u, err
}
