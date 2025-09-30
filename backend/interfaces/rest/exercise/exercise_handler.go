package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
	"github.com/xurenjun0806/workout-log/backend/interfaces/rest"
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

func (e *ExerciseHandler) Save(c echo.Context) error {
	exercise := new(exercise.Exercise)
	if err := c.Bind(exercise); err != nil {
		return err
	}

	ctx := c.Request().Context()
	err := e.useCase.Save(ctx, exercise)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, exercise)
}

func (e *ExerciseHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	exercise, err := e.useCase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, exercise)
}

func (e *ExerciseHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	err := e.useCase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case exercise.ErrInvalidExerciseID, exercise.ErrInvalidExerciseName:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
