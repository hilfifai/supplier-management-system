package routes

import (
	"supplier-management-backend/handlers"
	"supplier-management-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.RouterGroup, db *gorm.DB) {
	supplierService := services.NewSupplierService(db)
	supplierHandler := handlers.NewSupplierHandler(supplierService)

	dashboardService := services.NewDashboardService(db)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	suppliers := r.Group("/suppliers")
	{
		suppliers.GET("", supplierHandler.GetSuppliers)
		suppliers.GET("/:id", supplierHandler.GetSupplier)
		suppliers.POST("", supplierHandler.CreateSupplier)
		suppliers.PUT("/:id", supplierHandler.UpdateSupplier)
		suppliers.DELETE("/:id", supplierHandler.DeleteSupplier)
		suppliers.POST("/:id/block", supplierHandler.BlockSupplier)
		suppliers.POST("/:id/unblock", supplierHandler.UnblockSupplier)
	}

	dashboard := r.Group("/dashboard")
	{
		dashboard.GET("/stats", dashboardHandler.GetStats)
	}
}

