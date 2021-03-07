package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"ten80/proto/common/pingv1"
	partnerservice "ten80/proto/partner/servicev1"
	spservice "ten80/proto/sp/servicev1"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	serverType := os.Args[1]
	if serverType != "partner" && serverType != "sp" {
		log.Printf("unsupported serverType %s", serverType)
		usage()
	}
	address := os.Args[2]

	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	log.Printf("dialing %s", address)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("failed to dial %s", address)
	}

	switch serverType {
	case "partner":
		client := partnerservice.NewPartnerTo1080ServiceClient(conn)
		response, err := client.Ping(context.Background(), &pingv1.PingRequest{})
		if err != nil {
			log.Fatalf("failed to ping %s: %s", address, err.Error())
		}

		log.Printf("successful response: %+v", response)
	case "sp":
		client := spservice.NewSPTo1080ServiceClient(conn)
		response, err := client.Ping(context.Background(), &pingv1.PingRequest{})
		if err != nil {
			log.Fatalf("failed to ping %s: %s", address, err.Error())
		}

		log.Printf("successful response: %+v", response)
	default:
		log.Fatalf("unsupported serverType %s", serverType)
	}
}

func usage() {
	log.Fatalf("usage: ./ten80 (partner|sp) <host:port>")
}
