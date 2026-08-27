//go:build windows

// Copyright (c) 2026 The GoPlus Authors. Licensed under the Apache License 2.0.

package net

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

// AddrInfo matches the Winsock ADDRINFOA layout. ai_addrlen is size_t on
// Windows, unlike the socklen_t field used by Unix implementations.
type AddrInfo struct {
	Flags     c.Int
	Family    c.Int
	SockType  c.Int
	Protocol  c.Int
	AddrLen   uintptr
	CanOnName *c.Char
	Addr      *SockAddr
	Next      *AddrInfo
}

//go:linkname Getaddrinfo C.llgo_net_getaddrinfo
func Getaddrinfo(host *c.Char, port *c.Char, addrInfo *AddrInfo, result **AddrInfo) c.Int

//go:linkname freeaddrinfo C.llgo_net_freeaddrinfo
func freeaddrinfo(addrInfo *AddrInfo)

// Freeaddrinfo matches the existing cross-platform API. The native function
// returns void, so a completed call reports zero.
func Freeaddrinfo(addrInfo *AddrInfo) c.Int {
	freeaddrinfo(addrInfo)
	return 0
}
