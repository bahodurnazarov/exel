package slct

import (
	"database/sql"

	exel "github.com/bahodurnazarov/exel/exelFile"
	lg "github.com/bahodurnazarov/exel/utils"
)

func GetData(db *sql.DB) {

	userSql, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		lg.Errl.Fatal("Failed to execute query: ", err)
	}
	defer userSql.Close()

	exel.MakeFile(userSql)

}
