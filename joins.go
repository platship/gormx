package gormx

import (
	"gorm.io/gorm"
)

func Joins(query string, db *gorm.DB, val bool, wss []*WhereOne) *gorm.DB {
	if val {
		var sts []ScopeType
		for _, w := range wss {
			sts = append(sts, Where(w, query))
		}
		return db.Scopes(sts...)
	}
	return db

}
