package mysql

import (
	"context"
	"database/sql"

	"github.com/xurenjun0806/workout-log/backend/domain/aggregates/exercise"
	"github.com/xurenjun0806/workout-log/backend/infrastructure/id"
	"github.com/xurenjun0806/workout-log/backend/infrastructure/repositories"
)

type ExerciseRepository struct {
	Conn *sql.DB
}

func NewExerciseRepository(conn *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{Conn: conn}
}

func (repo *ExerciseRepository) Create(ctx context.Context, e *exercise.Exercise) error {
	if e.ID.HasId() {
		return repositories.ErrExistsExercise
	}

	query := `INSERT exercise SET id=?, name=? , body_part=? , description=?, created_at=?, updated_at=?`
	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	newId := exercise.ExerciseID(id.NewID())
	_, err = stmt.ExecContext(ctx, newId, e.Name, e.BodyPart, e.Description, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return err
	}

	e.ID = newId
	return nil
}

func (repo *ExerciseRepository) Fetch(ctx context.Context, limit int64) ([]exercise.Exercise, error) {
	return make([]exercise.Exercise, 0), nil
}

func (repo *ExerciseRepository) GetByID(ctx context.Context, id exercise.ExerciseID) (exercise.Exercise, error) {
	return exercise.Exercise{}, nil
}

func (repo *ExerciseRepository) Delete(ctx context.Context, id exercise.ExerciseID) error {
	return nil
}
