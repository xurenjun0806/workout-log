package exercise

import "errors"

var (
	ErrInvalidExerciseID   = errors.New("invalid exercise ID")
	ErrInvalidExerciseName = errors.New("invalid exercise Name")
)
