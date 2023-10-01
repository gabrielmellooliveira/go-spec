package spec

// OrSpecification used to create a new specification that is the OR of two other specifications.
type OrSpecification[T any] struct {
	BaseSpecification[T]
	left  Specification[T]
	right Specification[T]
}

func Or[T any](left Specification[T], right Specification[T]) Specification[T] {
	return &OrSpecification[T]{left: left, right: right}
}

func (spec *OrSpecification[T]) IsSatisfiedBy(t T) bool {
	return spec.left.IsSatisfiedBy(t) || spec.right.IsSatisfiedBy(t)
}
