package main

import (
	"fmt"
	"os"

	"github.com/jsfour/awsdyndns/v1/src"
	"github.com/spf13/pflag"
)

var zoneID string
var hostname string

func init() {
	pflag.StringVarP(&zoneID, "zoneid", "z", "", "Route 53 zone id")
	pflag.StringVarP(&hostname, "dnshostname", "d", "", "Hostname on the zone id to update")
	pflag.Parse()
}

func main() {

	if zoneID == "" {
		fmt.Println("You must specify a zoneid. Run with -h to see usage.")
		os.Exit(1)
	}

	if hostname == "" {
		fmt.Println("You must specify a hostname. Run with -h to see usage.")
		os.Exit(1)
	}

	ip, err := src.GetPublicIp()

	if err != nil {
		fmt.Println(err)
	}

	err = src.SetDns(ip, hostname, zoneID)
	if err != nil {
		fmt.Println(err)
	}
}
