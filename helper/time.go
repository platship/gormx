/*
 * @Author: Coller
 * @Date: 2022-05-18 15:47:37
 * @LastEditTime: 2024-01-02 15:54:45
 * @Desc: 时间解析
 */
package helper

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// MyTime 自定义时间
type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

// 将数据库中取出的数据，转换成时间
func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// 将写入数据库中数据还原类型
func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime, nil
}

// 将数据库中取出的数据，赋值给目标类型
func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
