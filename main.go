package main

import (
	"SecretCare/cli"
	"SecretCare/config"
	"SecretCare/handler"
)

type User struct {
	username string
}

type ContextKey string

const UserContextKey ContextKey = "user"

func main() {

	db := config.InitDatabase("root:@tcp(127.0.0.1:3306)/SecretCare2")
	defer db.Close()

	handler := handler.NewHandler(db)
	cli := cli.NewCli(handler)

	cli.MenuUtama()
}
