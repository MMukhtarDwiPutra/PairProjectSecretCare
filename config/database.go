package config

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func InitDatabase(url string) *sql.DB{
	db, err := sql.Open("mysql", url)
	if err != nil{
		fmt.Println(err.Error())
		return db
	}

	fmt.Println("Successfuly init new connection")

	return db
}