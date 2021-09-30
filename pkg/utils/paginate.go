package utils

import "gorm.io/gorm"

// 分页
func Paginate(p, s int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p == 0 {
			p = 1
		}
		switch {
		case s > 100:
			s = 100
		case s <= 0:
			s = 10
		}
		return db.Offset((p - 1) * s).Limit(s)
	}
}
