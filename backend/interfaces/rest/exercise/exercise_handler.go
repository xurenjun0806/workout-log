package exercise

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
	"github.com/xurenjun0806/workout-log/backend/interfaces/rest"
)

type ExerciseUseCase interface {
	Fetch(c context.Context, limit int64) ([]exercise.Exercise, error)
	GetByID(c context.Context, id string) (exercise.Exercise, error)
	//Search(ctx context.Context, parm SearchParam)
	CreateExercise(c context.Context, name string, bodyPart string, description string) (exercise.Exercise, error)
	Delete(c context.Context, id string) error
}

type ExerciseHandler struct {
	useCase ExerciseUseCase
}

type createExerciseRequest struct {
	Name        string `json:"name" validate:"required"`
	BodyPart    string `json:"body_part" validate:"required"` // TODO: 固定で選択式にしたいかも
	Description string `json:"description"`
}

func NewExerciseHandler(e *echo.Echo, useCase ExerciseUseCase) {
	handler := &ExerciseHandler{
		useCase: useCase,
	}
	e.POST("/exercises", handler.Create)
	e.GET("/exercises", handler.Fetch)
	e.GET("/exercises/:id", handler.GetByID)
	//e.GET("/exercises/search", handler.Search)
	e.DELETE("/exercises/:id", handler.Delete)
}

func (e *ExerciseHandler) Create(c echo.Context) error {
	input := new(createExerciseRequest)
	if err := c.Bind(input); err != nil {
		return err
	}

	ctx := c.Request().Context()
	exercise, err := e.useCase.CreateExercise(ctx, input.Name, input.BodyPart, input.Description)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, exercise)
}

func (e *ExerciseHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()
	// TODO: 具体的なLimitに関する実装は後で
	exercises, err := e.useCase.Fetch(ctx, 100)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, exercises)
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
