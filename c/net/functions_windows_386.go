//go:build windows && 386

// Copyright (c) 2026 The GoPlus Authors. Licensed under the Apache License 2.0.

package net

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

// Winsock uses __stdcall on Windows/386, which LLGo cannot currently express
// on a C declaration. This C file is only a calling-convention bridge; it
// preserves the native Winsock signatures and adds no cross-platform behavior.
const LLGoFiles = "_wrap/net_windows.c"

//go:linkname WSAStartup C.llgo_net_startup
func WSAStartup(version uint16, data *WSAData) c.Int

//go:linkname WSACleanup C.llgo_net_cleanup
func WSACleanup() c.Int

//go:linkname Getaddrinfo C.llgo_net_getaddrinfo
func Getaddrinfo(host *c.Char, port *c.Char, addrInfo *AddrInfo, result **AddrInfo) c.Int

//go:linkname Freeaddrinfo C.llgo_net_freeaddrinfo
func Freeaddrinfo(addrInfo *AddrInfo)

//go:linkname Socket C.llgo_net_socket
func Socket(domain c.Int, typ c.Int, protocol c.Int) SocketT

//go:linkname Bind C.llgo_net_bind
func Bind(sockfd SocketT, addr *SockAddr, addrlen SocklenT) c.Int

//go:linkname Connect C.llgo_net_connect
func Connect(sockfd SocketT, addr *SockAddr, addrlen SocklenT) c.Int

//go:linkname Listen C.llgo_net_listen
func Listen(sockfd SocketT, backlog c.Int) c.Int

//go:linkname Accept C.llgo_net_accept
func Accept(sockfd SocketT, addr *SockAddr, addrlen *SocklenT) SocketT

//go:linkname Closesocket C.llgo_net_close
func Closesocket(sockfd SocketT) c.Int

//go:linkname GetHostByName C.llgo_net_gethostbyname
func GetHostByName(name *c.Char) *Hostent

//go:linkname InetNtop C.llgo_net_inet_ntop
func InetNtop(af c.Int, src c.Pointer, dst *c.Char, size uintptr) *c.Char

//go:linkname InetAddr C.llgo_net_inet_addr
func InetAddr(value *c.Char) c.Uint

//go:linkname Send C.llgo_net_send
func Send(sockfd SocketT, buffer c.Pointer, length c.Int, flags c.Int) c.Int

//go:linkname Recv C.llgo_net_recv
func Recv(sockfd SocketT, buffer c.Pointer, length c.Int, flags c.Int) c.Int

//go:linkname SetSockOpt C.llgo_net_setsockopt
func SetSockOpt(sockfd SocketT, level c.Int, optionName c.Int, optionValue c.Pointer, optionLength SocklenT) c.Int

//go:linkname Ntohs C.llgo_net_ntohs
func Ntohs(value uint16) uint16

//go:linkname Htons C.llgo_net_htons
func Htons(value uint16) uint16

//go:linkname Ntohl C.llgo_net_ntohl
func Ntohl(value c.Uint) c.Uint

//go:linkname Htonl C.llgo_net_htonl
func Htonl(value c.Uint) c.Uint

//go:linkname WSAGetLastError C.llgo_net_last_error
func WSAGetLastError() c.Int
