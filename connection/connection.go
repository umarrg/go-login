package connection

import (
	"database/sql"
	"fmt"

	"github.com/umarrg/microservice/myError"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./store.db3")
	myError.HandleErr(err)

	createTableQuery := `create table if not exists students(
		id integer not null primary key autoincrement,
		firstName text,
		lastName text,
		email text,
		gender text,
		age integer
	)`

	res, err := Db.Exec(createTableQuery)
	myError.HandleErr(err)
	fmt.Println(res)
}
