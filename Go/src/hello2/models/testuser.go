package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func PrintUsers() {

	//打开驱动
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()

	if err != nil {
		return
	}

	//连接数据库
	stmt, err := db.Prepare("select *from user limit 10")
	defer stmt.Close()
	if err != nil {
		return
	}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
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
