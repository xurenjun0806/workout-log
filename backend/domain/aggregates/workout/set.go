package workout

type Set struct {
	SetNumber int     `json:"set_number"`
	Reps      int     `json:"reps"`
	WeightKg  float64 `json:"weight_kg"`
}

func NewSet(setNumber, reps int, weightKg float64) (*Set, error) {
	var s Set = Set{
		SetNumber: setNumber,
		Reps:      reps,
		WeightKg:  weightKg,
	}
	if !s.IsValid() {
		return nil, ErrInvalidSet
	}
	return &s, nil
}

func (s Set) IsValid() bool {
	if s.SetNumber <= 0 {
		return false
	}
	if s.Reps <= 0 {
		return false
	}
	// WeightKg can be zero (e.g., bodyweight exercises)
	return true
}
