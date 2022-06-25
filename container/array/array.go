package array

type (
	Array struct {
		arr []any
	}
)

// New array
func New() *Array {
	return &Array{arr: []any{}}
}
