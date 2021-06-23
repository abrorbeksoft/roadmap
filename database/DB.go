package database

import "database/sql"

func Connect() *sql.DB {
	conectStr:="user=root password=password dbname=roadmap sslmode=disable"
	db,err:=sql.Open("postgres",conectStr)
	if err!=nil {
		panic(err)
	}
	return db
}
