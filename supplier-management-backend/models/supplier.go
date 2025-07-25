package models

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID                 string              `json:"id" gorm:"primaryKey"`
	Name               string              `json:"name" gorm:"not null"`
	LogoURL            string              `json:"logo_url"`
	NickName           string              `json:"nick_name"`
	Status             string              `json:"status" gorm:"not null"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	Addresses          []Address           `json:"addresses,omitempty" gorm:"foreignKey:SupplierID"`
	Contacts           []Contact           `json:"contacts,omitempty" gorm:"foreignKey:SupplierID"`
	Groups             []SupplierGroup     `json:"groups,omitempty" gorm:"foreignKey:SupplierID"`
	Materials          []Material          `json:"materials,omitempty" gorm:"foreignKey:SupplierID"`
	OtherInformations  []OtherInformation  `json:"other_informations,omitempty" gorm:"foreignKey:SupplierID"`
	OrderHistories     []OrderHistory      `json:"order_histories,omitempty" gorm:"foreignKey:SupplierID"`
	Invoices           []Invoice           `json:"invoices,omitempty" gorm:"foreignKey:SupplierID"`
}

type Address struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	SupplierID string    `json:"supplier_id" gorm:"not null"`
	Type       string    `json:"type"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	IsMain     bool      `json:"is_main" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Contact struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	SupplierID  string    `json:"supplier_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	JobPosition string    `json:"job_position"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Mobile      string    `json:"mobile"`
	IsMain      bool      `json:"is_main" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SupplierGroup struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	SupplierID string    `json:"supplier_id" gorm:"not null"`
	GroupName  string    `json:"group_name"`
	Value      string    `json:"value"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Material struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	SupplierID    string    `json:"supplier_id" gorm:"not null"`
	MaterialGroup string    `json:"material_group"`
	MaterialID    string    `json:"material_id"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	Price         float64   `json:"price"`
	DeliveryTime  string    `json:"delivery_time"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type OtherInformation struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	SupplierID    string    `json:"supplier_id" gorm:"not null"`
	AttributeName string    `json:"attribute_name"`
	Value         string    `json:"value"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type OrderHistory struct {
	ID                    string     `json:"id" gorm:"primaryKey"`
	SupplierID            string     `json:"supplier_id" gorm:"not null"`
	OrderIDPO             string     `json:"order_id_po"`
	ShipmentDate          *time.Time `json:"shipment_date"`
	OrderStatus           string     `json:"order_status"`
	EstimatedDeliveryDate *time.Time `json:"estimated_delivery_date"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

type Invoice struct {
	ID                     string     `json:"id" gorm:"primaryKey"`
	SupplierID             string     `json:"supplier_id" gorm:"not null"`
	InvoiceNumber          string     `json:"invoice_number"`
	ProjectName            string     `json:"project_name"`
	Amount                 float64    `json:"amount"`
	AgingDays              int        `json:"aging_days"`
	ReceivedDate           *time.Time `json:"received_date"`
	EstimatedPaymentDate   *time.Time `json:"estimated_payment_date"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

type ReviewApproval struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	ProcessName  string     `json:"process_name"`
	CustomerID   string     `json:"customer_id"`
	CustomerName string     `json:"customer_name"`
	StageFlow    string     `json:"stage_flow"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Workflows    []Workflow `json:"workflows,omitempty" gorm:"foreignKey:ReviewApprovalID"`
}

type Workflow struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	ReviewApprovalID string    `json:"review_approval_id" gorm:"not null"`
	StageName        string    `json:"stage_name"`
	SLAValue         string    `json:"sla_value"`
	SLAUnit          string    `json:"sla_unit"`
	IsActive         bool      `json:"is_active" gorm:"default:true"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = generateID("SUPP")
	}
	return nil
}

func (a *Address) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = generateID("ADDR")
	}
	return nil
}

func (c *Contact) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = generateID("CONT")
	}
	return nil
}

func (sg *SupplierGroup) BeforeCreate(tx *gorm.DB) error {
	if sg.ID == "" {
		sg.ID = generateID("SGRP")
	}
	return nil
}

func (m *Material) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = generateID("MATL")
	}
	return nil
}

func (oi *OtherInformation) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = generateID("OTHR")
	}
	return nil
}

func (oh *OrderHistory) BeforeCreate(tx *gorm.DB) error {
	if oh.ID == "" {
		oh.ID = generateID("ORDR")
	}
	return nil
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) error {
	if i.ID == "" {
		i.ID = generateID("INVC")
	}
	return nil
}

func (ra *ReviewApproval) BeforeCreate(tx *gorm.DB) error {
	if ra.ID == "" {
		ra.ID = generateID("REVW")
	}
	return nil
}

func (w *Workflow) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = generateID("WFLW")
	}
	return nil
}

