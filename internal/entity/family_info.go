package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Family struct {
	ID          string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	FullName    string    `gorm:"column:full_name;size:100;not null"`
	BirthDate   time.Time `gorm:"column:birth_of_date;not null"`
	FullAddress string    `gorm:"column:full_addresses;size:255;not null"`
	Phone       string    `gorm:"column:phone;size:15;unique;not null"`
	Employment  string    `gorm:"column:employment"`
	StudentID   string
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Family) TableName() string {
	return "family"
}

func (s *Family) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
