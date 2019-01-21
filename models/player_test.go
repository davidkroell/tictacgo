package models

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	want := "david"
	got := NewPlayer(want)

	if got.Name != want {
		t.Errorf("Want %s, got %s\nUnique ID was %s", want, got.Name, got.ID)
	}
}
