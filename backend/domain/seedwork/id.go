package seedwork

type ID string

func (id *ID) IsZero() bool {
	return id == nil || *id == ""
}
