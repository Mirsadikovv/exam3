package apis

import (
	"login/api/handler"
	"login/pkg/logger"
	"login/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(service, log)

	r := gin.Default()

	// r.Use(authMiddleware)

	r.POST("/customer", h.CreateCustomer)
	r.PUT("/customer/update/:id", h.UpdateCustomer)
	r.GET("/customer", h.GetAllCustomers)
	r.GET("customer/:id", h.GetCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)
	r.PATCH("/customer/update_birthday/:id", h.UpdateBirthday)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
