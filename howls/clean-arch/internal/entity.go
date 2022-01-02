package internal

import "strconv"

type ID int64

func ValidId(id string) error {
	_, err := strconv.ParseInt(id, 10, 64)
	return err
}
func FromString(id string) ID {
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	return ID(parseInt)
}
