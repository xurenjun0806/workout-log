package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExercise_Valid(t *testing.T) {
	name := "Squat"
	bodyPart := "Legs"
	description := "Compound lift"

	e, err := NewExercise(name, bodyPart, description)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	if assert.NotNil(t, e) {
		assert.Equal(t, name, e.Name)
		assert.Equal(t, bodyPart, e.BodyPart)
		assert.Equal(t, description, e.Description)
	}
}

func TestNewExercise_InvalidName(t *testing.T) {
	e, err := NewExercise("", "Legs", "")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidExerciseName, err)
	assert.Nil(t, e)
}
