package main

import "fmt"

type IPAddr [4]byte

// String ...
func (ip IPAddr) String() string {
	res := ""
	for _, v := range ip {
		res += fmt.Sprint(v)
		res += "."
	}

	return res[:len(res)-1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
