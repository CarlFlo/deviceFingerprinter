package devicefingerprint

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

func GetFingerprint() string {
	// Gathering system information
	osInfo := runtime.GOOS     // Operating system
	arch := runtime.GOARCH     // Architecture
	numCPU := runtime.NumCPU() // Number of CPUs

	// Gathering information about network interfaces
	var netInfo []string
	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {
		if interf.Flags&net.FlagUp == 0 || interf.Flags&net.FlagLoopback != 0 {
			continue
		}
		if addr := interf.HardwareAddr.String(); addr != "" {
			netInfo = append(netInfo, addr)
		}
	}
	netInfoString := strings.Join(netInfo, "")

	// Gathering environment variables
	envVars := os.Environ()
	envString := strings.Join(envVars, "")

	// Combing and hashes (sha256) the values to create a fingerprint
	fingerprint := fmt.Sprintf("%s%s%d%s%s", osInfo, arch, numCPU, netInfoString, envString)
	fingerprintHash := sha256.New()
	fingerprintHash.Write([]byte(fingerprint))
	return hex.EncodeToString(fingerprintHash.Sum(nil))
}
