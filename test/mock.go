package test

import "github.com/gabrielmellooliveira/go-spec"

type User struct {
	minAge int
	maxAge int
	role   string
}

// User age Specification

type UserAgeSpecification struct {
	Specification spec.Specification[User]
}

func NewUserAgeSpecification() spec.Specification[User] {
	userAgeSpecification := &UserAgeSpecification{}
	userAgeSpecification.Specification = spec.NewSpecification[User](userAgeSpecification.IsSatisfiedBy)

	return userAgeSpecification.Specification
}

func (spec *UserAgeSpecification) IsSatisfiedBy(entity User) bool {
	return entity.minAge == entity.maxAge
}

// User role Specification

type UserRoleSpecification struct {
	Specification spec.Specification[User]
}

func NewUserRoleSpecification() spec.Specification[User] {
	userRoleSpecification := &UserRoleSpecification{}
	userRoleSpecification.Specification = spec.NewSpecification[User](userRoleSpecification.IsSatisfiedBy)

	return userRoleSpecification.Specification
}

func (spec *UserRoleSpecification) IsSatisfiedBy(entity User) bool {
	return entity.role == "ADMIN"
}
