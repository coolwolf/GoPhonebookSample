package db

func CreateTables() {
	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		in_use INTEGER DEFAULT 1,
		inserted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		inserted_by INTEGER,
		updated_at DATETIME,
		updated_by INTEGER
	);`
	_, err := DB.Exec(createUsers)
	if err != nil {
		panic(err)
	}

	createContacts := `
	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone TEXT NOT NULL,
		in_use INTEGER DEFAULT 1,
		inserted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		inserted_by INTEGER,
		updated_at DATETIME,
		updated_by INTEGER
	);`
	_, err = DB.Exec(createContacts)
	if err != nil {
		panic(err)
	}
}
