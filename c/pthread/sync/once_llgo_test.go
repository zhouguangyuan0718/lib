//go:build llgo

/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sync_test

import (
	stdsync "sync"
	"testing"

	llsync "github.com/goplus/lib/c/pthread/sync"
)

var (
	rawOnceCount int
	rawOnceDelta int
)

func addRawOnceDelta() {
	rawOnceCount += rawOnceDelta
}

func incrementRawOnceCount() {
	rawOnceCount++
}

func TestOnceDoRawCallback(t *testing.T) {
	once := llsync.OnceInit
	rawOnceCount, rawOnceDelta = 0, 2

	if result := once.Do(addRawOnceDelta); result != 0 {
		t.Fatalf("first Once.Do returned %d", result)
	}
	if result := once.Do(incrementRawOnceCount); result != 0 {
		t.Fatalf("second Once.Do returned %d", result)
	}
	if rawOnceCount != 2 {
		t.Fatalf("raw callback ran incorrectly: got %d, want 2", rawOnceCount)
	}
}

func TestOnceDoFuncCapturingClosure(t *testing.T) {
	once := llsync.OnceInit
	count, delta := 0, 3

	if result := once.DoFunc(func() { count += delta }); result != 0 {
		t.Fatalf("first Once.DoFunc returned %d", result)
	}
	if result := once.DoFunc(func() { count++ }); result != 0 {
		t.Fatalf("second Once.DoFunc returned %d", result)
	}
	if count != 3 {
		t.Fatalf("capturing closure ran incorrectly: got %d, want 3", count)
	}
}

func TestOnceDoFuncConcurrentContexts(t *testing.T) {
	once := llsync.OnceInit
	value := 0
	results := make(chan int, 4)
	var callers stdsync.WaitGroup
	callers.Add(4)

	for candidate := 1; candidate <= 4; candidate++ {
		candidate := candidate
		go func() {
			defer callers.Done()
			results <- int(once.DoFunc(func() { value = candidate }))
		}()
	}
	callers.Wait()
	close(results)

	for result := range results {
		if result != 0 {
			t.Fatalf("Once.DoFunc returned %d", result)
		}
	}
	if value < 1 || value > 4 {
		t.Fatalf("no concurrent context ran: got %d", value)
	}
}

func TestOnceDoFuncNested(t *testing.T) {
	outer := llsync.OnceInit
	inner := llsync.OnceInit
	value := 0
	innerResult := 0

	if result := outer.DoFunc(func() {
		innerResult = int(inner.DoFunc(func() { value = 5 }))
	}); result != 0 {
		t.Fatalf("outer Once.DoFunc returned %d", result)
	}
	if innerResult != 0 {
		t.Fatalf("inner Once.DoFunc returned %d", innerResult)
	}
	if value != 5 {
		t.Fatalf("nested context ran incorrectly: got %d, want 5", value)
	}
}
