package main

import (
	"flag"
	"github.com/miekg/dns"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	port = flag.Int("port", 15353, "Run DNS port")
	host = flag.String("host", "0.0.0.0", "Run DNS host")
)

func serveDNS(server *dns.Server) {
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func startDNSServer() {
	dns.HandleFunc(".", handlerQuery)
	udpServer := &dns.Server{
		Addr: *host + ":" + strconv.Itoa(*port),
		Net:  "udp",
	}
	defer udpServer.Shutdown()
	go serveDNS(udpServer)

	tcpServer := &dns.Server{
		Addr: *host + ":" + strconv.Itoa(*port),
		Net:  "tcp",
	}
	defer tcpServer.Shutdown()
	go serveDNS(tcpServer)

}

func handlerQuery(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]
	if q.Qtype != dns.TypeAXFR && q.Qtype != dns.TypeIXFR {
		log.Println("is not XFR")
		return
	}

	ch := make(chan *dns.Envelope)
	defer close(ch)
	tr := new(dns.Transfer)
	go tr.Out(w, r, ch)

	soa, err := dns.NewRR("example.com.	600	IN	SOA	ns.example.com. test.example.com. 2009032802 21600 7200 604800 3600")
	if err != nil {
		log.Printf("Create SOA Error: %x\n", err)
		return
	}
	a, err := dns.NewRR("www.example.com. 600 IN A 10.1.1.1")
	if err != nil {
		log.Printf("Create A Error: %x\n", err)
		return
	}
	mx, err := dns.NewRR("example.com.	600	IN	MX	1 example.com.")
	if err != nil {
		log.Printf("Create MX Error: %x\n", err)
		return
	}
        // TODO: Divide response size by dns query max size
	ch <- &dns.Envelope{RR: []dns.RR{soa, a, mx, soa}}
	w.Hijack()
}

func main() {
	startDNSServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
}
