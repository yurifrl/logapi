package store

import (
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	store := Create()

	store.Bump("foo")
	t.Errorf("expected '%v', got '%v'", 1, store.Get("foo"))

	go func() {
		store.Bump("foo")
	}()

	time.Sleep(200000) // to wait sync

	go func() {
		t.Errorf("expected '%v', got '%v'", 2, store.Get("foo"))
	}()

	t.Errorf("expected '%v', got '%v'", 3, store.Get("foo"))

	time.Sleep(200000) //to see the output
}
