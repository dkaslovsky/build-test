package main

import (
	"log"
	"testing"
)

func TestDummy(t *testing.T) {
	t.Run("dummy test", func(t *testing.T) {
		val1 := 101
		val2 := 101
		if val1 != val2 {
			log.Fatal("values are not equal")
		}
	})
}
