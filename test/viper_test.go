package test

import (
	"github.com/MuhamadAndre10/smartStudentApp/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewViper(t *testing.T) {
	// Arrange
	// handle panic
	defer func() {
		if r := recover(); r != nil {
			assert.NotNil(t, r)
		}
	}()

	// Act
	cfg := config.NewViper()

	// Assert
	assert.Equal(t, "smart-student-app", cfg.GetString("app.name"))
}
