package api

import (
	"testing"
)

func TestPut(t *testing.T) {
	kvEntry := struct {
		key   string
		value string
	}{
		"name1",
		"guoda2",
	}
	want := true
	got := Put(kvEntry.key, kvEntry.value)
	if got != want {
		t.Errorf("Put(%q,%q) == %q, want %q", kvEntry.key, kvEntry.value, got, want)
	}
}

func TestGet(t *testing.T) {
	key := "name"
	got := Get(key)
	want := "guoda2"
	if got != want {
		t.Errorf("Get(%q) == %q, want %q", key, got, want)
	}
}

func TestListAllKV(t *testing.T) {
	kvListStr := ListAllKV()
	t.Log(kvListStr)
}
