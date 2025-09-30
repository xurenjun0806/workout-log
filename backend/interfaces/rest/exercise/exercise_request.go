package exercise

type CreateExerciseRequest struct {
	Name        string `json:"name" validate:"required"`
	BodyPart    string `json:"body_part" validate:"required"`
	Description string `json:"description"`
}
