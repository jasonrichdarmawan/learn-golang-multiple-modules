package main

import (
	"log"

	"github.com/kidfrom/learn-golang-multiple-modules/hello"
	"github.com/kidfrom/learn-golang-multiple-modules/world"
	"golang.org/x/example/stringutil"
)

func HelloWorld() {
	hello.Hello()
	world.World()
	log.Println(stringutil.Reverse("Hello"))
}
