package database

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/packages/setting"
)

var dataSourceName string

func Connect() *ent.Client {
	dataSourceName = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Sslmode,
		// setting.DatabaseSetting.Binary_parameters,
	)

	sqlDriver, err := sql.Open(dialect.Postgres, dataSourceName)
	if err != nil {
		panic(err)
	}

	db := sqlDriver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(sqlDriver))

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

	// defer migrateDB(sqlDriver)

	// if setting.AppSetting.Env == "dev" || setting.AppSetting.Env == "docker" {
	// 	Seed(ctx, client)
	// }
	return client

}
