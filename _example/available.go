package main

import (
	"github.com/octu0/unsafecgo"
)

func main() {
	if unsafecgo.Available() {
		println("available!")
	} else {
		println("not yet...")
	}
}
