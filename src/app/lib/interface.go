package lib

import "main-api-store-management/app/models"

type StoreManagement interface {
	CreateUser(user *models.User) error
	GetUserByID(userID string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateStore(store *models.Store, user *models.User) error
	GetAllStores(user *models.User) ([]*models.Store, error)
	CreatePolicy(store *models.Store, user *models.User) error
	DeleteStoreByID(storeID string) error
}
