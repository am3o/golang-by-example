package main

import (
	"fmt"
)

type Number interface {
	~int | int16 | int32 | int64
}

type Value[T Number] interface {
	Get() T
	Set(value T)
}

type Entity[T Number] struct {
	value T
}

func (e Entity[T]) Get() T {
	return e.value
}

func (e *Entity[T]) Set(value T) {
	e.value = value
}

type Statistic[T ~[]V, V Number] struct {
	values T
}

func (s *Statistic[T, V]) Add(entity V) {
	s.values = append(s.values, entity)
}

func (s *Statistic[T, V]) combine() V {
	var result V
	return result
}

func main() {
	container := []any{
		Entity[int]{
			value: 2,
		},
		Entity[int32]{
			value: 36,
		},
	}

	for _, value := range container {
		switch value.(type) {
		case Entity[int]:
			fmt.Println(value)
		case Entity[int32]:
			fmt.Println(value)
		default:
			fmt.Printf("%T\n", value)
		}
	}

	fmt.Println(container)
}
