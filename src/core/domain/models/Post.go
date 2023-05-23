package models

import (
	"github.com/google/uuid"
	domain "github.com/kainguyen/go-scrapper/src/core/domain/common"
	"gorm.io/gorm"
)

type Post struct {
	domain.AuditableEntity

	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title      string
	PostImages []PostImage `gorm:"constraint:OnDelete:CASCADE"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = uuid.New()
	return
}

type PostImage struct {
	domain.AuditableEntity

	Id     uuid.UUID `gorm:"primaryKey;type:uuid;uuid_generate_v4()"`
	Url    string
	PostId uuid.UUID
}

func (pi *PostImage) BeforeCreate(tx *gorm.DB) (err error) {
	pi.Id = uuid.New()
	return
}
