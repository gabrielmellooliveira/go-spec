package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecification(t *testing.T) {
	userAgeSpecification := NewUserAgeSpecification()
	userRoleSpecification := NewUserRoleSpecification()

	user := User{minAge: 1, maxAge: 1, role: "ADMIN"}

	assert.True(t, userAgeSpecification.IsSatisfiedBy(user))
	assert.True(t, userAgeSpecification.And(userRoleSpecification).IsSatisfiedBy(user))

	user = User{minAge: 1, maxAge: 2, role: "USER"}

	assert.False(t, userAgeSpecification.IsSatisfiedBy(user))
	assert.False(t, userAgeSpecification.And(userRoleSpecification).IsSatisfiedBy(user))
	assert.True(t, userAgeSpecification.Not(userRoleSpecification).IsSatisfiedBy(user))

	user = User{minAge: 1, maxAge: 1, role: "USER"}

	assert.True(t, userAgeSpecification.Or(userRoleSpecification).IsSatisfiedBy(user))
}
