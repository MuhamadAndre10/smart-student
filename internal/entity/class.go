package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID                 string `gorm:"column:id;primaryKey;unique;not null"`
	ClassName          string `gorm:"column:class_name;not null"`
	AcademicYear       string `gorm:"column:academic_year;not null"`
	ClassroomInfo      string `gorm:"column:classroom_info"`
	NumberOfStudent    string `gorm:"column:number_of_student;not null"`
	Level              string `gorm:"column:level;not null"`
	Description        string `gorm:"column:description"`
	StatusActivationID string
	StatusActive       StatusActive `gorm:"foreignKey:StatusActivationID;references:ID"`
	TeacherID          string
	Teacher            Teacher    `gorm:"foreignKey:TeacherID;references:ID"`
	Student            []Student  `gorm:"foreignKey:ClassID;references:ID"`
	Schedule           []Schedule `gorm:"foreignKey:ClassID;references:ID"`
	CreatedAt          string     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt          string     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Class) TableName() string {
	return "classes"
}

func (s *Class) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
