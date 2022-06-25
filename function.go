package fn

type (
	Consumer[E any] func(element E)

	BiConsumer[X any, Y any] func(x X, y Y)

	Function[I any, O any] func(in I) O
)
