package puppy

import (
	"github.com/manwarcr/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark(s string) string {
	return dog.WhenGrownUp(s)
}
