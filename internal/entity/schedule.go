package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	ID          string    `gorm:"column:id;primaryKey;size:100;unique;not null"`
	StudyName   string    `gorm:"column:study_name"`
	StartTime   time.Time `gorm:"column:start_time;not null"`
	EndTime     time.Time `gorm:"column:end_time;not null"`
	DateTime    time.Time `gorm:"column:date_time;not null"`
	Description string    `gorm:"column:description"`
	ClassID     string
	TeacherID   string
	BookID      string
	Book        Book      `gorm:"foreignKey:BookID;references:ID"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (*Schedule) TableName() string {
	return "schedules"
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid.String()

	return nil
}
