package main

import (
	"fmt"
	"github.com/aydinnyunus/ipValidator/ipValidator"
)

func main() {
	// Create a filter with default settings (all flags set to false)
	defaultFilter := ipValidator.NewDefaultValidator()

	testIPs := []string{
		"127.0.0.1",
		"::1",
		"169.254.0.1",
		"192.168.1.1",
		"224.0.0.1",
		"8.8.8.8",
		"2001:db8::",
		"[::1]",
		"216.58.212.14",
	}

	fmt.Println("Default Validator :")
	for _, ipStr := range testIPs {
		if defaultFilter.IsReserved(ipStr) == 1 {
			fmt.Printf("\t- %s is a reserved IP address\n", ipStr)
		} else if defaultFilter.IsReserved(ipStr) == -1 {
			fmt.Printf("\t- %s is not a valid IP address\n", ipStr)
		} else {
			fmt.Printf("\t- %s is not a reserved IP address\n", ipStr)
		}
	}

	fmt.Println()
	fmt.Println("Custom Validator with allowLoopback and allowMulticast:")
	customFilter := ipValidator.NewValidator(true, false, true, false, false)

	for _, ipStr := range testIPs {
		if customFilter.IsReserved(ipStr) == 1 {
			fmt.Printf("\t- %s is a reserved IP address\n", ipStr)
		} else if customFilter.IsReserved(ipStr) == -1 {
			fmt.Printf("\t- %s is not a valid IP address\n", ipStr)
		} else {
			fmt.Printf("\t- %s is not a reserved IP address\n", ipStr)
		}
	}

}
