package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID          string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	FullName    string    `gorm:"column:full_name"`
	Email       string    `gorm:"column:email;unique;not null"`
	Phone       int       `gorm:"column:no_phone;size:15;unique;not null"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
	Gender      string    `gorm:"column:gender;type:genderType;not null"`
	Photo       string    `gorm:"column:photo;size:255"`
	Password    string    `gorm:"column:password;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Admin) TableName() string {
	return "admins"
}

func (s *Admin) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
