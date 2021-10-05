package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Database Connection Pool
	Database *sql.DB
)

// NewDatabase - Create a new database connection
func NewDatabase() error {
	db, err := sql.Open("sqlite3", "./database.db?parseTime=true")
	if err != nil {
		return err
	}
	Database = db

	// Return no error if there isn't a failure
	return nil
}

type User struct {
	ID         int
	Name       string
	Department string
	Created    time.Time
}

func CreateUsersTable() error {
	stmt := `
	create table userinfo (uid integer not null primary key, username varchar, departname varchar, created date);
	delete from userinfo;
	`
	pstmt, err := Database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = pstmt.Exec()
	if err != nil {
		return err
	}

	// Return no error if there isn't a failure
	return nil
}

func DropUsersTable() error {
	stmt := `
	drop table userinfo;
	`
	pstmt, err := Database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = pstmt.Exec()
	if err != nil {
		return err
	}
	// Return no error if there isn't a failure
	return nil
}

func AddNewUser(username, department string) error {
	stmt, err := Database.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, department, time.Now())
	if err != nil {
		return err
	}

	// Return no error if there isn't a failure
	return nil
}

func GetUser(ID int) (*User, error) {
	rows, err := Database.Query("SELECT * FROM userinfo WHERE uid = ?", ID)
	if err != nil {
		return nil, err
	}
	user := &User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Department, &user.Created)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
