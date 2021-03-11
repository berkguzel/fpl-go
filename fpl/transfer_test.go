package fpl

import (
	"testing"
)

func TestListTransfers(t *testing.T) {

	c := NewClient(nil)

	got, err := c.ListTransfers("")
	if err != nil {
		t.Error(err)
	}

	if got == nil {
		t.Errorf("Could do not be succed, got %+v", got)
	}
}
