//go:build windows && !386

// Copyright (c) 2026 The GoPlus Authors. Licensed under the Apache License 2.0.

package net

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

// Windows/AMD64 and Windows/ARM64 use a single platform calling convention,
// so these declarations link directly to Winsock without an adapter.

//go:linkname WSAStartup C.WSAStartup
func WSAStartup(version uint16, data *WSAData) c.Int

//go:linkname WSACleanup C.WSACleanup
func WSACleanup() c.Int

//go:linkname Getaddrinfo C.getaddrinfo
func Getaddrinfo(host *c.Char, port *c.Char, addrInfo *AddrInfo, result **AddrInfo) c.Int

//go:linkname Freeaddrinfo C.freeaddrinfo
func Freeaddrinfo(addrInfo *AddrInfo)

//go:linkname Socket C.socket
func Socket(domain c.Int, typ c.Int, protocol c.Int) SocketT

//go:linkname Bind C.bind
func Bind(sockfd SocketT, addr *SockAddr, addrlen SocklenT) c.Int

//go:linkname Connect C.connect
func Connect(sockfd SocketT, addr *SockAddr, addrlen SocklenT) c.Int

//go:linkname Listen C.listen
func Listen(sockfd SocketT, backlog c.Int) c.Int

//go:linkname Accept C.accept
func Accept(sockfd SocketT, addr *SockAddr, addrlen *SocklenT) SocketT

//go:linkname Closesocket C.closesocket
func Closesocket(sockfd SocketT) c.Int

//go:linkname GetHostByName C.gethostbyname
func GetHostByName(name *c.Char) *Hostent

//go:linkname InetNtop C.inet_ntop
func InetNtop(af c.Int, src c.Pointer, dst *c.Char, size uintptr) *c.Char

//go:linkname InetAddr C.inet_addr
func InetAddr(value *c.Char) c.Uint

//go:linkname Send C.send
func Send(sockfd SocketT, buffer c.Pointer, length c.Int, flags c.Int) c.Int

//go:linkname Recv C.recv
func Recv(sockfd SocketT, buffer c.Pointer, length c.Int, flags c.Int) c.Int

//go:linkname SetSockOpt C.setsockopt
func SetSockOpt(sockfd SocketT, level c.Int, optionName c.Int, optionValue c.Pointer, optionLength SocklenT) c.Int

//go:linkname Ntohs C.ntohs
func Ntohs(value uint16) uint16

//go:linkname Htons C.htons
func Htons(value uint16) uint16

//go:linkname Ntohl C.ntohl
func Ntohl(value c.Uint) c.Uint

//go:linkname Htonl C.htonl
func Htonl(value c.Uint) c.Uint

//go:linkname WSAGetLastError C.WSAGetLastError
func WSAGetLastError() c.Int
