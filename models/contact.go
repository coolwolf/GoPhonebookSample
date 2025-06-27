package models

import (
	"database/sql"
	"log"
	"phonebook/db"
	"time"
)

type Contact struct {
	ID         int
	Name       string
	Phone      string
	InUse      int
	InsertedAt time.Time
	InsertedBy int
	UpdatedAt  sql.NullTime
	UpdatedBy  sql.NullInt64
}

// CreateContact inserts a new contact into the DB
func CreateContact(name, phone string, insertedBy int) error {
	_, err := db.DB.Exec(`
		INSERT INTO contacts (name, phone, inserted_by) 
		VALUES (?, ?, ?)`,
		name, phone, insertedBy)
	if err != nil {
		log.Println("Error inserting contact:", err)
	}
	return err
}

func UpdateContact(id int, name, phone string, updatedBy int) error {
	_, err := db.DB.Exec(`
		update contacts set name= ?, phone = ?, updated_by = ?, updated_at = ? where id = ?`,
		name, phone, updatedBy, time.Now(), id)
	if err != nil {
		log.Println("Error updating contact:", err)
	}
	return err
}

func ListContacts(query string) ([]Contact, error) {
	appendWhare := ""
	if query != "" {
		appendWhare = " and name like '%" + query + "%' OR phone like '%" + query + "%'"
	}
	queryString := "SELECT id, name, phone, in_use, inserted_at, inserted_by, updated_at, updated_by FROM contacts WHERE 1=1 " + appendWhare + " order by name"
	//log.Println("Contact list:", queryString)
	rows, err := db.DB.Query(queryString)
	if err != nil {
		log.Println("Error listing contacts:", err)
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var c Contact
		err := rows.Scan(&c.ID, &c.Name, &c.Phone, &c.InUse, &c.InsertedAt, &c.InsertedBy, &c.UpdatedAt, &c.UpdatedBy)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func GetContact(id int) (Contact, error) {
	row, err := db.DB.Query(`SELECT id, name, phone, in_use, inserted_at, inserted_by, updated_at, updated_by FROM contacts where id = ?`, id)
	if err != nil {
		log.Println("Error inserting contact:", err)
	}
	defer row.Close()

	var contact Contact
	if row.Next() {
		err := row.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.InUse, &contact.InsertedAt, &contact.InsertedBy, &contact.UpdatedAt, &contact.UpdatedBy)
		if err != nil {
			log.Println("Error scanning contact:", err)
			return contact, err
		}
	}

	return contact, nil
}

func DeleteContact(id int, updatedBy int) error {
	// Soft delete by setting in_use = 0
	_, err := db.DB.Exec(`UPDATE contacts SET in_use = 0, updated_at = CURRENT_TIMESTAMP, updated_by = ? WHERE id = ?`,
		updatedBy, id)
	return err
}
