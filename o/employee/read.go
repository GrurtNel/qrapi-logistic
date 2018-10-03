package employee

import (
	"gopkg.in/mgo.v2/bson"
)

func GetEmployees() ([]*Employee, error) {
	var employee []*Employee
	var err = employeeTable.FindAll(&employee)
	return employee, err
}

func DeleteEmployeeByID(id string) error {
	return employeeTable.DeleteByID(id)
}

func GetEmployeeByCustomer(customerID string) ([]*Employee, error) {
	var employees []*Employee
	if customerID == "" {
		return GetEmployees()
	}
	var err = employeeTable.FindWhere(bson.M{"customer_id": customerID}, &employees)
	return employees, err
}
