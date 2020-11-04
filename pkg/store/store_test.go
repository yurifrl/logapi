package store

import (
	"reflect"
	"testing"
)

func TestSomething(t *testing.T) {
	store := Create()

	store.Bump([]string{"instance", "app"})
	out, _ := store.GetAll()
	expect := make(map[string]string)
	expect["app"] = "1"
	expect["instance"] = "1/1"

	if !reflect.DeepEqual(out, expect) {
		t.Errorf("expected '%v', got '%v'", expect, out)
	}
}
