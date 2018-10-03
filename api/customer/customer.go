package customer

import (
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/middleware"
	"qrapi-logistic/o/order"
	"qrapi-logistic/o/product"
)

type CustomerServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewCustomerServer(parent *gin.RouterGroup, name string) *CustomerServer {
	var s = CustomerServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeCustomer)
	//order api
	s.POST("order/create", s.createOrder)
	s.POST("order/update", s.updateOrder)
	s.GET("order/delete", s.deleteOrder)
	s.GET("order/list", s.getOrders)
	//employee api
	s.GET("employee/list", s.getEmployees)
	s.POST("employee/create", s.createEmployee)
	s.GET("employee/delete", s.deleteEmployee)
	//process
	s.GET("process/list", s.getProcesses)
	s.POST("process/create", s.createProcess)

	return &s
}

func (s *CustomerServer) createProduct(c *gin.Context) {
	var product *product.Product
	c.BindJSON(&product)
	web.AssertNil(product.Create())
	s.SendData(c, product)
}

func (s *CustomerServer) createOrder(c *gin.Context) {
	var order *order.Order
	web.AssertNil(c.BindJSON(&order))
	web.AssertNil(order.Create())
	s.SendData(c, order)
}

func (s *CustomerServer) updateOrder(c *gin.Context) {
	var order *order.Order
	web.AssertNil(c.BindJSON(&order))
	web.AssertNil(order.Create())
	s.SendData(c, order)
}

func (s *CustomerServer) deleteOrder(c *gin.Context) {
	web.AssertNil(order.DeleteByID(c.Query("id")))
	s.Success(c)
}

func (s *CustomerServer) getOrders(c *gin.Context) {
	var customerID = c.MustGet("user_id")
	var orders, err = order.GetOrdersByCustomer(customerID.(string))
	web.AssertNil(err)
	s.SendData(c, orders)
}
