package exelFile

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

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
		lg.Server.Fatal(err)
	}

	var myUser User
	var ff int = 0
	for data.Next() {
		err := data.Scan(&myUser.ID, &myUser.Name, &myUser.Email, &myUser.Password, &myUser.Created_on, &myUser.last_login)
		if err != nil {
			lg.Errl.Panic(err)
		}
		ff++

		//lg.Server.Println(myUser.ID, " ", myUser.Name)

		fmt.Println("ID: ", myUser.Name, myUser.Password)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(2+ff), myUser.ID)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(2+ff), myUser.Name)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(2+ff), myUser.Email)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(2+ff), myUser.Password)
	}
	now := time.Now()
	f.SetCellValue("Sheet1", "B1", now.Format(time.ANSIC))

	if err := f.SaveAs("simple.xlsx"); err != nil {
		lg.Server.Fatal(err)
	}
}
