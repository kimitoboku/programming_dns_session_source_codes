package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/miekg/dns"
)

func queryHandler(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range m.Question {
			switch q.Qtype {
			case dns.TypeA:
				rr, err := dns.NewRR(fmt.Sprintf("%s %d IN A %s", q.Name, 10, "10.1.1.1"))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			case dns.TypeAAAA:
				rr, err := dns.NewRR(fmt.Sprintf("%s %d IN AAAA %s", q.Name, 10, "2001:db8::1"))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}

	}
	w.WriteMsg(m)
}

var (
	port = flag.Int("port", 15353, "Run DNS port")
	zone = flag.String("zone", ".", "Run DNS zone")
	host = flag.String("host", "0.0.0.0", "Run DNS host")
	ttl  = flag.Int("ttl", -1, "DNS TTL")
)

func main() {
	flag.Parse()

	dns.HandleFunc(*zone, queryHandler)
	server := &dns.Server{Addr: *host + ":" + strconv.Itoa(*port), Net: "udp"}
	log.Printf("%s zone DNS Server run %s:%d\n", *zone, *host, *port)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start: %s\n", err.Error())
	}
}
