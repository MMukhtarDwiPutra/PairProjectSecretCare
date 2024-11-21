package main

import (
	"SecretCare/cli"
	"SecretCare/config"
	"SecretCare/handler"
	"context"
)

func main() {

	db := config.InitDatabase("root:@tcp(127.0.0.1:3307)/SecretCare")
	defer db.Close()

	var ctx context.Context = context.Background()
	
	handlerUser := handler.NewHandlerUser(db)
	handlerProduct := handler.NewHandlerProduct(db)
	cli := cli.NewCli(handlerUser, handlerProduct)


	cli.MenuUtama()
}
