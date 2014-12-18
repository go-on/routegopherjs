package routegopherjs

import (
	"github.com/gopherjs/gopherjs/js"
	"gopkg.in/go-on/method.v1"
	"gopkg.in/go-on/router.v2/route"
)

// Service is and interface that may be fulfilled by client side libraries for gopherjs.
// If should be passed once to RegisterService to allow common kind of requests via the *route.Get, ...
// methods.
type Service interface {
	Get(url string, callback func(js.Object))
	Post(url string, data interface{}, callback func(js.Object))
	Put(url string, data interface{}, callback func(js.Object))
	Patch(url string, data interface{}, callback func(js.Object))
	Delete(url string, callback func(js.Object))
	Options(url string, callback func(js.Object))
}

// RegisterService allows a central registration for client side  libraries.
// It may only be called once and must be called before using any of the *route.Get etc methods.
func RegisterService(aj Service) {
	if xhr != nil {
		panic(ErrServiceAlreadyRegistered{})
	}
	xhr = aj
}

// Get provides a shortcut for a GET request via the centralized  service.
// RegisterService must have been called refore.
func Get(r *route.Route, callback func(js.Object), params ...string) {
	xhrMustBeRegistered()
	if !r.HasMethod(method.GET) {
		panic(&ErrMethodNotDefined{method.GET, r})
	}
	xhr.Get(r.MustURL(params...), callback)
}

// Delete provides a shortcut for a DELETE request via the centralized  service.
// RegisterService must have been called refore.
func Delete(r *route.Route, callback func(js.Object), params ...string) {
	xhrMustBeRegistered()
	if !r.HasMethod(method.DELETE) {
		panic(&ErrMethodNotDefined{method.DELETE, r})
	}
	xhr.Delete(r.MustURL(params...), callback)
}

// Post provides a shortcut for a POST request via the centralized  service.
// RegisterService must have been called refore.
func Post(r *route.Route, data interface{}, callback func(js.Object), params ...string) {
	xhrMustBeRegistered()
	if !r.HasMethod(method.POST) {
		panic(&ErrMethodNotDefined{method.POST, r})
	}
	xhr.Post(r.MustURL(params...), data, callback)
}

// Patch provides a shortcut for a PATCH request via the centralized  service.
// RegisterService must have been called refore.
func Patch(r *route.Route, data interface{}, callback func(js.Object), params ...string) {
	xhrMustBeRegistered()
	if !r.HasMethod(method.PATCH) {
		panic(&ErrMethodNotDefined{method.PATCH, r})
	}
	xhr.Patch(r.MustURL(params...), data, callback)
}

// Put provides a shortcut for a PUT request via the centralized  service.
// RegisterService must have been called refore.
func Put(r *route.Route, data interface{}, callback func(js.Object), params ...string) {
	xhrMustBeRegistered()
	if !r.HasMethod(method.PUT) {
		panic(&ErrMethodNotDefined{method.PUT, r})
	}
	xhr.Put(r.MustURL(params...), data, callback)
}

// Options provides a shortcut for a OPTIONS request via the centralized  service.
// RegisterService must have been called refore.
func Options(r *route.Route, callback func(js.Object), params ...string) {
	xhrMustBeRegistered()
	if !r.HasMethod(method.OPTIONS) {
		panic(&ErrMethodNotDefined{method.OPTIONS, r})
	}
	xhr.Options(r.MustURL(params...), callback)
}

// Funcs offers an easy way to create an adapter to fulfill the Service interface
// by providing a set of functions
type Funcs struct {
	GET     func(url string, callback func(js.Object))
	POST    func(url string, body interface{}, callback func(js.Object))
	PUT     func(url string, body interface{}, callback func(js.Object))
	PATCH   func(url string, body interface{}, callback func(js.Object))
	DELETE  func(url string, callback func(js.Object))
	OPTIONS func(url string, callback func(js.Object))
}

// Get implements the Service.
func (ps *Funcs) Get(url string, callback func(js.Object)) {
	ps.GET(url, callback)
}

// Post implements the Service.
func (ps *Funcs) Post(url string, data interface{}, callback func(js.Object)) {
	ps.POST(url, data, callback)
}

// Put implements the Service.
func (ps *Funcs) Put(url string, data interface{}, callback func(js.Object)) {
	ps.PUT(url, data, callback)
}

// Patch implements the Service.
func (ps *Funcs) Patch(url string, data interface{}, callback func(js.Object)) {
	ps.PATCH(url, data, callback)
}

// Delete implements the Service.
func (ps *Funcs) Delete(url string, callback func(js.Object)) {
	ps.DELETE(url, callback)
}

// Options implements the Service.
func (ps *Funcs) Options(url string, callback func(js.Object)) {
	ps.OPTIONS(url, callback)
}

var xhr Service = nil

func xhrMustBeRegistered() {
	if xhr == nil {
		panic(ErrServiceNotRegistered{})
	}
}
