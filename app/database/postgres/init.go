package db

import (
	"gorm.io/gorm"
	"main-api-store-management/app/lib"
	"main-api-store-management/app/models"
)

type Database struct {
	lib.StoreManagement
	DB *gorm.DB
}

func (db *Database) GetUserByID(userID string) (*models.User, error) {
	user := new(models.User)
	err := db.DB.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (db *Database) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := db.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (db *Database) CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (db *Database) CreateStore(store *models.Store, user *models.User) error {
	err := db.DB.Create(store).Error
	if err != nil {
		return err
	}
	err = db.CreatePolicy(store, user, "owner")
	return err
}

func (db *Database) GetAllStores(user *models.User) ([]*models.Store, error) {
	stores := make([]*models.Store, 20)
	err := db.DB.Select("stores.*").Joins("JOIN policies ON stores.id = policies.store_id").Where("policies.user_id = ?", user.ID).Find(&stores).Error
	return stores, err
}

func (db *Database) CreatePolicy(store *models.Store, user *models.User, role string) error {
	policy := new(models.Policy)
	policy.Role = role
	policy.StoreID = store.ID
	policy.UserID = user.ID
	return db.DB.Create(policy).Error
}

func (db *Database) DeleteStoreByID(storeID string) error {
	store := new(models.Store)
	err := db.DB.Where("id = ?", storeID).Unscoped().Delete(store).Error
	return err
}
