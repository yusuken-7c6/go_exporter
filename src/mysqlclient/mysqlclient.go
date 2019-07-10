package mysqlclient

import (
	"dbconfig"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // initだけする
)

const DB_TYPE = "mysql"

func Close(conn *sql.DB) {
	conn.Close()
}

func Execute(conn *sql.DB, query string) {
	rows, err := conn.Query(query)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}
	rows
}

func Connect(env_prefix string) *sql.DB {
	dbconf := dbconfig.Generate(env_prefix)
	conn, err := sql.Open(
			DB_TYPE,
			dbconf.Username + ":" + dbconf.Password + "@tcp(" + dbconf.Hostname + ":)/" + dbconf.Name)

	&conn
}
