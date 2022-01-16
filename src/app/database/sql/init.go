package db

import (
	"gorm.io/gorm"
	"main-api-store-management/app/lib"
)

type Database struct {
	lib.StoreManagement
	DB *gorm.DB
}
