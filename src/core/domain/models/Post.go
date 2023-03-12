package models

import (
	"github.com/google/uuid"
	domain "github.com/kainguyen/go-scrapper/src/core/domain/common"
)

type Post struct {
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title      string
	PostImages []PostImages
	domain.AuditableEntity
}

type PostImages struct {
	Id     uint `gorm:"primaryKey,autoIncrement"`
	Url    string
	PostId uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	domain.AuditableEntity
}
