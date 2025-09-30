package exercise

import (
	"time"

	"github.com/xurenjun0806/workout-log/backend/domain/seedwork"
)

type ExerciseID struct {
	seedwork.ID
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
