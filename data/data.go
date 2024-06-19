package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
    "name" text unique
    );`
	createTableStationSQL := `create table if not exists station(
    "id" integer not null primary key autoincrement,
    "name" text unique
    );`
	createTableUserSQL := `create table if not exists user(
        "id" integer not null primary key autoincrement,
        "name" text,
        "login" text unique, 
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
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	fmt.Println("table role created")

	statement, err = db.Prepare(createTableStationSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	fmt.Println("table station created")

	statement, err = db.Prepare(createTableUserSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	fmt.Println("table user created")

	statement, err = db.Prepare(createTableStationVisitSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	fmt.Println("table stationVisit created")
}

// Roles

type Role struct {
	Id   int
	Name string
}

func InsertRole(role_name string) {
	insertRoleSQL := `
        insert into role (name) 
        values ($1)
    `
	_, err := db.Exec(insertRoleSQL, role_name)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Роль %s успешно добавлена \n", role_name)
}

func GetNamesRoles() []string {
	selectRoleSQL := `
        select name
        from role 
        order by id desc
    `

	rows, err := db.Query(selectRoleSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	roles := []string{}

	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			fmt.Println(err)
			continue
		}
		roles = append(roles, r)
	}
	return roles
}

func GetRoles() []Role {
	selectRoleSQL := `
		select *
		from role 
		order by id
	`
	rows, err := db.Query(selectRoleSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	roles := []Role{}

	for rows.Next() {
		r := Role{}
		err := rows.Scan(&r.Id, &r.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		roles = append(roles, r)
	}

	return roles
}

func getRoleIdByName(role string) int {
	getRoleIdSQL := `
	select id 
	from role
	where name = $1 
	limit 1
`

	row := db.QueryRow(getRoleIdSQL, role)
	var role_id int
	err := row.Scan(&role_id)

	if err != nil {
		log.Fatal(err.Error())
	}

	return role_id
}

// Stations

type Station struct {
	Id   int
	Name string
}

func InsertStations(station_name string) {
	insertStationSQL := `
        insert into station (name) 
        values ($1)
    `
	_, err := db.Exec(insertStationSQL, station_name)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Станция %s успешно добавлена\n", station_name)
}

func GetStations() []Station {
	getStationsSQL := `
		select * 
		from station
		order by id
	`
	rows, err := db.Query(getStationsSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	stations := []Station{}

	for rows.Next() {
		station := Station{}
		err := rows.Scan(&station.Id, &station.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		stations = append(stations, station)
	}
	return stations
}

func GetStationsName() []string {
	selectStationsNameSQL := `
		select name 
		from station
		order by id
	`
	rows, err := db.Query(selectStationsNameSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	stations := []string{}

	for rows.Next() {
		var station string

		err := rows.Scan(&station)

		if err != nil {
			fmt.Println(err)
		}

		stations = append(stations, station)
	}
	return stations
}

func GetStationIdByName(station_name string) int {
	getStationSQL := `
		select id 
		from station 
		where name = $1
		limit 1
	`
	row := db.QueryRow(getStationSQL, station_name)
	var station_id int

	err := row.Scan(&station_id)

	if err != nil {
		log.Fatal(err.Error())
	}

	return station_id

}

// Users

type User struct {
	Id        int
	Name      string
	Login     string
	Password  string
	Role_id   int
	Role_name string
}

func InsertUser(name, login, password, role string) {
	insertUserSQL := `
        insert into user (name, login, password, role_id)
        values($1, $2, $3, $4)
    `

	role_id := getRoleIdByName(role)

	_, err := db.Exec(insertUserSQL, name, login, password, role_id)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Пользователь %s успешно добавлен\n", name)
}

func GetUsers() []User {
	getUsersSQL := `
		select user.id, user.name, user.login, user.password, user.role_id, role.name
		from user
		join role on user.role_id = role.id
	`

	rows, err := db.Query(getUsersSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Login, &u.Password, &u.Role_id, &u.Role_name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users
}

// StationVisit

type StationVisit struct {
	Id           int
	Datetime     string
	Station_id   int
	Station_name string
}

func InsertStationVisit(station_id int) {
	insertStationVisitSQL := `
        insert into stationVisit (station_id, datetime)
        values ($1 , $2)
    `

	_, err := db.Exec(insertStationVisitSQL, station_id, time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetStationVisits() []StationVisit {
	selectStationsVisitSQL := `
		select stationVisit.id, stationVisit.datetime, stationVisit.station_id , station.name
		from stationVIsit
		join station on stationVisit.station_id = station.id
		order by stationVisit.id desc
	`
	rows, err := db.Query(selectStationsVisitSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	stationVisits := []StationVisit{}

	for rows.Next() {
		stationVisit := StationVisit{}

		err := rows.Scan(&stationVisit.Id, &stationVisit.Datetime, &stationVisit.Station_id, &stationVisit.Station_name)

		if err != nil {
			fmt.Println(err)
		}

		stationVisits = append(stationVisits, stationVisit)
	}

	return stationVisits
}
