package repos

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	source = "../../../test.db"
)

var tables = []string{"follows", "history", "users", "segments"}

var db *sql.DB

// TODO fabric db
func TestMain(m *testing.M) {
	createTables()
	code := m.Run()
	for _, t := range tables {
		_, err := db.Exec(fmt.Sprintf("drop table %s;", t))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	os.Exit(code)
}

func createTables() {
	sqlite, err := sql.Open("sqlite3", source)
	if err != nil {
		log.Fatal(err.Error())
	}

	query :=
		`create table users(
			id integer primary key not null
		);
			
		create table segments(
			id integer primary key autoincrement,
			name varchar not null unique
		);
			
		create table history(
			id integer primary key autoincrement,
			user_id integer not null,
			segment_name varchar not null,
			operation varchar not null,
			operation_time timestamp default CURRENT_TIMESTAMP,
			foreign key (user_id) references users(id)
		);
			
		create table follows(
			user_id integer not null,
			segment_id integer not null,
			expire timestamp default null,
			unique (user_id, segment_id),
			foreign key (user_id) references users (id) on delete cascade,
			foreign key (segment_id) references segments (id) on delete cascade
		);
		`

	_, err = sqlite.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	db = sqlite
}

func clearDB() {
	for _, t := range tables {
		_, err := db.Exec(fmt.Sprintf("delete from %s;", t))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}
