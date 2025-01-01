package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	controllers "restaraunt_golang/controllers"
)

func InvoiceRoutes(routes *gin.Engine, db *gorm.DB) {
	routes.GET("/invoices", controllers.GetInvoices(db))
	routes.GET("/invoices/:invoice_id", controllers.GetInvoice(db))
	routes.POST("/invoices", controllers.CreateInvoice(db))
	routes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice(db))
}
