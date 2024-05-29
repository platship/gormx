package gormx

import (
	"regexp"
	"strings"

	"github.com/duke-git/lancet/v2/slice"

	"gorm.io/gorm"
)

func Where(w *WhereOne, table string) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
		w.Field = CamelToSnake(reg.ReplaceAllString(w.Field, ""))
		if w.Field == "" {
			return db
		}
		if table != "" {
			w.Field = table + "." + w.Field
		}
		switch w.Rule {
		case WhereRuleEq:
			return db.Where(w.Field+"=?", w.Val)
		case WhereRuleNEq:
			return db.Where(w.Field+"!=?", w.Val)
		case WhereRuleEqTrue:
			return db.Where(w.Field + "=true")
		case WhereRuleEqFalse:
			return db.Where(w.Field + "=false")
		case WhereRuleNull:
			return db.Where(w.Field + "=''")
		case WhereRuleNNull:
			return db.Where(w.Field + "!=''")
		case WhereRuleGt:
			return db.Where(w.Field+"> ?", w.Val)
		case WhereRuleGtE:
			return db.Where(w.Field+">= ?", w.Val)
		case WhereRuleLt:
			return db.Where(w.Field+"< ?", w.Val)
		case WhereRuleLtE:
			return db.Where(w.Field+"<= ?", w.Val)
		case WhereRuleIn:
			return db.Where(w.Field+" IN ?", strings.Split(w.Val, ","))
		case WhereRuleInInt:
			return db.Where(w.Field+" IN ?", slice.IntSlice(strings.Split(w.Val, ",")))
		case WhereRuleLikes:
			vals := strings.Split(w.Val, ",")
			for _, v := range vals {
				db.Or(w.Field+" LIKE ?", "%"+v+"%")
			}
			return db
		case WhereRuleNLikes:
			vals := strings.Split(w.Val, ",")
			for _, v := range vals {
				db.Or(w.Field+" NOT LIKE ?", "%"+v+"%")
			}
			return db
		case WhereRuleLike:
			return db.Where(w.Field+" LIKE ?", "%"+w.Val+"%")
		case WhereRuleLikeBf:
			return db.Where(w.Field+" LIKE ?", w.Val+"%")
		case WhereRuleLikeAf:
			return db.Where(w.Field+" LIKE ?", "%"+w.Val)
		case WhereRuleBtw:
			timeArr := strings.Split(w.Val, ",")
			return db.Where(w.Field+" BETWEEN ? AND ?", timeArr[0], timeArr[1])
		case WhereRuleNBtw:
			return db.Where(w.Field+" IN ?", slice.IntSlice(strings.Split(w.Val, ",")))
		case WhereRuleJArr:
			return db.Where("JSON_CONTAINS(`" + w.Field + "`,JSON_ARRAY(" + w.Val + "))")
		case WhereRuleJObj:
			nameVal := strings.Split(w.Val, ",")
			return db.Where("JSON_CONTAINS(`" + w.Field + "`,JSON_OBJECT('" + nameVal[0] + "','" + nameVal[1] + "'))")
		default:
			return db.Where(w.Field+" LIKE ?", "%"+w.Val+"%")
		}
	}
}
