package api

import (
	"testing"
)

func TestReloadData(t *testing.T) {
	res := ReloadData()
	if res {
		t.Log("successful")
	}
}
