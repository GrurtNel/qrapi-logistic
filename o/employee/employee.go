package employee

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi-logistic/g/x/web"
	"qrapi-logistic/x/logger"
	"qrapi-logistic/x/mongodb"
)

type Employee struct {
	mongodb.Model  `bson:",inline"`
	Name           string   `bson:"name" json:"name"`
	Uname          string   `bson:"uname" json:"uname"`
	HashedPassword string   `bson:"password" json:"-"`
	Password       Password `bson:"-" json:"password"`
	CustomerID     string   `bson:"customer_id" json:"customer_id"`
}

var employeeTable = mongodb.NewTable("employee", "emp")
var employeeLog = logger.NewLogger("employee")

var (
	errExistUname = "Xin lỗi! Người dùng đã tồn tại trong hệ thống"
)

func GetEmployeeByUname(uname string) (*Employee, error) {
	var employee *Employee
	var err = employeeTable.FindOne(bson.M{
		"uname": uname,
	}, &employee)
	return employee, err
}

func (u *Employee) Create() error {
	hashed, _ := u.Password.Hash()
	u.HashedPassword = hashed
	var existEmail, _ = GetEmployeeByUname(u.Uname)
	if existEmail != nil {
		return web.BadRequest(errExistUname)
	}
	return employeeTable.Create(u)
}

func Login(uname, pwd string) (*Employee, error) {
	var employee *Employee
	var query = bson.M{"uname": uname}
	err := employeeTable.FindOne(query, &employee)
	if err != nil {
		if err.Error() == "not found" {
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	if err := Password(pwd).ComparePassword(employee.HashedPassword); err != nil {
		if err.Error() == ErrMismatchedHashAndPassword {
			employeeLog.Error(err)
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	// if !customer.Activated {
	// 	return nil, web.BadRequest("Tài khoản chưa được kích hoạt vui lòng liên hệ để được kích hoạt")
	// }
	return employee, nil
}
