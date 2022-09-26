package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/201R/go_api_boilerplate/ent"
	_ "github.com/201R/go_api_boilerplate/ent/runtime"
	"github.com/201R/go_api_boilerplate/packages/setting"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var db_url string

func Connect() *ent.Client {

	db_url = fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		setting.DatabaseSetting.Type,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Sslmode,
	)

	db, err := sql.Open("pgx", db_url)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	drv := entsql.OpenDB(dialect.Postgres, db)

	if setting.AppSetting.Env == "docker" {
		<-time.After(42 * time.Second)
	}

	// err = client.Schema.Create(
	// 	ctx,
	// 	migrate.WithDropIndex(true),
	// 	migrate.WithFixture(true),
	// )
	// if err != nil {
	// 	logger.Fatalf("failed creating schema resources: %v", err)
	// }

	// defer migrateDB(drv)

	// if setting.AppSetting.Env == "dev" || setting.AppSetting.Env == "docker" {
	// 	Seed(ctx, client)
	// }
	return ent.NewClient(ent.Driver(drv))

}

// func migrateDB(sqlDriver *entsql.Driver) {
// 	var m *gomigrate.Migrate
// 	var err error

// 	pg, _ := mpg.WithInstance(sqlDriver.DB(), &mpg.Config{})
// 	m, err = gomigrate.NewWithDatabaseInstance("file://migrations", "postgre", pg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := m.Up(); err != nil && err != gomigrate.ErrNoChange {
// 		log.Fatal(err)
// 	}
// }
