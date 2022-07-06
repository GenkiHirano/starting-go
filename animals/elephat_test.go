package animals

import (
	"testing"
)

func TestElephatFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephatFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
