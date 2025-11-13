package main

type Equalizer[T any] interface {
	Equal(other T) bool
}

type ID int

func (id ID) Equal(other ID) bool {
	return id == other
}

func HasMyInterface[T Equalizer[T]](list []T, value T) bool {
	for _, v := range list {
		if value.Equal(v) {
			return true
		}
	}

	return false
}
