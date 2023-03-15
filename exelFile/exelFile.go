package exelFile

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	lg "github.com/bahodurnazarov/exel/utils"
	"github.com/xuri/excelize/v2"
)

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Created_on sql.NullTime
	last_login sql.NullTime
}

func MakeFile(data *sql.Rows) {
	f, err := excelize.OpenFile("simple.xlsx")
	if err != nil {
        log.Fatal(err)
    }

	var myUser User
	var ff int = 0
	for data.Next() {
		err := data.Scan(&myUser.ID, &myUser.Name, &myUser.Email, &myUser.Password, &myUser.Created_on, &myUser.last_login)
		if err != nil {
			lg.Errl.Panic(err)
		}
		for i := 0; i < myUser.ID; i++ {

			fmt.Println(myUser.ID, " ", myUser.Name)
		}
		ff++
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(+3), myUser.ID)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(+3), myUser.Name)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(+3), myUser.Password)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(+3), myUser.Email)

		//lg.Server.Println(myUser.ID, " ", myUser.Name)

	}
	if err := f.SaveAs("simple.xlsx"); err != nil {
		log.Fatal(err)
	}
}
