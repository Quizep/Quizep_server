package reposistories

import "gorm.io/gorm"

func NewRepositories(dns string) *gorm.DB {
	sql, err := NewMySQL(dns)
	if err != nil {
		panic(err)
	}
	return sql
}
