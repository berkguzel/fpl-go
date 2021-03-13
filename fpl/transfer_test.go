package fpl

import (
	"testing"
)

func TestListTransfers(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.ListTransfers("")
	if got != nil {
		t.Errorf("Could do not be succed, got %+v", got)
	}
}
