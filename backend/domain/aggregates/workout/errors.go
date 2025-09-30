package workout

import "errors"

var (
	ErrInvalidWorkoutSession = errors.New("invalid workout session")
	ErrInvalidWorkoutItem    = errors.New("invalid workout item")
	ErrInvalidSet            = errors.New("invalid workout set")
)
