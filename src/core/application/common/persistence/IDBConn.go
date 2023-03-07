package persistence

import (
	"gorm.io/gorm"
)

type IDBConn interface {
	DBConn() (*gorm.DB, error)
}
