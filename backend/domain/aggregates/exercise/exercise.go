package exercise

import "github.com/xurenjun0806/workout-log/backend/domain/seedwork"

type ExerciseID struct {
	seedwork.ID
}

// Root
type Exercise struct {
	ID          ExerciseID
	Name        string `json:"name" validate:"required"`
	BodyPart    string `json:"body_part" validate:"required"` // TODO: 固定で選択式にしたいかも
	Description string `json:"description"`
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
