package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	sqldb "main-api-store-management/app/database/sql"
	"main-api-store-management/app/lib"
	"main-api-store-management/app/models"
	"sync"
)

var onceDataEngine sync.Once
var onceMySQL sync.Once
var databaseGetter func() (lib.StoreManagement, error)
var SQL *sqldb.Database
var connectionErr error

func GetDatabase() (lib.StoreManagement, error) {
	onceDataEngine.Do(func() {
		switch lib.DB_ENGINE {
		case "SQL":
			databaseGetter = newSQLDatabase
		default:
			databaseGetter = func() (lib.StoreManagement, error) {
				return nil, fmt.Errorf("Unknown DB_ENGINE: '%s'.", lib.DB_ENGINE)
			}
		}
	})
	return databaseGetter()
}

func newSQLDatabase() (lib.StoreManagement, error) {
	onceMySQL.Do(func() {
		SQL = new(sqldb.Database)
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", lib.DB_USER, lib.DB_PASSWORD, lib.DB_HOST, lib.DB_PORT, lib.DB_NAME)
		database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			connectionErr = err
			return
		}
		connectionErr = migration(database)
		SQL.DB = database
	})
	return SQL, connectionErr
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
