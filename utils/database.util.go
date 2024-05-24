package utils

import "gorm.io/gorm"

func Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := page * limit
		return db.Offset(offset).Limit(limit)
	}
}
