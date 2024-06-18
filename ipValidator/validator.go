package ipValidator

import (
	"net"
)

// Checker holds the configuration for IP checks.
type Validator struct {
	AllowLoopback           bool
	AllowLinkLocal          bool
	AllowMulticast          bool
	AllowPrivate            bool
	AllowLinkLocalMulticast bool
}

// NewValidator creates a new Checker with the specified flags.
// If no flags are provided, the default is to allow all reserved IP types.
func NewValidator(allowLoopback, allowLinkLocal, allowMulticast, allowPrivate, allowLinkLocalMulticast bool) *Validator {
	return &Validator{
		AllowLoopback:           allowLoopback,
		AllowLinkLocal:          allowLinkLocal,
		AllowMulticast:          allowMulticast,
		AllowPrivate:            allowPrivate,
		AllowLinkLocalMulticast: allowLinkLocalMulticast,
	}
}

// NewDefaultValidator creates a new Checker with all flags set to true.
func NewDefaultValidator() *Validator {
	return &Validator{
		AllowLoopback:           false,
		AllowLinkLocal:          false,
		AllowMulticast:          false,
		AllowPrivate:            false,
		AllowLinkLocalMulticast: false,
	}
}

// IsReserved checks if the given IP address falls into any reserved category.
func (c *Validator) IsReserved(ipStr string) int64 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return -1
	}

	// Check loopback
	if !c.AllowLoopback && ip.IsLoopback() {
		return 1
	}

	// Check link-local unicast
	if !c.AllowLinkLocal && ip.IsLinkLocalUnicast() {
		return 1
	}

	// Check link-local unicast
	if !c.AllowLinkLocalMulticast && ip.IsLinkLocalMulticast() {
		return 1
	}

	// Check multicast
	if !c.AllowMulticast && ip.IsMulticast() {
		return 1
	}

	// Check private
	if !c.AllowPrivate && ip.IsPrivate() {
		return 1
	}

	return 0
}
