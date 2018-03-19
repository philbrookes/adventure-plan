package mysql

import (
	"database/sql"

	//used for mysql... herpaderp
	_ "github.com/go-sql-driver/mysql"
	"github.com/philbrookes/adventure-plan/pkg/config"
)

//Factory returns a connection to a MySQL database
type Factory func() (*sql.DB, error)

//GetDB returns a function which creates a connection to a MySQL server specified in the config
func GetDB(config *config.MySQLConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.GetConnectionString())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

//InitDB creates the database and tables required by this application
func InitDB(db *sql.DB, config *config.MySQLConfig) error {
	stmt, err := db.Prepare("CREATE DATABASE IF NOT EXISTS " + config.GetDBName())
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	_, err = db.Exec("USE " + config.GetDBName())
	if err != nil {
		return err
	}

	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS maps (
			id INT NOT NULL AUTO_INCREMENT,
			owner INT,
			center_lat DECIMAL(10,8),
			center_lng DECIMAL(10,8),
			zoom INT NULL,
			title VARCHAR(45) NULL,

			PRIMARY KEY (id),
			INDEX (id),
			INDEX (owner)
		);
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS points (
			id INT NOT NULL AUTO_INCREMENT,
			map_id INT NOT NULL,
			latitude DECIMAL(10,8), 
			longitude DECIMAL(10,8), 
			title varchar(45),
			notes varchar(128),

			PRIMARY KEY (id),
			INDEX (id),
			INDEX (map_id),
			FOREIGN KEY (map_id) REFERENCES maps(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
