# gomap

![GitHub](https://img.shields.io/github/license/JustinTimperio/gomap)
[![Go Reference](https://pkg.go.dev/badge/github.com/JustinTimperio/gomap.svg)](https://pkg.go.dev/github.com/JustinTimperio/gomap)
[![Go Report Card](https://goreportcard.com/badge/github.com/JustinTimperio/gomap)](https://goreportcard.com/report/github.com/JustinTimperio/gomap)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/47e878568ce04a819e82af10d3734062)](https://www.codacy.com/gh/JustinTimperio/gomap/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=JustinTimperio/gomap&amp;utm_campaign=Badge_Grade)

## What is gomap?
Gomap is a fully self-contained nmap like module for Golang. Unlike other projects which provide nmap bindings or rely on other local dependencies, gomap is a fully implemented in Go. Since this is a small library, it only focuses on providing a few core features for applications that require a completely embedded solution. 


## Features
  - Parallel port scanning using go routines
  - Automated CIDR range scanning
  - Service perdiction by port number
  - SYN (Silent) Scanning Mode
  - UDP Scanning (Non-Stealth)
  - Fast and detailed scanning for common ports
  - Pure Go with zero dependencies
  - Easily integrated into other projects

## Upcoming Features
  - More Known Ports for Detection
  - Stable Release

## Example Usage - 1
Performs a fastscan for the most common ports on every IP on a local range
### Create Files
 1. Create `quickscan.go`
```go
package main

import (
        "fmt"

        "github.com/JustinTimperio/gomap"
)

func main() {
        var (
                proto    = "tcp"
                fastscan = true
                syn      = false
        )

        scan := gomap.ScanRange(proto, fastscan, syn)
        fmt.Printf(scan.String())
}

```
 2. `go mod init quickscan`
 3. `go mod tidy`
 4. `go run quickscan.go`

### Example Output

```
Host: computer-name (192.168.1.132)
        |     Port      Service
        |     ----      -------
        |---- 22        ssh
 
Host: server-nginx (192.168.1.143)
        |     Port      Service
        |     ----      -------
        |---- 443       https
        |---- 80        http
        |---- 22        ssh
 
Host: server-minio (192.168.1.112)
        |     Port      Service
        |     ----      -------
        |---- 22        ssh

Host: some-phone (192.168.1.155)
        |- No Open Ports
 
```

## Example Usage - 2
Performs a detailed stealth scan on a single IP

### Create Files
 1. Create `stealthmap.go`
```go
package main

import (
        "fmt"

        "github.com/JustinTimperio/gomap"
)

func main() {
        // Stealth scans MUST be run as root/admin
        var (
                fastscan = false
                syn      = true
                proto    = "tcp"
                ip       = "192.168.1.120"
        )

        scan := gomap.ScanIP(ip, proto, fastscan, syn)
        fmt.Printf(scan.String())
}


```
 2. `go mod init stealthmap`
 3. `go mod tidy`
 4. `sudo go run stealthmap.go`

### Example Output

```
Host: 192.168.1.120 | Ports Scanned 3236/3236
Host: Voyager (192.168.1.120)
        |     Port      Service
        |     ----      -------
        |---- 22        SSH Remote Login Protocol
        |---- 80        World Wide Web HTTP
        |---- 443       HTTP protocol over TLS/SSL

```
