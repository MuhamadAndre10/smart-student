package test

import (
	"github.com/MuhamadAndre10/smartStudentApp/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TestNewDatabase failed: %v", r)
		}
	}()

	// Arrange
	viper := config.NewViper()
	logger := config.NewLog()

	// Act
	db := config.NewDatabase(viper, logger)

	// Assert
	assert.NotNil(t, db)
}
