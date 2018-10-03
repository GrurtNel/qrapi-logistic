package order_status

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi-logistic/x/logger"
	"qrapi-logistic/x/mongodb"
)

var orderStatusLog = logger.NewLogger("order_status")
var orderStatusTable = mongodb.NewTable("order_status", "ors")

type OrderStatus struct {
	mongodb.Model `bson:",inline"`
	OrderID       string `bson:"order_id" json:"order_id"`
	ProcessID     string `bson:"process_id" json:"process_id"`
	Quantity      int    `bson:"quantity" json:"quantity"`
}

func (orderStatus *OrderStatus) Create() error {
	var condition = bson.M{
		"process_id": orderStatus.ProcessID,
		"order_id":   orderStatus.OrderID,
	}
	var existOrderStatus *OrderStatus
	orderStatusTable.FindWhere(condition, &existOrderStatus)
	if existOrderStatus != nil {
		return orderStatusTable.Update(condition, bson.M{
			"$inc": bson.M{
				"number_of_scan": 1,
			},
		})
	}
	orderStatus.Quantity = 1
	return orderStatusTable.Insert(orderStatus)
}
