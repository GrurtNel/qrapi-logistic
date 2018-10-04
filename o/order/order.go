package order

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi-logistic/x/logger"
	"qrapi-logistic/x/mongodb"
)

const (
	PENDING_STATE  = "pending"
	DELIVERY_STATE = "delivery"
	DONE_STATE     = "done"
)

var orderLog = logger.NewLogger("tbl_order")
var orderTable = mongodb.NewTable("order", "prd")

type Order struct {
	mongodb.Model `bson:",inline"`
	Name          string   `bson:"name" json:"name"`
	CustomerID    string   `bson:"customer_id" json:"customer_id"`
	Quantity      int      `bson:"quantity" json:"quantity"`
	Deadline      int64    `bson:"deadline" json:"deadline"`
	Processes     []string `bson:"processes" json:"processes"`
	// Resources     []*Resource `bson:"resources" json:"resources"`
}

type Resource struct {
	EmployeeID string   `bson:"employee_id" json:"employee_id"`
	ProcessID  []string `bson:"process_id" json:"process_id"`
}

func (order *Order) Create() error {
	return orderTable.Create(order)
}

func GetOrdersByCustomer(customerID string) ([]*Order, error) {
	var orders []*Order
	if customerID == "" {
		return GetOrders()
	}
	var err = orderTable.FindWhere(bson.M{"customer_id": customerID}, &orders)
	return orders, err
}

func GetOrders() ([]*Order, error) {
	var order []*Order
	var err = orderTable.FindAll(&order)
	return order, err
}

func DeliveryOrder(id string) error {
	return orderTable.UpdateId(id, bson.M{
		"$set": bson.M{
			"activated": true,
		},
	})
}

func GetOrderByID(id string) (*Order, error) {
	var order *Order
	var err = orderTable.FindID(id, &order)
	return order, err
}

func GetOrderByIDAndEmployee(id string) (interface{}, error) {
	var result = &struct {
		ID         string `bson:"_id" json:"id"`
		Name       string `bson:"name" json:"name"`
		Deadline   int64  `bson:"deadline" json:"deadline"`
		CustomerID int64  `bson:"customer_id" json:"customer_id"`
		Processes  []struct {
			ID   string `json:"id" bson:"_id"`
			Name string `json:"name" bson:"name"`
		} `json:"processes" bson:"processes"`
	}{}
	var err = orderTable.Pipe([]bson.M{
		bson.M{
			"$match": bson.M{"_id": id},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "process",
				"localField":   "processes",
				"foreignField": "_id",
				"as":           "processes",
			},
		},
	}).One(result)
	return result, err
}
