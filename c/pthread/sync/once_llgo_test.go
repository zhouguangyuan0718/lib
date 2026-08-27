//go:build llgo && !windows

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
	"testing"

	llsync "github.com/goplus/lib/c/pthread/sync"
)

var onceCount int

func incrementOnceCount() {
	onceCount++
}

func TestOnceDo(t *testing.T) {
	once := llsync.OnceInit
	onceCount = 0

	if result := once.Do(incrementOnceCount); result != 0 {
		t.Fatalf("first Once.Do returned %d", result)
	}
	if result := once.Do(incrementOnceCount); result != 0 {
		t.Fatalf("second Once.Do returned %d", result)
	}
	if onceCount != 1 {
		t.Fatalf("callback ran %d times, want 1", onceCount)
	}
}
