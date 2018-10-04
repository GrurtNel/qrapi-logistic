package order_history

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi-logistic/x/logger"
	"qrapi-logistic/x/mongodb"
)

var orderHistoryLog = logger.NewLogger("order_history")
var orderHistoryTable = mongodb.NewTable("order_history", "orh")

type OrderHistory struct {
	mongodb.Model `bson:",inline"`
	OrderID       string `bson:"order_id" json:"order_id"`
	ProductID     string `bson:"product_id" json:"product_id"`
	EmployeeID    string `bson:"employee_id" json:"employee_id"`
	EmployeeName  string `bson:"employee_name" json:"employee_name"`
	ProcessID     string `bson:"process_id" json:"process_id"`
	ProcessName   string `bson:"process_name" json:"process_name"`
}

func (orderHistory *OrderHistory) Create() error {
	return orderHistoryTable.Create(orderHistory)
}

func GetHistoryByProductID(orderID, productID string) ([]*OrderHistory, error) {
	var result []*OrderHistory
	var query = bson.M{}
	if orderID != "" {
		query["order_id"] = orderID
	}
	if productID != "" {
		query["product_id"] = productID
	}
	var err = orderHistoryTable.FindWhere(query, &result)
	return result, err
}
