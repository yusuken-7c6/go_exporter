package main

import (
	"fmt"
	"sync"
	"runtime"
	"mysqlclient"
	"time"
)

var PREFIX "SRC_TEST_DB_"

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	sync.Add(runtime.NumCPU())

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	table_name := ARGV[1]
	sql = query_string(table_name, 1, 1000)
	conn := mysqlclient.Connect(conn)
	res := mysqlclient.Execute(conn, sql)
	for rows.Next() {
		var id int
		var name string
		var created_at time.Time
		var description string
		var code string
		var register_date time.Date
		if err := rows.scan(
			&id,
			&name,
			&created_at,
			&description,
			&code,
			&register_date
		); err != nil {
			fmt.Println(
				id,
				name,
				created_at,
				description,
				code,
				register_date
			)
		}
	}
	mysqlclient.Close(conn)
}

func query_string(table_name string, start_id int, end_id int) {
	"select * from " + table_name + " where id > " + str(start_id) + " and id <= " + end_id + ";"
}