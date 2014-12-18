package routegopherjs

import (
	"fmt"
	"gopkg.in/go-on/router.v2/route"
	"reflect"
	"testing"

	"github.com/gopherjs/gopherjs/js"
	"gopkg.in/go-on/method.v1"
)

func errorMustBe(err interface{}, class interface{}) string {
	classTy := reflect.TypeOf(class)
	if err == nil {
		return fmt.Sprintf("error must be of type %s but is nil", classTy)
	}

	errTy := reflect.TypeOf(err)
	if errTy.String() != classTy.String() {
		return fmt.Sprintf("error must be of type %s but is of type %s", classTy, errTy)
	}
	return ""
}

func TestDoubleRegisteredService(t *testing.T) {
	xhr = nil
	aj := &Funcs{}
	RegisterService(aj)

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, ErrServiceAlreadyRegistered{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		_ = e.(ErrServiceAlreadyRegistered).Error()
	}()

	RegisterService(aj)
}

func TestNotRegisteredService(t *testing.T) {
	xhr = nil
	rt := route.New("/", method.GET)
	route.Mount("/", rt)
	defer func() {
		e := recover()
		errMsg := errorMustBe(e, ErrServiceNotRegistered{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		_ = e.(ErrServiceNotRegistered).Error()
	}()

	Get(rt, func(js.Object) {})
}

func TestUnknownMethod(t *testing.T) {

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, route.ErrUnknownMethod{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(route.ErrUnknownMethod)
		_ = err.Error()

		if err.Method.String() != "unknown" {
			t.Errorf("wrong method: %#v, expected: %v", err.Method, "unknown")
		}
	}()

	route.New("/route", method.Method("unknown"))
}

func TestErrPairParams(t *testing.T) {
	rt := route.New("/route", method.GET)

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, route.ErrPairParams{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(route.ErrPairParams)
		_ = err.Error()
	}()

	rt.MustURL("param1")
}

func TestErrMissingParams(t *testing.T) {
	rt := route.New("/route/:name", method.GET)

	route.Mount("/a", rt)

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, route.ErrMissingParam{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(route.ErrMissingParam)
		_ = err.Error()

		if err.Param != "name" {
			t.Errorf("wrong param: %#v, expected: %v", err.Param, "name")
		}

		if err.MountedPath != "/a/route/:name" {
			t.Errorf("wrong mountedPath: %#v, expected: %v", err.MountedPath, "/a/route/:name")
		}
	}()

	rt.MustURL()
}

func TestDoubleMounted(t *testing.T) {
	rt := route.New("/route/:name", method.GET)

	route.Mount("/a", rt)

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, &route.ErrDoubleMounted{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(*route.ErrDoubleMounted)
		_ = err.Error()

		if err.Path != "/a" {
			t.Errorf("wrong Path: %#v, expected: %v", err.Path, "/a")
		}

		if err.Route != rt {
			t.Errorf("wrong route: %#v, expected: %v", err.Route.DefinitionPath, rt.DefinitionPath)
		}
	}()

	route.Mount("/b", rt)
}

func testMethodNotDefined(has method.Method, hasNot method.Method, t *testing.T) {
	xhr = nil
	rt := route.New("/route/:name", has)

	route.Mount("/a", rt)

	x := &Funcs{}
	RegisterService(x)

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, &ErrMethodNotDefined{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(*ErrMethodNotDefined)
		_ = err.Error()

		if err.Method != hasNot {
			t.Errorf("wrong method: %#v, expected: %v", err.Method.String(), hasNot.String())
		}

		if err.Route != rt {
			t.Errorf("wrong route: %#v, expected: %v", err.Route.DefinitionPath, rt.DefinitionPath)
		}
	}()

	switch hasNot {
	case method.GET:
		Get(rt, nil)
	case method.POST:
		Post(rt, nil, nil)
	case method.PUT:
		Put(rt, nil, nil)
	case method.PATCH:
		Patch(rt, nil, nil)
	case method.DELETE:
		Delete(rt, nil)
	case method.OPTIONS:
		Options(rt, nil)
	}
}

func TestMethodNotDefined(t *testing.T) {
	testMethodNotDefined(method.POST, method.GET, t)
	testMethodNotDefined(method.POST, method.PUT, t)
	testMethodNotDefined(method.POST, method.PATCH, t)
	testMethodNotDefined(method.POST, method.DELETE, t)
	testMethodNotDefined(method.POST, method.OPTIONS, t)
	testMethodNotDefined(method.GET, method.POST, t)
	testMethodNotDefined(method.OPTIONS, method.GET, t)
}

func TestRouteIsNil(t *testing.T) {
	var rt *route.Route

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, route.ErrRouteIsNil{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(route.ErrRouteIsNil)
		_ = err.Error()

	}()

	rt.MustURL()
}

/*
func TestHandlerAlreadyDefined(t *testing.T) {
	route := New("/route")
	route.SetHandlerForMethod(noop{}, method.GET)

	defer func() {
		e := recover()
		errMsg := errorMustBe(e, ErrHandlerAlreadyDefined{})

		if errMsg != "" {
			t.Error(errMsg)
			return
		}

		err := e.(ErrHandlerAlreadyDefined)
		_ = err.Error()

		if err.Method != method.GET {
			t.Errorf("wrong method: %#v, expected: %v", err.Method, method.GET)
		}
	}()

	route.SetHandlerForMethod(noop{}, method.GET)
}



// ErrPairParams

*/
