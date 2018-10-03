package customer

import (
	"github.com/gin-gonic/gin"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/o/process"
)

func (s *CustomerServer) getProcesses(c *gin.Context) {
	var customerID = c.MustGet("user_id").(string)
	var process, err = process.GetByCustomerID(customerID)
	web.AssertNil(err)
	s.SendData(c, process)
}

func (s *CustomerServer) createProcess(c *gin.Context) {
	var customerID = c.MustGet("user_id").(string)
	var process *process.Process
	c.BindJSON(&process)
	process.CustomerID = customerID
	web.AssertNil(process.Create())
	s.SendData(c, process)
}
