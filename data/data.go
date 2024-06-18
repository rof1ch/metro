package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
    
	return db.Ping()
}

func CreateTable() {
	createTableRoleSQL := `create table if not exists role(
    "id" integer not null primary key autoincrement,
    "name" text
    );`
	createTableStationSQL := `create table if not exists station(
    "id" integer not null primary key autoincrement,
    "name" text
    );`
	createTableUserSQL := `create table if not exists user(
        "id" integer not null primary key autoincrement,
        "name" text,
        "login" text, 
        "password" text, 
        "role_id" integer,
        foreign key(role_id) references role(id)
    );`
	createTableStationVisitSQL := `create table if not exists stationVisit(
        "id" integer not null primary key autoincrement,
        "datetime" text,
        "station_id" integer,
        foreign key(station_id) references station(id)
        );`

    statement, err := db.Prepare(createTableRoleSQL)
    if err != nil{
        log.Fatal(err.Error())
    }
    statement.Exec()
    fmt.Println("table role created")

    statement, err = db.Prepare(createTableStationSQL)
    if err != nil{
        log.Fatal(err.Error())
    }
    statement.Exec()
    fmt.Println("table station created")

    statement, err = db.Prepare(createTableUserSQL)
    if err != nil{
        log.Fatal(err.Error())
    }
    statement.Exec()
    fmt.Println("table user created")

    statement, err = db.Prepare(createTableStationVisitSQL)
    if err != nil{
        log.Fatal(err.Error())
    }
    statement.Exec()
    fmt.Println("table stationVisit created")
}

// func InsertRole() {
//     insertRoleSQL
// }