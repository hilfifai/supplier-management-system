package database

import (
	"supplier-management-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Supplier{},
		&models.Address{},
		&models.Contact{},
		&models.SupplierGroup{},
		&models.Material{},
		&models.OtherInformation{},
		&models.OrderHistory{},
		&models.Invoice{},
		&models.ReviewApproval{},
		&models.Workflow{},
	)
}

