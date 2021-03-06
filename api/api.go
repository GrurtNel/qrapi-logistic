package api

import (
	"qrapi-logistic/api/admin"
	"qrapi-logistic/api/auth"
	"qrapi-logistic/api/customer"
	"qrapi-logistic/api/employee"
	"qrapi-logistic/api/public"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	admin.NewAdminServer(root, "admin")
	auth.NewAuthServer(root, "auth")
	public.NewPublicServer(root, "public")
	employee.NewEmployeeServer(root, "employee")
	customer.NewCustomerServer(root, "customer")
}
