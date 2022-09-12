package shared

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestClientDial(target string, roles []string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	interceptor, err := JWTClientInterceptor(roles, "")
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	return grpc.Dial(target, opts...)
}

func SecureClientDial(target string, pemCertificateAuthorities []byte, pemClientCertificate []byte, pemClientKey []byte, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	cp := x509.NewCertPool()
	if ok := cp.AppendCertsFromPEM(pemCertificateAuthorities); !ok {
		return nil, fmt.Errorf("unable to add certificte authorities to certificate pool")
	}
	cert, err := tls.X509KeyPair(pemClientCertificate, pemClientKey)
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
		Certificates:       []tls.Certificate{cert},
	}
	credential := credentials.NewTLS(config)
	opts = append(opts, grpc.WithTransportCredentials(credential))
	return grpc.Dial(target, opts...)
}
