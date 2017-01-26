package main

import (
	"fmt"
)

type EnvVar struct {
	key string
	val string
}

func (e *EnvVar) get() string {
	return fmt.Sprintf("KEY: %s, VALUE: %s", e.key, e.val)
}
