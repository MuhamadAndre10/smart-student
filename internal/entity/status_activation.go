package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type StatusActive struct {
	ID        string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	Status    string    `gorm:"column:status_name;size:50;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*StatusActive) TableName() string {
	return "status_activation"
}

func (s *StatusActive) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
