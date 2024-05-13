# Device Fingerprint Package for Go

![Tests](https://github.com/CarlFlo/deviceFingerprinter/actions/workflows/go.yml/badge.svg)

This Go package allows you to generate a unique fingerprint for a device based on various system properties such as operating system, architecture, number of CPUs, network interface details, and environment variables. The fingerprint is created using SHA-256 hashing.

## How it works
* Collects system information including operating system, architecture, and number of CPUs.
* Gathers active network interface MAC addresses (excluding loopback).
* Includes environmental variables in the fingerprint.
* Generates a SHA-256 hash to create a unique fingerprint.

## Usage
```go
package main

import (
    "fmt"
    "github.com/CarlFlo/deviceFingerprinter"
)

func main() {
    fingerprint := deviceFingerprinter.GetFingerprint()
    fmt.Printf("Device Fingerprint: %s\n", fingerprint)
}
```