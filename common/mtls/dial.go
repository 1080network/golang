// Package mtls builds gRPC client connections to the Mica network.
//
// Every connection to Mica uses mutual TLS: the client presents a certificate issued by Mica and verifies
// Mica's server certificate against the root certificate authority Mica provides.
package mtls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Config holds the PEM-encoded material for a mutual TLS connection.
type Config struct {
	// CertificateAuthorities is the PEM-encoded root certificate(s) used to verify Mica's server certificate.
	CertificateAuthorities []byte
	// Certificate is the PEM-encoded client certificate chain issued by Mica.
	Certificate []byte
	// Key is the PEM-encoded private key for Certificate.
	Key []byte
}

// ConfigFromFiles reads a Config from PEM files on disk.
func ConfigFromFiles(caFile, certFile, keyFile string) (Config, error) {
	var cfg Config
	var err error
	if cfg.CertificateAuthorities, err = os.ReadFile(caFile); err != nil {
		return Config{}, fmt.Errorf("reading certificate authorities: %w", err)
	}
	if cfg.Certificate, err = os.ReadFile(certFile); err != nil {
		return Config{}, fmt.Errorf("reading client certificate: %w", err)
	}
	if cfg.Key, err = os.ReadFile(keyFile); err != nil {
		return Config{}, fmt.Errorf("reading client key: %w", err)
	}
	return cfg, nil
}

// TLSConfig builds the client-side *tls.Config for the configuration.
func (c Config) TLSConfig() (*tls.Config, error) {
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(c.CertificateAuthorities) {
		return nil, errors.New("no certificates found in the certificate authorities PEM")
	}
	cert, err := tls.X509KeyPair(c.Certificate, c.Key)
	if err != nil {
		return nil, fmt.Errorf("loading client certificate and key: %w", err)
	}
	return &tls.Config{
		MinVersion:   tls.VersionTLS12,
		RootCAs:      pool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

// Dial opens a client connection to target (for example "api.example.mica.io:443") using mutual TLS.
// Extra grpc.DialOptions are applied after the transport credentials.
func Dial(target string, cfg Config, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	tlsConfig, err := cfg.TLSConfig()
	if err != nil {
		return nil, err
	}
	opts = append([]grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))}, opts...)
	return grpc.NewClient(target, opts...)
}
