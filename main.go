package main

import(
	"SecretCare/config"
	"SecretCare/handler"
	"SecretCare/cli"
)

func main(){

	db := config.InitDatabase("root:@tcp(127.0.0.1:3306)/SecretCare2")
	defer db.Close()
	
	handlerUser := handler.NewHandlerUser(db)
	handlerProduct := handler.NewHandlerProduct(db)
	cli := cli.NewCli(handlerUser, handlerProduct)

	cli.MenuUtama()
}