package exercise

import (
	"context"

	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
)

type Repository interface {
	Create(ctx context.Context, e *exercise.Exercise) error
	Fetch(ctx context.Context, limit int64) ([]exercise.Exercise, error)
	GetByID(ctx context.Context, id exercise.ExerciseID) (exercise.Exercise, error)
	Delete(ctx context.Context, id exercise.ExerciseID) error
}

type UseCase struct {
	repository Repository
}

func (u *UseCase) Fetch(ctx context.Context, limit int64) ([]exercise.Exercise, error) {
	if limit == 0 {
		return make([]exercise.Exercise, 0), nil
	}

	exercises, err := u.repository.Fetch(ctx, limit)
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func (u *UseCase) GetByID(ctx context.Context, id string) (exercise.Exercise, error) {
	e, err := u.repository.GetByID(ctx, exercise.ExerciseID(id))
	if err != nil {
		return exercise.Exercise{}, err
	}
	return e, nil
}

func (u *UseCase) CreateExercise(ctx context.Context, name string, bodyPart string, description string) (exercise.Exercise, error) {
	e, err := exercise.NewExercise(name, bodyPart, description)
	if err != nil {
		return exercise.Exercise{}, err
	}

	err = u.repository.Create(ctx, e)
	if err != nil {
		return exercise.Exercise{}, err
	}

	return *e, nil
}

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if err := u.repository.Delete(ctx, exercise.ExerciseID(id)); err != nil {
		return err
	}
	return nil
}
