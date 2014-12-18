package routegopherjs

import (
	"gopkg.in/go-on/router.v2/route"
	"net/http"

	"github.com/gopherjs/gopherjs/js"
	"gopkg.in/go-on/method.v1"

	"testing"
)

type noop struct{}

var allMethods = []method.Method{
	method.GET,
	method.POST,
	method.PUT,
	method.PATCH,
	method.DELETE,
	method.OPTIONS,
}

func (noop) ServeHTTP(rw http.ResponseWriter, req *http.Request) {}

func TestHXR(t *testing.T) {
	rt := route.New("/", method.GET, method.POST, method.PATCH, method.PUT, method.DELETE, method.OPTIONS)

	route.Mount("/", rt)

	aj := &Funcs{}

	methCalled := []method.Method{}

	aj.GET = func(url string, callback func(js.Object)) {
		methCalled = append(methCalled, method.GET)
	}

	aj.POST = func(url string, data interface{}, callback func(js.Object)) {
		methCalled = append(methCalled, method.POST)
	}

	aj.PUT = func(url string, data interface{}, callback func(js.Object)) {
		methCalled = append(methCalled, method.PUT)
	}

	aj.PATCH = func(url string, data interface{}, callback func(js.Object)) {
		methCalled = append(methCalled, method.PATCH)
	}

	aj.DELETE = func(url string, callback func(js.Object)) {
		methCalled = append(methCalled, method.DELETE)
	}

	aj.OPTIONS = func(url string, callback func(js.Object)) {
		methCalled = append(methCalled, method.OPTIONS)
	}
	xhr = nil
	RegisterService(aj)
	expectedMethCalled := 0

	Get(rt, nil)
	expectedMethCalled++
	if len(methCalled) != expectedMethCalled {
		t.Errorf("ajax %s not called", method.GET)
	}

	Post(rt, nil, nil)
	expectedMethCalled++
	if len(methCalled) != expectedMethCalled {
		t.Errorf("ajax %s not called", method.POST)
	}

	Put(rt, nil, nil)
	expectedMethCalled++
	if len(methCalled) != expectedMethCalled {
		t.Errorf("ajax %s not called", method.PUT)
	}

	Patch(rt, nil, nil)
	expectedMethCalled++
	if len(methCalled) != expectedMethCalled {
		t.Errorf("ajax %s not called", method.PATCH)
	}

	Delete(rt, nil)
	expectedMethCalled++
	if len(methCalled) != expectedMethCalled {
		t.Errorf("ajax %s not called", method.DELETE)
	}

	Options(rt, nil)
	expectedMethCalled++
	if len(methCalled) != expectedMethCalled {
		t.Errorf("ajax %s not called", method.OPTIONS)
	}
}
