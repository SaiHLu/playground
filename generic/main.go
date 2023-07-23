package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (c *CustomMap[K, V]) Insert(k K, v V) error {
	c.data[k] = v
	return nil
}

func NewCustomMapFunc[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func printAny[V any](val V) {
	fmt.Println(val)
}

func main() {
	newCustomMap := NewCustomMapFunc[string, int]()

	newCustomMap.Insert("Ko Sai", 1)
	newCustomMap.Insert("Ko Sai2", 2)

	fmt.Println("Data: ", newCustomMap.data)

	newCustomMap2 := NewCustomMapFunc[string, string]()

	newCustomMap2.Insert("Ko Sai", "Lol")
	newCustomMap2.Insert("Ko Sai2", "Cool")

	fmt.Println("Data: ", newCustomMap2.data)

	printAny[int](1)
	printAny[string]("One")
	printAny[float64](1.123)
}
