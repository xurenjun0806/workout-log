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

func (m *ExerciseRepository) fetch(ctx context.Context, query string, args ...any) ([]exercise.Exercise, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
		}
	}()

	result := make([]exercise.Exercise, 0)
	for rows.Next() {
		t := exercise.Exercise{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.BodyPart,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
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
	query := `SELECT id, name, body_part, description, created_at, updated_at FROM exercise WHERE ID = ?`
	list, err := repo.fetch(ctx, query, id)
	if err != nil {
		return exercise.Exercise{}, err
	}

	if len(list) > 0 {
		return list[0], nil
	} else {
		return exercise.Exercise{}, exercise.ErrNotFound
	}
}

func (repo *ExerciseRepository) Delete(ctx context.Context, id exercise.ExerciseID) error {
	return nil
}
