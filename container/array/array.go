package array

import (
	"github.com/fzdwx/fn"
)

type (
	// Array the element container
	Array[E any] struct {
		arr []E
	}
)

// New array
func New[E any]() *Array[E] {
	return &Array[E]{arr: []E{}}
}

// From new Array from arr
func From[E any](arr []E) *Array[E] {
	return &Array[E]{arr: arr}
}

// Len return the length of Array
func (a Array[E]) Len() int {
	return len(a.arr)
}

// Add element to Array
func (a *Array[E]) Add(element E) *Array[E] {
	a.arr = append(a.arr, element)
	return a
}

// Foreach Array with consumer
func (a Array[E]) Foreach(consumer fn.BiConsumer[int, E]) {
	for i, e := range a.arr {
		consumer(i, e)
	}
}

// Map 什么时候才能加方法的类型参数？ todo
func (a Array[E]) Map(mapping fn.Function[E, any]) *Array[any] {
	result := make([]any, len(a.arr))

	for i, e := range a.arr {
		result[i] = mapping(e)
	}

	return From[any](result)

}
