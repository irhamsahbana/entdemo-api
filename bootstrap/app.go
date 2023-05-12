package bootstrap

import (
	"context"
	"entdemo-api/ent"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

var App *Application

type Application struct {
	Ent *ent.Client
}

func init() {
	initApp()
}

func initApp() {
	App = &Application{}
	App.Ent = InitEntGo()
}

func InitEntGo() *ent.Client {
	driver := os.Getenv(`DB_DRIVER`)
	host := os.Getenv(`DB_HOST`)
	port := os.Getenv(`DB_PORT`)
	database := os.Getenv(`DB_DATABASE`)
	user := os.Getenv(`DB_USERNAME`)
	pass := os.Getenv(`DB_PASSWORD`)

	if driver == "" {
		driver = dialect.MySQL
	}

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "3306"
	}

	if database == "" {
		database = "entdemo_api"
	}

	if user == "" {
		user = "root"
	}

	if pass == "" {
		pass = ""
	}

	dsn := ""
	switch driver {
	case dialect.MySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, pass, host, port, database)
	case dialect.Postgres:
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, database, pass)
	case dialect.SQLite:
		dsn = fmt.Sprintf("file:%s?cache=shared&mode=rwc", database)
	default:
		log.Fatal("Unknown driver")
	}

	client, err := ent.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
