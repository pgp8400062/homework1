package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// 查询数据库，模拟出错
func selectDB() error {
	return errors.Wrap(sql.ErrNoRows, "select data from db failed")
}

// 模拟下层调用无法处理
func xxService() error {
	return errors.WithMessage(selectDB(), "no data")
}

// 再下一层业务处理，下一步进阶错误统一处理
func main() {
	err := xxService()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("抛出业务错误码：100001")
		return
	}
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
