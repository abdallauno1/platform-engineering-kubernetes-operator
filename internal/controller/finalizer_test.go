package controller

import "testing"

func TestFinalizerLifecycle(t *testing.T) {
	fs, added := EnsureFinalizer(nil)
	if !added || len(fs) != 1 {
		t.Fatal("finalizer not added")
	}
	fs, added = EnsureFinalizer(fs)
	if added || len(fs) != 1 {
		t.Fatal("finalizer must be idempotent")
	}
	fs, removed := RemoveFinalizer(fs)
	if !removed || len(fs) != 0 {
		t.Fatal("finalizer not removed")
	}
}
