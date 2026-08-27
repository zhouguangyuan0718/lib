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
	"runtime"
	stdsync "sync"
	"testing"
	"time"

	"github.com/goplus/lib/c"
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

func TestOnceDoCRawCallback(t *testing.T) {
	once := llsync.OnceInit
	rawOnceCount, rawOnceDelta = 0, 2

	if result := once.DoC(addRawOnceDelta); result != 0 {
		t.Fatalf("first Once.DoC returned %d", result)
	}
	if result := once.DoC(incrementRawOnceCount); result != 0 {
		t.Fatalf("second Once.DoC returned %d", result)
	}
	if rawOnceCount != 2 {
		t.Fatalf("raw callback ran incorrectly: got %d, want 2", rawOnceCount)
	}
}

type onceDoer interface {
	Do(func()) c.Int
}

var _ onceDoer = (*llsync.Once)(nil)

func TestOnceDoCompatibility(t *testing.T) {
	once := llsync.OnceInit
	count, delta := 0, 2
	if result := once.Do(func() { count += delta }); result != 0 {
		t.Fatalf("Once.Do returned %d", result)
	}
	if count != 2 {
		t.Fatalf("Once.Do callback ran incorrectly: got %d, want 2", count)
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

func TestMutexAndRWLock(t *testing.T) {
	var mutex llsync.Mutex
	if result := mutex.Init(nil); result != 0 {
		t.Fatalf("Mutex.Init returned %d", result)
	}
	defer mutex.Destroy()
	mutex.Lock()
	mutex.Unlock()
	if result := mutex.TryLock(); result != 0 {
		t.Fatalf("Mutex.TryLock returned %d", result)
	}
	mutex.Unlock()

	var rw llsync.RWLock
	if result := rw.Init(nil); result != 0 {
		t.Fatalf("RWLock.Init returned %d", result)
	}
	defer rw.Destroy()
	if result := rw.TryRLock(); result != 0 {
		t.Fatalf("RWLock.TryRLock returned %d", result)
	}
	rw.RUnlock()
	if result := rw.TryLock(); result != 0 {
		t.Fatalf("RWLock.TryLock returned %d", result)
	}
	rw.Unlock()
}

func TestCondSignal(t *testing.T) {
	var mutex llsync.Mutex
	if result := mutex.Init(nil); result != 0 {
		t.Fatalf("Mutex.Init returned %d", result)
	}
	defer mutex.Destroy()

	var cond llsync.Cond
	if result := cond.Init(nil); result != 0 {
		t.Fatalf("Cond.Init returned %d", result)
	}
	defer cond.Destroy()

	waiting := make(chan struct{})
	done := make(chan c.Int, 1)
	go func() {
		mutex.Lock()
		close(waiting)
		result := cond.Wait(&mutex)
		mutex.Unlock()
		done <- result
	}()

	<-waiting
	mutex.Lock()
	result := cond.Signal()
	mutex.Unlock()
	if result != 0 {
		t.Fatalf("Cond.Signal returned %d", result)
	}

	select {
	case result := <-done:
		if result != 0 {
			t.Fatalf("Cond.Wait returned %d", result)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Cond.Wait did not wake after Cond.Signal")
	}
}

func TestSemaphore(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("macOS does not implement unnamed POSIX semaphores")
	}

	var sem llsync.Sem
	if result := sem.Init(0, 1); result != 0 {
		t.Fatalf("Sem.Init returned %d", result)
	}
	defer func() {
		if result := sem.Destroy(); result != 0 {
			t.Errorf("Sem.Destroy returned %d", result)
		}
	}()

	var value c.Int
	if result := sem.GetValue(&value); result != 0 || value != 1 {
		t.Fatalf("Sem.GetValue = (%d, %d), want (0, 1)", result, value)
	}
	if result := sem.TryWait(); result != 0 {
		t.Fatalf("Sem.TryWait returned %d", result)
	}
	if result := sem.TryWait(); result == 0 {
		t.Fatal("Sem.TryWait succeeded with no available value")
	}
	if result := sem.Post(); result != 0 {
		t.Fatalf("Sem.Post returned %d", result)
	}
	if result := sem.Wait(); result != 0 {
		t.Fatalf("Sem.Wait returned %d", result)
	}
}
