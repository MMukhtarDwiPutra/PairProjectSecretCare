package main

import (
	"SecretCare/cli"
	"SecretCare/config"
	"SecretCare/handler"
	"context"
)

func main() {

	db := config.InitDatabase("root:a!@tcp(127.0.0.1:3306)/SecretCare")
	defer db.Close()

	var ctx context.Context = context.Background()
	
	handler := handler.NewHandler(ctx, db)
	cli := cli.NewCli(handler, ctx)


	cli.MenuUtama()
}
