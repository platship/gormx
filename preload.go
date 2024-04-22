package gormx

import (
	"gorm.io/gorm"
)

func Preload(val bool, ws []*WhereOne) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if val {
			var sts []ScopeType
			for _, w := range ws {
				sts = append(sts, Where(w, ""))
			}
			return db.Scopes(sts...)
		}
		return db
	}
}
