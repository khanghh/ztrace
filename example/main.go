package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/zartbot/ztrace"
)

var (
	dst      string  = ""
	src      string  = ""
	asndb    string  = ""
	geodb    string  = ""
	protocol string  = "udp"
	maxPath  int     = 16
	maxTTL   int     = 64
	pps      float64 = 1
	wmode    bool    = false
	lat      float64 = 31.02
	long     float64 = 121.1
)

func init() {
	flag.StringVar(&protocol, "proto", protocol, "Protocol[icmp|tcp|udp]")
	flag.StringVar(&src, "src", src, "Source ")
	flag.StringVar(&asndb, "asndb", asndb, "ASN Database")
	flag.StringVar(&geodb, "geodb", geodb, "Geo Database")
	flag.IntVar(&maxPath, "path", maxPath, "Max ECMP Number")
	flag.IntVar(&maxPath, "p", maxPath, "Max ECMP Number")
	flag.IntVar(&maxTTL, "ttl", maxTTL, "Max TTL")
	flag.Float64Var(&pps, "rate", pps, "Packet Rate per second")
	flag.Float64Var(&pps, "r", pps, "Packet Rate per second")
	flag.BoolVar(&wmode, "wide", wmode, "Widescreen mode")
	flag.BoolVar(&wmode, "w", wmode, "Widescreen mode")
	flag.Float64Var(&lat, "lat", lat, "Latitude")
	flag.Float64Var(&lat, "long", long, "Longitude")
	flag.Parse()
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  ./ztrace [-src source] [-proto protocol] [-ttl ttl] [-rate packetRate] [-wide Widescreen mode] [-path NumOfECMPPath] [-ansdb asndb] [-geodb geodb] host")
	fmt.Println("Example:")
	fmt.Println(" ./ztrace www.cisco.com")
	fmt.Println(" ./ztrace -ttl 30 -rate 1 -path 8 -wide www.cisco.com")
	fmt.Println("Option:")
	flag.PrintDefaults()
}

func main() {
	if flag.NArg() != 1 {
		PrintUsage()
		return
	} else {
		dst = flag.Arg(0)
		fmt.Println(dst)
	}

	if asndb == "" {
		asndb = os.Getenv("ASNDB")
	}

	if geodb == "" {
		geodb = os.Getenv("GEODB")
	}

	t := ztrace.New(protocol, dst, src, maxPath, uint8(maxTTL), float32(pps), 0, wmode, asndb, geodb)
	t.Latitude = lat
	t.Longitude = long

	t.Start()
	go t.Report(time.Second)
	time.Sleep(time.Second * 100)
	t.Stop()
}
