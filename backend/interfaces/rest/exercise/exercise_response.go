package exercise

import (
	"time"

	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
)

type ExerciseResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	BodyPart    string    `json:"body_part"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func toExerciseResponse(e exercise.Exercise) ExerciseResponse {
	return ExerciseResponse{
		ID:          string(e.ID),
		Name:        e.Name,
		BodyPart:    e.BodyPart,
		Description: e.Description,
	}
}
