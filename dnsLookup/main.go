package main

import (
	"fmt"
	"net"
	"os"
)

func lookIp(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostName(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	input := args[1]
	IPAddress := net.ParseIP(input)

	if IPAddress == nil {
		IPs, err := lookHostName(input)
		if err == nil {
			for _, singleIP := range IPs {
				fmt.Println(singleIP)
			}
		}
	} else {
		hosts, err := lookIp(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}
	}
}
