package utils

import (
	"fmt"
	"time"
)

func RandomNameGenerator() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	fmt.Println(name)
}
