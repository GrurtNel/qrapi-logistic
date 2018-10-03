package admin

import (
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/middleware"
	"qrapi-logistic/o/admin"
	"qrapi-logistic/o/auth"
	"qrapi-logistic/o/customer"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("auth/login", s.login)
	s.Use(middleware.MustBeAdmin)
	//product api
	s.POST("product/create", s.createProduct)
	s.GET("product/list", s.getProducts)
	s.GET("product/delete", s.deleteProduct)
	s.POST("product/update", s.updateProduct)
	//order api
	s.GET("order/delivery", s.deliveryOrder)
	s.GET("order/list", s.getOrders)
	s.POST("order/create", s.createOrder)
	s.POST("order/update", s.updateOrder)
	s.GET("order/delete", s.deleteOrder)
	s.GET("order/generate", s.generateCSV)
	//customer api
	s.GET("customer/list", s.getCustomers)
	return &s
}

func (s *AdminServer) login(c *gin.Context) {
	var loginInfo = struct {
		Phone    string
		Password string
	}{}
	c.BindJSON(&loginInfo)
	user, err := admin.Login(loginInfo.Phone, loginInfo.Password)
	web.AssertNil(err)
	var auth = auth.Create(user.ID, "admin")
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *AdminServer) getCustomers(c *gin.Context) {
	var users, err = customer.GetCustomers()
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *AdminServer) generateCSV(c *gin.Context) {
	// var quantity, _ = strconv.Atoi(c.Query("quantity"))
	// var orderID = c.Query("order_id")
	// var order, err = order.GetOrderByID(orderID)
	// web.AssertNil(err)
	// record := []string{"Link sản phẩm", "Mã thẻ cào"}
	// b := &bytes.Buffer{}
	// wr := csv.NewWriter(b)

}
