Port Scanner Tool

This is a simple Go-based command line tool for port scanning. It allows you to scan a hostname or IP address for open ports and provides options for customizing the scan (e.g. specifying a range of ports to scan, spoofing the IP address, and choosing a scan type). The tool also supports output in JSON or CSV format.
Usage

To run the tool, navigate to the directory where the portscan.go file is located and run the following command:

go run portscan.go --host [hostname or IP address]

For example, to scan the host scanme.nmap.org, you would run:

go run portscan.go --host scanme.nmap.org

Advanced options

In addition to the --host flag, the following advanced options are available:

    --startport: specify a starting port number (default is 1)
    --endport: specify an ending port number (default is 65535)
    --spoofIP: specify a spoofed IP address
    --scanType: choose a scan type (SYN, FIN, XMAS, NULL, ACK)
    --outputFormat: choose an output format (JSON, CSV)

For example, to scan all ports from 100 to 200 on the host scanme.nmap.org, you would run:

go run portscan.go --host scanme.nmap.org --startport 100 --endport 200

Note

Please use this tool responsibly and with proper permissions. This tool will scan all ports from 1 to 65535 by default.

