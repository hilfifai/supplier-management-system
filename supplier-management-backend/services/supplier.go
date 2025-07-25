package services

import (
	"supplier-management-backend/models"

	"gorm.io/gorm"
)

type SupplierService struct {
	db *gorm.DB
}

func NewSupplierService(db *gorm.DB) *SupplierService {
	return &SupplierService{db: db}
}

func (s *SupplierService) GetSuppliers(page, limit int, search, status string) ([]models.Supplier, int, error) {
	var suppliers []models.Supplier
	var total int64

	query := s.db.Model(&models.Supplier{})

	if search != "" {
		query = query.Where("name ILIKE ? OR nick_name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&suppliers).Error
	if err != nil {
		return nil, 0, err
	}

	return suppliers, int(total), nil
}

func (s *SupplierService) GetSupplierByID(id string) (*models.Supplier, error) {
	var supplier models.Supplier
	err := s.db.Preload("Addresses").
		Preload("Contacts").
		Preload("Groups").
		Preload("Materials").
		Preload("OtherInformations").
		Preload("OrderHistories").
		Preload("Invoices").
		First(&supplier, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (s *SupplierService) CreateSupplier(supplier *models.Supplier) (*models.Supplier, error) {
	err := s.db.Create(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *SupplierService) UpdateSupplier(supplier *models.Supplier) (*models.Supplier, error) {
	err := s.db.Save(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *SupplierService) DeleteSupplier(id string) error {
	return s.db.Delete(&models.Supplier{}, "id = ?", id).Error
}

func (s *SupplierService) BlockSupplier(id, reason string) error {
	return s.db.Model(&models.Supplier{}).Where("id = ?", id).Update("status", "Blocked").Error
}

func (s *SupplierService) UnblockSupplier(id, reason string) error {
	return s.db.Model(&models.Supplier{}).Where("id = ?", id).Update("status", "Active").Error
}

