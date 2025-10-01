package exercise

import (
	"time"
)

type ExerciseID string

func (wi ExerciseID) HasId() bool {
	return wi != ""
}

// Root
type Exercise struct {
	ID          ExerciseID
	Name        string
	BodyPart    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewExercise(name string, bodyPart string, description string) (*Exercise, error) {
	if name == "" {
		return nil, ErrInvalidExerciseName
	}
	return &Exercise{
		Name:        name,
		BodyPart:    bodyPart,
		Description: description,
	}, nil
}
