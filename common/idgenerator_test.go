package common

import (
	"testing"
)

func TestFakeIDGenerator(t *testing.T) {
	id := "1234-abcd"

	f := FakeIDGenerator{id}

	fid, err := f.GenerateID()

	if err != nil {
		t.Fatalf("idgen failed: %v", err)
	}

	if id != fid {
		t.Fatalf("ids don't match")
	}
}
