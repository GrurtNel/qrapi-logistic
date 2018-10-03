package order

import (
	"qrapi-logistic/g/x/web"
)

func DeleteByID(id string) error {
	var order, err = GetOrderByID(id)
	if err != nil || order == nil {
		return web.BadRequest("Không tồn tại đơn hàng")
	}
	return orderTable.DeleteByID(id)
}
