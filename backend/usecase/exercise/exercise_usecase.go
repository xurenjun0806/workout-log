package exercise

import (
	"context"

	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
)

type UseCase struct{}

func (u *UseCase) Fetch(c context.Context, limit int64) ([]exercise.Exercise, error) {
	return make([]exercise.Exercise, 0), nil
}

func (u *UseCase) GetByID(c context.Context, id string) (exercise.Exercise, error) {
	return exercise.Exercise{
		ID:          exercise.ExerciseID{ID: "test"},
		Name:        "testName",
		BodyPart:    "脚",
		Description: "ダミーデータ",
	}, nil
}

func (u *UseCase) Save(c context.Context, exercise *exercise.Exercise) error {
	return nil
}
func (u *UseCase) Delete(c context.Context, id string) error {
	return nil
}
