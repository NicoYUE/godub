package main

import (
	"fmt"
	"github.com/NicoYUE/godub/audioop"
)

func main() {
	e := audioop.NewError("Hello, world: %d", 100)
	fmt.Println(e.Error())
}
