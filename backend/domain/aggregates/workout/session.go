package workout

import (
	"time"

	"github.com/xurenjun0806/workout-log/backend/domain/seedwork"
)

type WorkoutSessionID struct {
	seedwork.ID
}

type WorkoutSession struct {
	ID        WorkoutSessionID `json:"id"`
	StartedAt time.Time        `json:"started_at" validate:"required"`
	EndedAt   *time.Time       `json:"ended_at"`
	Items     []WorkoutItem    `json:"items"`
	Tags      []Tag            `json:"tags"`
	Notes     string           `json:"notes"`
}

func NewWorkoutSession(startedAt time.Time, notes string, tags []Tag) (*WorkoutSession, error) {
	var ws = WorkoutSession{
		StartedAt: startedAt,
		Items:     []WorkoutItem{},
		Notes:     notes,
		Tags:      tags,
	}
	return &ws, nil
}

func (ws *WorkoutSession) AddItem(item WorkoutItem) (*WorkoutSession, error) {
	if !item.IsValid() {
		return ws, ErrInvalidWorkoutItem
	}

	ws.Items = append(ws.Items, item)
	return ws, nil
}
