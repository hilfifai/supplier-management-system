package services

import (
	"supplier-management-backend/models"

	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

func NewDashboardService(db *gorm.DB) *DashboardService {
	return &DashboardService{db: db}
}

func (s *DashboardService) GetDashboardStats() (*models.DashboardStats, error) {
	var stats models.DashboardStats

	var totalSuppliers int64
	err := s.db.Model(&models.Supplier{}).Count(&totalSuppliers).Error
	if err != nil {
		return nil, err
	}
	stats.TotalSuppliers = int(totalSuppliers)

	var newSuppliers int64
	err = s.db.Model(&models.Supplier{}).Where("created_at >= NOW() - INTERVAL '30 days'").Count(&newSuppliers).Error
	if err != nil {
		return nil, err
	}
	stats.NewSuppliers = int(newSuppliers)

	var blockedSuppliers int64
	err = s.db.Model(&models.Supplier{}).Where("status = ?", "Blocked").Count(&blockedSuppliers).Error
	if err != nil {
		return nil, err
	}
	stats.BlockedSuppliers = int(blockedSuppliers)

	var totalOrders int64
	err = s.db.Model(&models.OrderHistory{}).Count(&totalOrders).Error
	if err != nil {
		return nil, err
	}
	stats.TotalOrders = int(totalOrders)

	var onTimeDelivery int64
	err = s.db.Model(&models.OrderHistory{}).Where("order_status = ?", "Delivered").Count(&onTimeDelivery).Error
	if err != nil {
		return nil, err
	}
	stats.OnTimeDelivery = int(onTimeDelivery)

	var lateDelivery int64
	err = s.db.Model(&models.OrderHistory{}).Where("order_status = ?", "Late").Count(&lateDelivery).Error
	if err != nil {
		return nil, err
	}
	stats.LateDelivery = int(lateDelivery)

	var ordersInProgress int64
	err = s.db.Model(&models.OrderHistory{}).Where("order_status = ?", "In Progress").Count(&ordersInProgress).Error
	if err != nil {
		return nil, err
	}
	stats.OrdersInProgress = int(ordersInProgress)

	var pendingOrders int64
	err = s.db.Model(&models.OrderHistory{}).Where("order_status = ?", "Pending").Count(&pendingOrders).Error
	if err != nil {
		return nil, err
	}
	stats.PendingOrders = int(pendingOrders)

	var totalInvoices int64
	err = s.db.Model(&models.Invoice{}).Count(&totalInvoices).Error
	if err != nil {
		return nil, err
	}
	stats.TotalInvoices = int(totalInvoices)

	stats.InvoicesInProgress = 1869
	stats.PaidInvoices = 999
	stats.OutstandingInvoices = 999
	stats.AvgCostPerSupplier = 320300000

	return &stats, nil
}

