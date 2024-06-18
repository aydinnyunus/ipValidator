# ipValidator

`ipValidator` is a Go library designed to help developers easily filter out and validate IP addresses based on various criteria. It allows you to check if an IP address falls into reserved categories such as loopback, link-local, multicast, or private addresses. This can be useful in various scenarios, such as preventing requests to non-routable addresses or securing network communications.

## Features

- **Loopback Address Check**: Detects and optionally disallows loopback addresses (e.g., `127.0.0.1` and `::1`).
- **Link-Local Address Check**: Detects and optionally disallows link-local addresses (e.g., `169.254.0.0/16` for IPv4 and `fe80::/10` for IPv6).
- **Multicast Address Check**: Detects and optionally disallows multicast addresses (e.g., `224.0.0.0/4` for IPv4 and `ff00::/8` for IPv6).
- **Private Address Check**: Detects and optionally disallows private addresses (e.g., `10.0.0.0/8`, `172.16.0.0/12`, and `192.168.0.0/16` for IPv4).
- **Customizable Filters**: Easily enable or disable specific checks based on your requirements.

## Installation

To install `ipValidator`, use `go get`:

```bash
go get github.com/aydinnyunus/ipValidator/ipValidator
```

## Usage

### Creating a Default Filter

By default, all reserved IP checks are enabled:

```go
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

}
```

### Creating a Custom Filter

You can customize the filter to disallow specific types of addresses:

```go
package main

import (
	"fmt"
	"github.com/aydinnyunus/ipValidator/ipValidator"
)

func main() {
	// Create a custom filter with specific flags
	customFilter := ipValidator.NewValidator(true, false, true, false, false)

	testIPs := []string{
		"127.0.0.1",
		"::1",
		"169.254.0.1",
		"192.168.1.1",
		"224.0.0.1",
		"8.8.8.8",
		"2001:db8::",
	}

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
```

## License

`ipValidator` is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

[<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/linkedin.png" title="LinkedIn">](https://linkedin.com/in/yunus-ayd%C4%B1n-b9b01a18a/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/github.png" title="Github">](https://github.com/aydinnyunus/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/instagram-new.png" title="Instagram">](https://instagram.com/aydinyunus_/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/twitter-squared.png" title="LinkedIn">](https://twitter.com/aydinnyunuss)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/usestrix/cli.svg?style=for-the-badge

[contributors-url]: https://github.com/aydinnyunus/ipValidator/graphs/contributors


[forks-url]: https://github.com/aydinnyunus/ipValidator/network/members


[stars-url]: https://github.com/aydinnyunus/ipValidator/stargazers


[issues-url]: https://github.com/aydinnyunus/ipValidator/issues


[license-url]: https://github.com/aydinnyunus/ipValidator/blob/master/LICENSE.txt

[linkedin-url]: https://linkedin.com/in/aydinnyunus

[latest-release]: https://github.com/aydinnyunus/ipValidator/releases
