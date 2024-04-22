/*
 * @Author: Coller
 * @Date: 2022-05-18 15:47:37
 * @LastEditTime: 2024-04-21 16:20:20
 * @Desc: 自定义日期，接收字符串，存日期
 */
package helper

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Date string

// 写入数据库转换类型
func (t Date) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	if len(t) != 10 {
		return time.Now, errors.New("输入的日期数据不符合规范")
	}
	times, _ := time.Parse("2006-01-02", string(t))
	return times, nil
}

// 取出做类型转换
func (t *Date) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		tTime := time.Time(vt).Format("2006-01-02")
		*t = Date(tTime)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}
