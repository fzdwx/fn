package array

import (
	"fmt"
	"testing"
)

func Test_Foreach(t *testing.T) {
	a := New[int]()

	a.Add(1)
	a.Add(2)
	a.Add(3)

	a.Foreach(func(i int, e int) {
		fmt.Println("idx:", i, " - val:", e)
	})
}

func TestArray_Map(t *testing.T) {
	a := New[int]()

	a.Add(1)
	a.Add(2)
	a.Add(3)

	anies := a.Map(func(e int) any {
		return 1
	})

	anies.Foreach(func(i int, e any) {
		fmt.Println(e)
	})
}
