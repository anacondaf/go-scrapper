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
	PostImages []PostImages
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = uuid.New()
	return
}

type PostImages struct {
	domain.AuditableEntity

	Id     uuid.UUID `gorm:"primaryKey;type:uuid;uuid_generate_v4()"`
	Url    string
	PostId uuid.UUID
}

func (pi *PostImages) BeforeCreate(tx *gorm.DB) (err error) {
	pi.Id = uuid.New()
	return
}
