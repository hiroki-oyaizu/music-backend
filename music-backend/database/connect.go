package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	user := "webuser"
	password := "webpass"
	host := "localhost"
	port := "3308"
	database_name := "music_development"

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func CreateTableIfNotExists(db *sql.DB) error {
	// テーブルが存在しない場合に作成するクエリ
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tests (
		id int AUTO_INCREMENT,
		name varchar(100),
		age int,
		title varchar(255),
		PRIMARY KEY(id)
	);
	`
	// クエリを実行
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}
