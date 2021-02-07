package help

import (
	"testing"
)

type Str struct {
	Name string
	Is   bool
}

func Test_strings(t *testing.T) {

	ffxNames := []Str{
		{
			Name: "Rikku",
			Is:   false,
		},
		{
			Name: "Wakka",
			Is:   false,
		},
		{
			Name: "",
			Is:   true,
		},
		{
			Name: "Tidus",
			Is:   false,
		},
	}

	for _, persona := range ffxNames {
		if got := IsEmpty(persona.Name); got != persona.Is {
			t.Errorf("result: %v but expected: %v", got, persona.Is)

		}
	}
}
