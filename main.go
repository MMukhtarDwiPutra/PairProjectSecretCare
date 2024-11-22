package main

import (
    "SecretCare/cli"
    "SecretCare/config"
    "SecretCare/handler"
    "context"
)

func main() {
    connString := "postgresql://postgres.tsbsgibxzmhmjoaifosb:AnakGanteng@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

    db := config.InitDatabase(connString)
    defer db.Close()

    var ctx context.Context = context.Background()

    handler := handler.NewHandler(ctx, db)
    cli := cli.NewCli(handler, ctx)

    cli.MenuUtama()
}
