package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	postgresdb "main-api-store-management/app/database/postgres"
	"main-api-store-management/app/lib"
	"main-api-store-management/app/models"
	"sync"
)

var onceDataEngine sync.Once
var onceMySQL sync.Once
var databaseGetter func() (lib.StoreManagement, error)
var postgresql *postgresdb.Database
var connectionErr error

func GetDatabase() (lib.StoreManagement, error) {
	onceDataEngine.Do(func() {
		switch lib.DB_ENGINE {
		case "PostgreSQL":
			databaseGetter = newPostgreSQLDatabase
		default:
			databaseGetter = func() (lib.StoreManagement, error) {
				return nil, fmt.Errorf("Unknown DB_ENGINE: '%s'.", lib.DB_ENGINE)
			}
		}
	})
	return databaseGetter()
}

func newPostgreSQLDatabase() (lib.StoreManagement, error) {
	onceMySQL.Do(func() {
		postgresql = new(postgresdb.Database)
		dnsPostgres := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			lib.DB_HOST,
			lib.DB_USER,
			lib.DB_PASSWORD,
			lib.DB_NAME,
			lib.DB_PORT)

		database, err := gorm.Open(postgres.Open(dnsPostgres), &gorm.Config{})
		if err != nil {
			connectionErr = err
			return
		}
		connectionErr = migration(database)
		postgresql.DB = database
	})
	return postgresql, connectionErr
}

func migration(db *gorm.DB) error {
	schema := []interface{}{
		&models.Store{},
		&models.Bill{},
		&models.Category{},
		&models.CategoryOption{},
		&models.CategoryOptionValue{},
		&models.Cheque{},
		&models.Necessity{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Product{},
		&models.Report{},
		&models.User{},
		&models.Policy{},
	}
	return db.AutoMigrate(schema...)
}
