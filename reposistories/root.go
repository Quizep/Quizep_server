package reposistories

import "gorm.io/gorm"

func NewReposistories(dns string) (*gorm.DB, error) {
	sql, err := NewMySQL(dns)

	return sql, err
}
