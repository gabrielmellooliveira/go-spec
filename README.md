# go-spec

[![GoDoc Widget]][GoDoc]

`go-spec` is a lib to implement Specification Pattern.

Specification is to separate the statement of how to match a candidate, from the candidate object that it is matched against. As well as its usefulness in selection, it is also valuable for validation and for building to order.

## Install

`go get -u github.com/gabrielmellooliveira/go-spec`

## Examples

### Create a specification

```go
package main

import "github.com/gabrielmellooliveira/go-spec"

type User struct {
	minAge int
	maxAge int
}

type UserAgeSpecification struct {
	Specification spec.Specification[User]
}

func NewUserAgeSpecification() spec.Specification[User] {
	specification := &UserAgeSpecification{}
	specification.Specification = spec.NewSpecification[User](specification.IsSatisfiedBy)

	return specification.Specification
}

func (spec *UserAgeSpecification) IsSatisfiedBy(entity User) bool {
	return entity.minAge == entity.maxAge
}
```

### Use a specification

```go
package main

import "github.com/gabrielmellooliveira/go-spec"

func main() {
    userAgeSpecification := NewUserAgeSpecification()

    user := User{minAge: 18, maxAge: 18}

    if userAgeSpecification.IsSatisfiedBy(user) {
        ...
    } else {
        ...
    }
}
```

### Combine specifications

```go
package main

import "github.com/gabrielmellooliveira/go-spec"

func main() {
    userAgeSpecification := NewUserAgeSpecification()
    userRoleSpecification := NewUserRoleSpecification()

    user := User{minAge: 18, maxAge: 18}

    if userAgeSpecification.And(userRoleSpecification).IsSatisfiedBy(user) {
        ...
    } else {
        ...
    }
}
```

## License

Copyright (c) 2015-present [Gabriel Mello de Oliveira](https://github.com/gabrielmellooliveira)

Licensed under [MIT License](./LICENSE)

[GoDoc]: https://pkg.go.dev/github.com/gabrielmellooliveira/go-spec
[GoDoc Widget]: https://godoc.org/github.com/gabrielmellooliveira/go-spec?status.svg
