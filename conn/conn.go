package conn

import (
	"database/sql"

	slct "github.com/bahodurnazarov/exel/select"
	lg "github.com/bahodurnazarov/exel/utils"
	_ "github.com/lib/pq"
)

func Conn() {
	psqlInfo := "postgresql://postgres:postgres@localhost/exel?sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		lg.Errl.Panic(err)
	}
	defer db.Close()

	slct.GetData(db)

}
