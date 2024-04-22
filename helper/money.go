/*
 * @Author: Coller
 * @Date: 2023-01-30 11:50:59
 * @LastEditTime: 2024-04-21 17:11:20
 * @Desc: 货币处理
 */
package helper

import (
	"database/sql/driver"
)

type Moneys float64

func (t Moneys) Value() (driver.Value, error) {
	if t == 0 {
		return nil, nil
	}
	return t, nil
}

// func (t *Moneys) Scan(v interface{}) error {
// 	if value, ok := v.(string); ok {
// 		*t = Moneys(numberx.MoneyToFormatFloat(value))
// 		return nil
// 	}
// 	return nil
// }
