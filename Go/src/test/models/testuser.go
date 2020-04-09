package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func PrintUsers() {
	fmt.Println("PrintUsers1")
	db, err := sql.Open("mysql", "root:2422503084.@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()

	if err != nil {
		fmt.Println("PrintUsers2")
		return
	}

	stmt, err := db.Prepare("select * from user")
	defer stmt.Close()
	if err != nil {
		fmt.Println("PrintUsers3")
		return
	}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Println("PrintUsers")
		return
	}

	for rows.Next() {
		var id int
		var name string
		var sex int8
		rows.Scan(&id, &name, &sex)

		fmt.Println(id, name, sex)
	}

}
