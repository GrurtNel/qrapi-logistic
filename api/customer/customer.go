package customer

import (
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/math"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/middleware"
	"qrapi-logistic/o/order"
	"qrapi-logistic/o/order_history"
	"qrapi-logistic/o/product"
	"strconv"
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
	s.GET("order/history", s.getOrderHistory)
	s.GET("order/tracking", s.getOrderTracking)
	s.GET("order/generate", s.generateCSV)
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

func (s *CustomerServer) getOrderHistory(c *gin.Context) {
	var orderID = c.Query("order_id")
	var result, err = order_history.GetHistoryByID(orderID, "")
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *CustomerServer) getOrderTracking(c *gin.Context) {
	var orderID = c.Query("order_id")
	var result, err = order_history.GetTrackingByID(orderID)
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *CustomerServer) getOrders(c *gin.Context) {
	var customerID = c.MustGet("user_id")
	var orders, err = order.GetOrdersByCustomer(customerID.(string))
	web.AssertNil(err)
	s.SendData(c, orders)
}

func (s *CustomerServer) generateCSV(c *gin.Context) {
	var quantity, _ = strconv.Atoi(c.Query("quantity"))
	var orderID = c.Query("order_id")
	var order, err = order.GetOrderByID(orderID)
	web.AssertNil(err)
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	var record []string
	for i := 0; i < quantity; i++ {
		record = []string{orderID + "$$" + math.RandString("P", 11)}
		wr.Write(record)
	}
	wr.Flush()                                        // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))
	c.Writer.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+order.Name+".csv")
	c.Writer.Write(b.Bytes())
}
