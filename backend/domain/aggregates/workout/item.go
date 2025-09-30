package workout

import (
	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
	"github.com/xurenjun0806/workout-log/backend/domain/seedwork"
)

type WorkoutItemID struct {
	seedwork.ID
}
type WorkoutItem struct {
	ID         WorkoutItemID       `json:"id"`
	ExerciseID exercise.ExerciseID `json:"exercise_id" validate:"required"`
	Position   int                 `json:"position" validate:"required"`
	Sets       []Set               `json:"sets"`
}

func NewWorkoutItem(exerciseID exercise.ExerciseID, position int) (*WorkoutItem, error) {
	if exerciseID.IsZero() {
		return nil, exercise.ErrInvalidExerciseID
	}
	var wi WorkoutItem = WorkoutItem{
		ExerciseID: exerciseID,
		Position:   position,
		Sets:       []Set{},
	}
	if !wi.IsValid() {
		return nil, ErrInvalidWorkoutItem
	}
	return &wi, nil
}

func (wi *WorkoutItem) IsNew() bool {
	return wi.ID.IsZero()
}

func (wi *WorkoutItem) IsValid() bool {
	if wi == nil {
		return false
	}
	if wi.ExerciseID.IsZero() {
		return false
	}
	return true
}

func (wi *WorkoutItem) AddSet(set Set) (*WorkoutItem, error) {
	if !wi.IsValid() {
		return wi, ErrInvalidWorkoutItem
	}
	if !set.IsValid() {
		return wi, ErrInvalidSet
	}
	wi.Sets = append(wi.Sets, set)
	return wi, nil
}
