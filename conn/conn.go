package conn

import (
	"database/sql"
	"fmt"
	"log"

	lg "github.com/bahodurnazarov/exel/utils"
	_ "github.com/lib/pq"
)

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Created_on sql.NullTime
	last_login sql.NullTime
}

func Conn() {
	psqlInfo := "postgresql://postgres:postgres@localhost/exel?sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		lg.Errl.Panic(err)
	}
	defer db.Close()

	var myUser User

	userSql, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	defer userSql.Close()

	for userSql.Next() {
		err := userSql.Scan(&myUser.ID, &myUser.Name, &myUser.Email, &myUser.Password, &myUser.Created_on, &myUser.last_login)
		if err != nil {
			panic(err)
		}
		fmt.Println(myUser.ID, " ", myUser.Name)
	}
	err = userSql.Err()
	if err != nil {
		panic(err)
	}
}
