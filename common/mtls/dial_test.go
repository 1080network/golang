package mtls_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"testing"
	"time"

	pingv1 "github.com/1080network/golang/common/v2/mica/member/ping/v1"
	"github.com/1080network/golang/common/v2/mtls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type pingServer struct {
	pingv1.UnimplementedPingServiceServer
}

func (pingServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return &pingv1.PingResponse{Status: pingv1.PingResponse_STATUS_SUCCESS, ServiceType: "test"}, nil
}

// issue creates a certificate signed by parent (or self-signed when parent is nil).
func issue(t *testing.T, name string, isCA bool, parent *x509.Certificate, parentKey *ecdsa.PrivateKey) (*x509.Certificate, *ecdsa.PrivateKey, []byte, []byte) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	serial, _ := rand.Int(rand.Reader, big.NewInt(1<<62))
	tmpl := &x509.Certificate{
		SerialNumber:          serial,
		Subject:               pkix.Name{CommonName: name},
		NotBefore:             time.Now().Add(-time.Minute),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
		BasicConstraintsValid: true,
		IsCA:                  isCA,
	}
	if isCA {
		tmpl.KeyUsage |= x509.KeyUsageCertSign
	}
	signer, signerKey := tmpl, key
	if parent != nil {
		signer, signerKey = parent, parentKey
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, signer, &key.PublicKey, signerKey)
	if err != nil {
		t.Fatal(err)
	}
	cert, _ := x509.ParseCertificate(der)
	keyDER, _ := x509.MarshalECPrivateKey(key)
	return cert, key,
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
		pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})
}

// TestDialPing makes a real mutual TLS Ping call to an in-process server, checking that the server
// requires and verifies the client certificate that mtls.Dial presents.
func TestDialPing(t *testing.T) {
	ca, caKey, caPEM, _ := issue(t, "test ca", true, nil, nil)
	_, _, serverCertPEM, serverKeyPEM := issue(t, "server", false, ca, caKey)
	_, _, clientCertPEM, clientKeyPEM := issue(t, "client", false, ca, caKey)

	serverPair, err := tls.X509KeyPair(serverCertPEM, serverKeyPEM)
	if err != nil {
		t.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(ca)
	server := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverPair},
		ClientCAs:    pool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})))
	pingv1.RegisterPingServiceServer(server, pingServer{})
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	go server.Serve(lis)
	t.Cleanup(server.Stop)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("with a client certificate", func(t *testing.T) {
		conn, err := mtls.Dial(lis.Addr().String(), mtls.Config{CertificateAuthorities: caPEM, Certificate: clientCertPEM, Key: clientKeyPEM})
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()
		resp, err := pingv1.NewPingServiceClient(conn).Ping(ctx, &pingv1.PingRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if resp.GetStatus() != pingv1.PingResponse_STATUS_SUCCESS {
			t.Fatalf("unexpected status %v", resp.GetStatus())
		}
	})

	t.Run("a server certificate from an unknown authority is rejected", func(t *testing.T) {
		_, _, otherCAPEM, _ := issue(t, "other ca", true, nil, nil)
		conn, err := mtls.Dial(lis.Addr().String(), mtls.Config{CertificateAuthorities: otherCAPEM, Certificate: clientCertPEM, Key: clientKeyPEM})
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()
		if _, err := pingv1.NewPingServiceClient(conn).Ping(ctx, &pingv1.PingRequest{}); err == nil {
			t.Fatal("expected the call to fail")
		}
	})
}

func TestBadConfig(t *testing.T) {
	if _, err := mtls.Dial("localhost:1", mtls.Config{}); err == nil {
		t.Fatal("expected an error for an empty configuration")
	}
}
