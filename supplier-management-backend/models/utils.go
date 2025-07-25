package models

import (
	"fmt"
	"math/rand"
	"time"
)

func generateID(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%d%d", prefix, time.Now().Unix(), rand.Intn(10000))
}

type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type APIResponse struct {
	Data       interface{}         `json:"data,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
	Error      *APIError           `json:"error,omitempty"`
}

type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type DashboardStats struct {
	TotalSuppliers        int     `json:"total_suppliers"`
	NewSuppliers          int     `json:"new_suppliers"`
	AvgCostPerSupplier    float64 `json:"avg_cost_per_supplier"`
	BlockedSuppliers      int     `json:"blocked_suppliers"`
	TotalOrders           int     `json:"total_orders"`
	OnTimeDelivery        int     `json:"on_time_delivery"`
	LateDelivery          int     `json:"late_delivery"`
	OrdersInProgress      int     `json:"orders_in_progress"`
	PendingOrders         int     `json:"pending_orders"`
	TotalInvoices         int     `json:"total_invoices"`
	InvoicesInProgress    int     `json:"invoices_in_progress"`
	PaidInvoices          int     `json:"paid_invoices"`
	OutstandingInvoices   int     `json:"outstanding_invoices"`
}

