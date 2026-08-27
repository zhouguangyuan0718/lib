//go:build !windows

// Copyright (c) 2026 The GoPlus Authors. Licensed under the Apache License 2.0.

package net

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

const LLGoPackage = true

type AddrInfo struct {
	Flags     c.Int
	Family    c.Int
	SockType  c.Int
	Protocol  c.Int
	AddrLen   c.Uint
	CanOnName *c.Char
	Addr      *SockAddr
	Next      *AddrInfo
}

//go:linkname Getaddrinfo C.getaddrinfo
func Getaddrinfo(host *c.Char, port *c.Char, addrInfo *AddrInfo, result **AddrInfo) c.Int

//go:linkname freeaddrinfo C.freeaddrinfo
func freeaddrinfo(addrInfo *AddrInfo)

// Freeaddrinfo retains the package's original result for source compatibility.
// The C function returns void, so a completed call reports zero.
func Freeaddrinfo(addrInfo *AddrInfo) c.Int {
	freeaddrinfo(addrInfo)
	return 0
}
