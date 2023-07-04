package main

import (
	"github.com/geremachek/arzhur/frame"
)

func main() {
	f, _ := frame.NewFrame("foo", "bar", "far")
	f.Start()
}