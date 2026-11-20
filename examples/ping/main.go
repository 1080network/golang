// Command ping connects to Mica with mutual TLS and calls the Ping service, to check that your
// certificates and network path work.
//
//	go run ./examples/ping -target api.example.mica.io:443 -ca rootca.crt -cert client.crt -key client.key
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pingv1 "github.com/1080network/golang/common/v2/mica/member/ping/v1"
	"github.com/1080network/golang/common/v2/mtls"
)

func main() {
	target := flag.String("target", "", "Mica endpoint, host:port")
	ca := flag.String("ca", "rootca.crt", "root certificate authority PEM file")
	cert := flag.String("cert", "client.crt", "client certificate PEM file")
	key := flag.String("key", "client.key", "client private key PEM file")
	flag.Parse()
	if *target == "" {
		log.Fatal("-target is required")
	}

	cfg, err := mtls.ConfigFromFiles(*ca, *cert, *key)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := mtls.Dial(*target, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	resp, err := pingv1.NewPingServiceClient(conn).Ping(ctx, &pingv1.PingRequest{})
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}
	fmt.Printf("status=%v service=%s version=%s\n", resp.GetStatus(), resp.GetServiceType(), resp.GetBuildVersion())
}
