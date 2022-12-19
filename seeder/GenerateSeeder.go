package seeder

import (
	"fmt"
	"github.com/balqisgautama/okami-auth/server"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
	"strconv"
)

func DBMigrate() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./postgreSQL"),
	}
	n, err := migrate.Exec(server.ServerConfig.DBConnection, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
		os.Exit(3) // exit 3
	} else {
		fmt.Println("Applied " + strconv.Itoa(n) + " migrations!")
	}
}
