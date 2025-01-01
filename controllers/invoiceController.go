package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"restaraunt_golang/models"
)

// GetInvoices returns a list of all invoices.
func GetInvoices(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invoices []models.Invoice
		if err := db.Find(&invoices).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch invoices"})
			return
		}
		ctx.JSON(http.StatusOK, invoices)
	}
}

// GetInvoice fetches a specific invoice by ID.
func GetInvoice(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoiceID := ctx.Param("invoice_id")
		var invoice models.Invoice
		if err := db.Where("id = ?", invoiceID).First(&invoice).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
			return
		}
		ctx.JSON(http.StatusOK, invoice)
	}
}

// CreateInvoice creates a new invoice.
func CreateInvoice(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invoice models.Invoice
		if err := ctx.ShouldBindJSON(&invoice); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// You can also add some logic to fetch the associated order and ensure it exists.
		if err := db.Create(&invoice).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
			return
		}

		ctx.JSON(http.StatusCreated, invoice)
	}
}

// UpdateInvoice updates an existing invoice.
func UpdateInvoice(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoiceID := ctx.Param("invoice_id")
		var invoice models.Invoice
		if err := db.Where("id = ?", invoiceID).First(&invoice).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
			return
		}

		// Bind new data to invoice struct
		if err := ctx.ShouldBindJSON(&invoice); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// Update the invoice in DB
		if err := db.Save(&invoice).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice"})
			return
		}

		ctx.JSON(http.StatusOK, invoice)
	}
}
