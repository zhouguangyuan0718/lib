/* Copyright (c) 2026 The GoPlus Authors. Licensed under Apache-2.0. */

#define WIN32_LEAN_AND_MEAN
#define _WINSOCK_DEPRECATED_NO_WARNINGS

#include <stdint.h>
#include <winsock2.h>
#include <ws2tcpip.h>

int llgo_net_startup(WORD version, WSADATA *data)
{
    return WSAStartup(version, data);
}

int llgo_net_cleanup(void) { return WSACleanup(); }

int llgo_net_getaddrinfo(
    const char *host, const char *service, const ADDRINFOA *hints,
    PADDRINFOA *result)
{
    return getaddrinfo(host, service, hints, result);
}

void llgo_net_freeaddrinfo(PADDRINFOA value)
{
    freeaddrinfo(value);
}

uintptr_t llgo_net_socket(int domain, int type, int protocol)
{
    return (uintptr_t)socket(domain, type, protocol);
}

int llgo_net_bind(uintptr_t socket_value, const struct sockaddr *address,
                  int address_length)
{
    return bind((SOCKET)socket_value, address, address_length);
}

int llgo_net_connect(uintptr_t socket_value, const struct sockaddr *address,
                     int address_length)
{
    return connect((SOCKET)socket_value, address, address_length);
}

int llgo_net_listen(uintptr_t socket_value, int backlog)
{
    return listen((SOCKET)socket_value, backlog);
}

uintptr_t llgo_net_accept(uintptr_t socket_value, struct sockaddr *address,
                          int *address_length)
{
    return (uintptr_t)accept((SOCKET)socket_value, address, address_length);
}

int llgo_net_close(uintptr_t socket_value)
{
    return closesocket((SOCKET)socket_value);
}

struct hostent *llgo_net_gethostbyname(const char *name)
{
    return gethostbyname(name);
}

const char *llgo_net_inet_ntop(int family, const void *source,
                               char *destination, size_t size)
{
    return inet_ntop(family, source, destination, size);
}

unsigned long llgo_net_inet_addr(const char *value)
{
    return inet_addr(value);
}

int llgo_net_send(uintptr_t socket_value, const void *buffer, int length,
                  int flags)
{
    return send((SOCKET)socket_value, (const char *)buffer, length, flags);
}

int llgo_net_recv(uintptr_t socket_value, void *buffer, int length,
                  int flags)
{
    return recv((SOCKET)socket_value, (char *)buffer, length, flags);
}

int llgo_net_setsockopt(uintptr_t socket_value, int level, int option_name,
                        const void *option_value, int option_length)
{
    return setsockopt((SOCKET)socket_value, level, option_name,
                      (const char *)option_value, option_length);
}

uint16_t llgo_net_ntohs(uint16_t value) { return ntohs(value); }
uint16_t llgo_net_htons(uint16_t value) { return htons(value); }
uint32_t llgo_net_ntohl(uint32_t value) { return ntohl(value); }
uint32_t llgo_net_htonl(uint32_t value) { return htonl(value); }
int llgo_net_last_error(void) { return WSAGetLastError(); }
