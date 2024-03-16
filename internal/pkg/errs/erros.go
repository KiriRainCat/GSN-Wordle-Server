package errs

import (
	"errors"
	"fmt"
)

// 服务器内部错误 - "500"
var ErrServer = errors.New("500")

// 数据库字段冲突 - "使用此 {field} 字段的内容已存在"
func ErrUniqueConstraint(field string) error {
	return fmt.Errorf("使用此 %s 字段的内容已存在", field)
}
