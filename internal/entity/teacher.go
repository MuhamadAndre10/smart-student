package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID          string     `gorm:"column:id;primaryKey;size:100;unique;not null"`
	FullName    string     `gorm:"column:full_name;size:100;not null"`
	Email       string     `gorm:"column:email;size:100;unique"`
	Phone       string     `gorm:"column:phone;size:15;unique;not null"`
	Gender      string     `gorm:"column:gender;type:genderType;not null"`
	BirthDate   time.Time  `gorm:"column:birth_of_date;not null"`
	FullAddress string     `gorm:"column:full_addresses;size:255;not null"`
	Photo       string     `gorm:"column:photo;size:255"`
	Skill       string     `gorm:"column:skill"`
	Schedule    []Schedule `gorm:"foreignKey:TeacherID;references:ID"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Teacher) TableName() string {
	return "teacher"
}

func (s *Teacher) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
