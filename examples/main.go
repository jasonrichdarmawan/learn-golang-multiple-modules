package main

import (
	"log"

	"github.com/kidfrom/learn-golang-multiple-modules/hello"
	"golang.org/x/example/stringutil"
)

func main() {
	hello.HelloWorld()
	log.Println(stringutil.Reverse("Hello"))
}
