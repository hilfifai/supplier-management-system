package handlers

import (
	"net/http"
	"strconv"
	"supplier-management-backend/models"
	"supplier-management-backend/services"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{service: service}
}

func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	status := c.Query("status")

	suppliers, total, err := h.service.GetSuppliers(page, limit, search, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	totalPages := (total + limit - 1) / limit

	c.JSON(http.StatusOK, models.APIResponse{
		Data: suppliers,
		Pagination: &models.PaginationResponse{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")

	supplier, err := h.service.GetSupplierByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Error: &models.APIError{
				Code:    "NOT_FOUND",
				Message: "Supplier not found",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data: supplier,
	})
}

func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: &models.APIError{
				Code:    "VALIDATION_ERROR",
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	createdSupplier, err := h.service.CreateSupplier(&supplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Data: createdSupplier,
	})
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: &models.APIError{
				Code:    "VALIDATION_ERROR",
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	supplier.ID = id
	updatedSupplier, err := h.service.UpdateSupplier(&supplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data: updatedSupplier,
	})
}

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteSupplier(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data: map[string]string{"message": "Supplier deleted successfully"},
	})
}

func (h *SupplierHandler) BlockSupplier(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Reason string `json:"reason" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: &models.APIError{
				Code:    "VALIDATION_ERROR",
				Message: "Reason is required",
			},
		})
		return
	}

	err := h.service.BlockSupplier(id, request.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data: map[string]string{"message": "Supplier blocked successfully"},
	})
}

func (h *SupplierHandler) UnblockSupplier(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Reason string `json:"reason" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: &models.APIError{
				Code:    "VALIDATION_ERROR",
				Message: "Reason is required",
			},
		})
		return
	}

	err := h.service.UnblockSupplier(id, request.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data: map[string]string{"message": "Supplier unblocked successfully"},
	})
}

