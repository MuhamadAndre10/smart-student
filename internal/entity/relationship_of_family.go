package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Relationship struct {
	ID           string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	RelationName string    `gorm:"column:relation_name;unique;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Relationship) TableName() string {
	return "relationship"
}

func (s *Relationship) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
