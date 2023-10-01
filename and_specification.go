package spec

// AndSpecification used to create a new specification that is the AND of two other specifications.
type AndSpecification[T any] struct {
	BaseSpecification[T]
	left  Specification[T]
	right Specification[T]
}

func And[T any](left Specification[T], right Specification[T]) Specification[T] {
	return &AndSpecification[T]{left: left, right: right}
}

func (spec *AndSpecification[T]) IsSatisfiedBy(entity T) bool {
	return spec.left.IsSatisfiedBy(entity) && spec.right.IsSatisfiedBy(entity)
}
