package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	const token = "6df4a6fe00930a"

	var targetIP string
	fmt.Printf("Target IP Address:  ")
	fmt.Scan(&targetIP)

	getInfo(targetIP, token)
}

func getInfo(ipAddr string, token string) {
	// params: httpClient, cache, token. `http.DefaultClient` and no cache will be used in case of `nil`.
	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ipAddr))
	if err != nil {
		log.Fatal(err)
	}

	if info != nil {
		fmt.Printf("\n------------------info------------------\n")
		fmt.Printf("Location: %v, %v\n", info.Country, info.City)
		fmt.Printf("Coordinates: %v\n", info.Location)
		fmt.Printf("Company: %v, Abuse: %v\n", info.Company, info.Abuse)
		fmt.Printf("ASN: %v\n", info.ASN)
		fmt.Printf("Hostname: %v\n", info.Hostname)
		fmt.Printf("Anycast: %v\n", info.Anycast)
		fmt.Printf("Bogon: %v\n", info.Bogon)
		fmt.Printf("Carrier: %v\n", info.Carrier)
		fmt.Printf("Postal: %v\n", info.Postal)
		fmt.Printf("Domains: %v\n", info.Domains)
	} else {
		fmt.Printf("No info found\n")
	}
	if info.Privacy != nil {
		fmt.Printf("\n-----------------privacy----------------\n")
		fmt.Printf("VPN: %v\n", info.Privacy.VPN)
		fmt.Printf("Tor: %v\n", info.Privacy.Tor)
		fmt.Printf("Proxy: %v\n", info.Privacy.Proxy)
		fmt.Printf("Relay: %v\n", info.Privacy.Relay)
		fmt.Printf("Hosting: %v\n", info.Privacy.Hosting)
	} else {
		fmt.Printf("\nNo info found on privacy\n")
	}
}
