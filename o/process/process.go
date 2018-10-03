package process

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi-logistic/x/logger"
	"qrapi-logistic/x/mongodb"
)

type Process struct {
	mongodb.Model `bson:",inline"`
	Name          string `bson:"name" json:"name"`
	CustomerID    string `bson:"customer_id" json:"customer_id"`
}

var processTable = mongodb.NewTable("process", "prc")
var processLog = logger.NewLogger("process")

func (p *Process) Create() error {
	return processTable.Create(p)
}

func GetProcesses() ([]*Process, error) {
	var processes []*Process
	var err = processTable.FindAll(&processes)
	return processes, err
}

func GetByCustomerID(customer string) ([]*Process, error) {
	var processes []*Process
	if customer == "" {
		return GetProcesses()
	}
	var err = processTable.FindWhere(bson.M{"customer_id": customer}, &processes)
	return processes, err
}
