package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	//Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	gorm.Model
	Title      string
	PostImages []PostImages
	//domain.AuditableEntity
}

//func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
//	p.Id = uuid.New()
//	return
//}

type PostImages struct {
	//Id     uint `gorm:"primaryKey,autoIncrement"`
	gorm.Model
	Url    string
	PostId uuid.UUID
	//domain.AuditableEntity
}
