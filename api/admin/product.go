package admin

import (
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/o/product"
)

func (s *AdminServer) createProduct(c *gin.Context) {
	var product *product.Product
	c.BindJSON(&product)
	web.AssertNil(product.Create())
	s.SendData(c, product)
}

func (s *AdminServer) updateProduct(c *gin.Context) {
	var product *product.Product
	c.BindJSON(&product)
	web.AssertNil(product.Create())
	s.SendData(c, product)
}

func (s *AdminServer) deleteProduct(c *gin.Context) {
	web.AssertNil(product.DeleteByID(c.Query("id")))
	s.Success(c)
}

func (s *AdminServer) getProducts(c *gin.Context) {
	var customerID = c.Query("customer_id")
	var products, err = product.GetProductsByCustomer(customerID)
	web.AssertNil(err)
	s.SendData(c, products)
}
