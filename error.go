package routegopherjs

import (
	"gopkg.in/go-on/method.v1"
	"gopkg.in/go-on/router.v2/route"
)

// ErrServiceAlreadyRegistered is raised if the Service has already been registered.
type ErrServiceAlreadyRegistered struct{}

func (ErrServiceAlreadyRegistered) Error() string {
	return " handler already registered"
}

// ErrServiceNotRegistered is raised if no Service has been registered.
type ErrServiceNotRegistered struct{}

func (ErrServiceNotRegistered) Error() string {
	return " handler not registered"
}

// ErrMethodNotDefined is raised if the given http method is not defined for the given route
type ErrMethodNotDefined struct {
	method.Method
	Route *route.Route
}

func (e *ErrMethodNotDefined) Error() string {
	return "method " + e.Method.String() + " is not defined for route " + e.Route.DefinitionPath
}
