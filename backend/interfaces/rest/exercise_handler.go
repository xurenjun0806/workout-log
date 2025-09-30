package rest

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
)

type ExerciseUseCase interface {
	Fetch(ctx context.Context, limit int64) ([]exercise.Exercise, string, error)
	GetByID(ctx context.Context, id string) (exercise.Exercise, error)
	//Search(ctx context.Context, parm SearchParam)
	Save(context.Context, *exercise.Exercise) error
	Delete(ctx context.Context, id string) error
}

type ExerciseHandler struct {
	useCase ExerciseUseCase
}

func NewExerciseHandler(e *echo.Echo, useCase ExerciseUseCase) {
	handler := &ExerciseHandler{
		useCase: useCase,
	}
	e.POST("/exercises", handler.Save)
	//e.GET("/exercises", handler.Fetch)
	e.GET("/exercises/:id", handler.GetByID)
	//e.GET("/exercises/search", handler.Search)
	e.DELETE("/exercises/:id", handler.Delete)
}

func (e *ExerciseHandler) Save(ctx echo.Context) error {
	return nil
}

func (e *ExerciseHandler) GetByID(ctx echo.Context) error {
	return nil
}

func (e *ExerciseHandler) Delete(ctx echo.Context) error {
	return nil
}
