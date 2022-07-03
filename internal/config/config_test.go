package config

import (
	"testing"
)

func TestSettings(t *testing.T) {
	s, err := New(".", "settings", "yaml")
	if err != nil {
		t.Fatal(err.Error())
	}
	if s == nil {
		t.Fatal("Empty settings")
	}
	t.Logf("%+v", s)
}
