package puppy

import (
	"fmt"

	"github.com/manwarcr/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func from11() {
	fmt.Println("I'm V1.1.0")
}
