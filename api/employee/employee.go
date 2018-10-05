package employee

import (
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/o/auth"
	"qrapi-logistic/o/employee"
	"qrapi-logistic/o/order"
	"qrapi-logistic/o/order_history"
)

type EmployeeServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewEmployeeServer(parent *gin.RouterGroup, name string) *EmployeeServer {
	var s = EmployeeServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("login", s.login)
	s.GET("order/get", s.getOrderByID)
	s.POST("work/done", s.workDone)
	s.GET("order/history/get", s.getOrderHistory)
	s.POST("order/history/create", s.createOrderHistory)
	return &s
}

func (s *EmployeeServer) login(c *gin.Context) {
	var loginInfo = struct {
		Uname    string
		Password string
	}{}
	c.BindJSON(&loginInfo)
	user, err := employee.Login(loginInfo.Uname, loginInfo.Password)
	web.AssertNil(err)
	var auth = auth.Create(user.ID, "employee")
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *EmployeeServer) getOrderByID(c *gin.Context) {
	var orderID = c.Query("order_id")
	var order, err = order.GetOrderByIDAndEmployee(orderID)
	web.AssertNil(err)
	s.SendData(c, order)
}

func (s *EmployeeServer) workDone(c *gin.Context) {
	var orh *order_history.OrderHistory
	web.AssertNil(c.BindJSON(&orh))
	web.AssertNil(orh.Create())
	s.Success(c)
}

func (s *EmployeeServer) getOrderHistory(c *gin.Context) {
	var productID = c.Query("product_id")
	var orderID = c.Query("order_id")
	var result, err = order_history.GetHistoryByID(orderID, productID)
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *EmployeeServer) createOrderHistory(c *gin.Context) {
	var orh *order_history.OrderHistory
	web.AssertNil(c.BindJSON(&orh))
	web.AssertNil(orh.Create())
	s.Success(c)
}
