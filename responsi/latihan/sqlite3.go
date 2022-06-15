package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"gorm.io/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type mahasiswa struct {
	ID    int
	NAMA  string
	NIM int
	ALAMAT string
	JK string
}

func (item mahasiswa) String() string {
	return fmt.Sprintf("%d, %s, %s, %d, %s", item.ID, item.NAMA, item.NIM, item.ALAMAT, item.JK)
}

func main() {
	db, err := sql.Open("sqlite3", "./ResponsiTekcloud.db")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected")
		// var id = 1
		rows, err := db.Query("select * from MAHASISWA ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []mahasiswa

		for rows.Next() {
			var each = mahasiswa{}
			var err = rows.Scan(&each.ID, &each.NAMA, &each.NIM, &each.ALAMAT, &each.JK)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, each := range result {
			fmt.Println(each.String())
		}
	}
	defer db.Close()