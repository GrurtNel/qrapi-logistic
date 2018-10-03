package customer

import (
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/o/employee"
)

func (s *CustomerServer) getEmployees(c *gin.Context) {
	var customerID = c.MustGet("user_id").(string)
	var users, err = employee.GetEmployeeByCustomer(customerID)
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *CustomerServer) createEmployee(c *gin.Context) {
	var employee *employee.Employee
	c.BindJSON(&employee)
	web.AssertNil(employee.Create())
	s.SendData(c, employee)
}

func (s *CustomerServer) deleteEmployee(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(employee.DeleteEmployeeByID(id))
	s.Success(c)
}
