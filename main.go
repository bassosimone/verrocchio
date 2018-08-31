package main

import (
	"log"

	"github.com/miekg/dns"
)

func main() {
	m := new(dns.Msg)
	m.SetQuestion("www.google.com.", dns.TypeA)
	c := new(dns.Client)
	in, rtt, err := c.Exchange(m, "194.116.72.177:53")
	log.Println(in, rtt, err)
}
