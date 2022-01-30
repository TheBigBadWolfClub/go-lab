package pkg

type Error string

func (d Error) Error() string {
	return string(d)
}

const (
	IndexNotFound    Error = "index not found"
	IndexOutOfBounds Error = "index out of bounds"
)
