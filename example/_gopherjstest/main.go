package main

import (
	"gopkg.in/go-on/method.v1"
	"gopkg.in/go-on/router.v2/route"
)

var (
	A = route.New("/a", method.GET, method.POST)
	B = route.New("/b/:x/:y", method.PATCH, method.PUT)
)

func Mount(path string) {
	route.Mount(path, A, B)
}

func main() {
	Mount("/app")
	println(A.MustURL())
}
