package banana

import "errors"

var (
	UserConflic    = errors.New("Người dùng đã tồn tại")
	SignUpFall     = errors.New("Đăng kí thất bại")
	UserNotFound   = errors.New("Người dùng không tồn tại")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại")
)
