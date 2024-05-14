package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID              string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	Title           string    `gorm:"column:title;size:100;not null"`
	Author          string    `gorm:"column:author;size:100;not null"`
	Publisher       string    `gorm:"column:publisher"`
	PublicationDate time.Time `gorm:"column:publication_date"`
	Edition         string    `gorm:"column:edition"`
	Price           int       `gorm:"column:price"`
	Description     string    `gorm:"column:description"`
	StatusKhatam    string    `gorm:"column:status_khatam"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Book) TableName() string {
	return "books"
}

func (s *Book) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
