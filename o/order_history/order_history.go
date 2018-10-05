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

func GetHistoryByID(orderID, productID string) ([]*OrderHistory, error) {
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

func GetTrackingByID(orderID string) (interface{}, error) {
	var result = []*struct {
		ProcessID   string `bson:"_id" json:"id"`
		ProcessName string `bson:"process_name" json:"process_name"`
		Count       int    `bson:"count" json:"count"`
	}{}
	var err = orderHistoryTable.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{"order_id": orderID},
		},
		bson.M{
			"$group": bson.M{
				"_id":          "$process_id",
				"process_name": bson.M{"$first": "$process_name"},
				"count":        bson.M{"$sum": 1},
			},
		},
	}).All(&result)
	if err != nil {
		orderHistoryLog.Error(err)
	}
	return result, err
}
