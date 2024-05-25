package entity

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type genderType string

const (
	Men   genderType = "LAKI-LAKI"
	Women genderType = "PEREMPUAN"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}
func (gt *genderType) Value() (driver.Value, error) {
	return string(*gt), nil
}

type Student struct {
	ID                 string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	FullName           string    `gorm:"column:full_name;size:100;not null"`
	BirthDate          time.Time `gorm:"column:birth_of_date;not null"`
	FullAddress        string    `gorm:"column:full_addresses;size:255;not null"`
	Email              string    `gorm:"column:email;size:100;unique"`
	Phone              string    `gorm:"column:phone;size:15;unique;not null"`
	Gender             string    `gorm:"column:gender;type:gender_type;not null"`
	Photo              string    `gorm:"column:photo;size:255"`
	StatusActivationID string
	StatusActive       StatusActive `gorm:"foreignKey:StatusActivationID;references:ID"`
	Family             []Family     `gorm:"foreignKey:StudentID;references:ID"`
	ClassID            string
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Student) TableName() string {
	return "students"
}

func (s *Student) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
