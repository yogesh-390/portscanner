package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	host := flag.String("host", "", "hostname or IP address to scan")
	startPort := flag.Int("startport", 1, "start port number")
	endPort := flag.Int("endport", 65535, "end port number")
	spoofIP := flag.String("spoofIP", "", "spoofed IP address")
	scanType := flag.String("scanType", "", "scan type (SYN, FIN, XMAS, NULL, ACK)")
	outputFormat := flag.String("outputFormat", "", "output format (JSON, CSV)")
	verbose := flag.Bool("verbose", false, "enable verbose output")
	flag.Parse()

	if *host == "" {
		fmt.Println("\033[1;32mWelcome to the port scanner tool!\033[0m")
		fmt.Println("\033[1;34mUsage:\033[0m go run portscan.go --host [hostname or IP address]")
		fmt.Println("Example: go run portscan.go --host scanme.nmap.org")
		fmt.Println("Advanced options: --startport [port number] --endport [port number] --spoofIP [IP address] --scanType [SYN, FIN, XMAS, NULL, ACK] --outputFormat [JSON, CSV] --verbose")
		fmt.Println("Note: This tool will scan all ports from 1 to 65535 by default.")
		fmt.Println("Please use it responsibly and with proper permissions.")
		return
	}

	fmt.Println("Starting scan on host", *host)
	fmt.Println("\033[1;31mCopyright © 2023 Nox | Mentor. All rights reserved.\033[0m")
	var wg sync.WaitGroup
	for i := *startPort; i <= *endPort; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			var address string
			if *spoofIP != "" {
				address = fmt.Sprintf("%s:%d", *spoofIP, port)
			} else {
				address = fmt.Sprintf("%s:%d", *host, port)
			}
			var conn net.Conn
			var err error
			switch *scanType {
			case "SYN":
				conn, err = net.DialTimeout("tcp", address, time.Duration(2*time.Second))
			case "FIN":
				// code for FIN scanning
			case "XMAS":
				// code for XMAS scanning
			case "NULL":
				// code for NULL scanning
			case "ACK":
				// code for ACK scanning
			default:
				conn, err = net.DialTimeout("tcp", address, time.Duration(2*time.Second))
			}
			if err != nil {
				return
			}
			conn.Close()
			if *outputFormat == "JSON" {
				// code to output in JSON format
			} else if *outputFormat == "CSV" {
				// code to output in CSV format
			} else {
				if *verbose {
					printVerboseOutput(port)
				}
			}
		}(i)
	}
	wg.Wait()

}

func printVerboseOutput(port int) {
	fmt.Printf("Port %d is open\n", port)
}
