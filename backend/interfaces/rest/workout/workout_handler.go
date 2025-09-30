package workout

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/workout"
	"github.com/xurenjun0806/workout-log/backend/interfaces/rest"
)

type WorkoutUseCase interface {
	Fetch(ctx context.Context, limit int64) ([]workout.WorkoutSession, string, error)
	GetByID(ctx context.Context, id string) (workout.WorkoutSession, error)
	Save(context.Context, *workout.WorkoutSession) error
	Delete(ctx context.Context, id string) error
}

type WorkoutHandler struct {
	useCase WorkoutUseCase
}

func NewWorkoutHandler(e *echo.Echo, useCase WorkoutUseCase) {
	handler := &WorkoutHandler{useCase: useCase}
	e.POST("/workout", handler.Save)
	//e.GET("/workout", handler.Fetch)
	e.GET("/workout/:id", handler.GetById)
	e.DELETE("/workout/:id", handler.Delete)
}

func (w *WorkoutHandler) Save(c echo.Context) error {
	ws := new(workout.WorkoutSession)
	if err := c.Bind(ws); err != nil {
		return err
	}

	ctx := c.Request().Context()
	err := w.useCase.Save(ctx, ws)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ws)
}

func (w *WorkoutHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	ws, err := w.useCase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ws)
}

func (w *WorkoutHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	err := w.useCase.Delete(ctx, id)
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
	case workout.ErrInvalidWorkoutSession, workout.ErrInvalidWorkoutItem, workout.ErrInvalidSet:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
