//go:build llgo && !windows

// Copyright (c) 2026 The GoPlus Authors. Licensed under the Apache License 2.0.

package pthread_test

import (
	"testing"

	"github.com/goplus/lib/c"
	"github.com/goplus/lib/c/pthread"
)

func keyDestructor(c.Pointer) {}

func TestKeyDestructorCallback(t *testing.T) {
	var key pthread.Key
	if result := key.Create(keyDestructor); result != 0 {
		t.Fatalf("pthread_key_create returned %d", result)
	}
	if result := key.Delete(); result != 0 {
		t.Fatalf("pthread_key_delete returned %d", result)
	}
}
