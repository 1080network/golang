package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"ten80/proto/common/authenticatev1"
	"ten80/proto/common/enums/realmv1"
	"ten80/proto/common/pingv1"
	partnerserviceaccount "ten80/proto/partner/serviceaccountv1"
	partnerservice "ten80/proto/partner/servicev1"
	spserviceaccount "ten80/proto/sp/serviceaccountv1"
	spservice "ten80/proto/sp/servicev1"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const usageDescription = `
This cli demonstrates how to generate a jwt private key credential to use endpoints requiring authn/authz
Example usage of this cli:
	Generate a private RSA key and it's corresponding public key locally:
		openssl genrsa -out private-key.pem 4096
		openssl rsa -in private-key.pem -pubout -out public-key.pem

	Authenticate with an admin account:
		export JWT=$(./ten80 partner auth -endpoint grpc.5m9gg.demo.1080.network:443 -username "Partner-Admin-[secretid]" -password "Partner-Admin-[secretid]")

	Upload the public key w/ comma separated list of roles using the admin account:
		 export CLIENT_ID=$(./ten80 partner create-service-account -endpoint grpc.5m9gg.$domain:443 -jwt $JWT -publickey=public-key.pem) -roles AclGroupPartner

	Create a JWT for the client to use:
		 export SA_JWT=$(./ten80 generate-jwt -clientid $CLIENT_ID -rsa_private_key_file private-key.pem)

	Call an authn/authz protected API using the signed JWT from above:
		grpcurl -rpc-header="Authorization: Bearer $SA_JWT"  grpc.5m9gg.$domain:443 partner.service.v1.PartnerTo1080Service/SearchServiceAccount

`

func withBearerJwtToContext(ctx context.Context, jwt string) context.Context {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", jwt)})
	authContext := metadata.NewOutgoingContext(ctx, md)
	return authContext
}

func generateAuthJwt(clientID string, privateKeyPEM []byte) (string, error) {
	saJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims{
		"iss": clientID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return "", err
	}
	newJwtToken, err := saJwt.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return newJwtToken, nil
}

func createPartnerClient(hostnamePort string) partnerservice.PartnerTo1080ServiceClient {
	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	conn, err := grpc.Dial(hostnamePort, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		fmt.Printf("failed to dial %s", hostnamePort)
		os.Exit(1)
	}
	client := partnerservice.NewPartnerTo1080ServiceClient(conn)

	return client
}

func createSPClient(hostnamePort string) spservice.SPTo1080ServiceClient {
	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	conn, err := grpc.Dial(hostnamePort, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		fmt.Printf("failed to dial %s", hostnamePort)
		os.Exit(1)
	}
	client := spservice.NewSPTo1080ServiceClient(conn)

	return client
}

func main() {
	var hostnamePort, username, password, authJwt, publicKeyPEM, roles string
	pingCmd := flag.NewFlagSet("ping", flag.ExitOnError)
	pingCmd.StringVar(&hostnamePort, "endpoint", "localhost:8080", "the endpoint to call")

	authCmd := flag.NewFlagSet("auth", flag.ExitOnError)
	authCmd.StringVar(&hostnamePort, "endpoint", "localhost:8080", "the endpoint to call")
	authCmd.StringVar(&username, "username", "username_value", "the username")
	authCmd.StringVar(&password, "password", "password_value", "the password")

	createServiceAccountCmd := flag.NewFlagSet("create-service-account", flag.ExitOnError)
	createServiceAccountCmd.StringVar(&hostnamePort, "endpoint", "localhost:8080", "the endpoint to call")
	createServiceAccountCmd.StringVar(&authJwt, "jwt", "[[sample jwt]]", "the credential")
	createServiceAccountCmd.StringVar(&publicKeyPEM, "publickey", "public_key.pem", "pem formatted public key file")
	createServiceAccountCmd.StringVar(&roles, "roles", "", "comma separate list of roles")

	var rsaPrivateKeyFile, clientID string
	generateJwtCmd := flag.NewFlagSet("create-service-account", flag.ExitOnError)
	generateJwtCmd.StringVar(&rsaPrivateKeyFile, "rsa_private_key_file", "private-key.pem", "the private rsa key pem file")
	generateJwtCmd.StringVar(&clientID, "clientid", "", "the client id")

	if len(os.Args) < 3 {
		usage()
	}

	switch os.Args[1] {
	case "generate-jwt":
		generateJwtCmd.Parse(os.Args[2:])
		fileContents, err := ioutil.ReadFile(rsaPrivateKeyFile)
		if err != nil {
			fmt.Printf("failed to open rsa private key: %s", err.Error())
			os.Exit(1)
		}
		if !strings.HasPrefix(clientID, "sa_") {
			fmt.Printf("client_id has the format 'sa_[[uuid]]'")
			os.Exit(1)
		}
		jwt, err := generateAuthJwt(clientID, fileContents)
		if err != nil {
			fmt.Printf("failed to generate jwt: %s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s", jwt)
	case "partner":
		switch os.Args[2] {
		case "ping":
			pingCmd.Parse(os.Args[3:])
			client := createPartnerClient(hostnamePort)
			response, err := client.Ping(context.Background(), &pingv1.PingRequest{})
			if err != nil {
				fmt.Printf("failed to ping %s: %s", hostnamePort, err.Error())
				os.Exit(1)
			}
			fmt.Printf("%+v", response)
		case "auth":
			authCmd.Parse(os.Args[3:])
			client := createPartnerClient(hostnamePort)

			authResponse, err := client.Authenticate(context.Background(), &authenticatev1.AuthenticateRequest{
				Realm:    realmv1.Realm_REALM_ADMIN,
				Username: username,
				Password: password,
			})
			if err != nil {
				fmt.Printf("failed to authenticate %s: %s", hostnamePort, err.Error())
				os.Exit(1)
			}
			if authResponse.Status == authenticatev1.AuthenticateResponse_STATUS_SUCCESS {
				fmt.Printf("%s", authResponse.Token)
			} else {
				fmt.Printf("bad response: %+v", authResponse)
				os.Exit(1)
			}
		case "create-service-account":
			createServiceAccountCmd.Parse(os.Args[3:])
			publicKeyFileContent, err := ioutil.ReadFile(publicKeyPEM)
			if err != nil {
				fmt.Printf("failed to read file %s: %s", publicKeyPEM, err.Error())
				os.Exit(1)
			}
			client := createPartnerClient(hostnamePort)
			authContext := withBearerJwtToContext(context.Background(), authJwt)
			roleArr := strings.Split(roles, ",")
			createResponse, err := client.CreateServiceAccount(authContext, &partnerserviceaccount.CreateServiceAccountRequest{
				Roles: roleArr,
				AuthenticationProperties: &partnerserviceaccount.CreateServiceAccountRequest_JwtInfo{
					JwtInfo: &partnerserviceaccount.JwtPrivateKey{
						VerificationPublicKeyPem: string(publicKeyFileContent),
					},
				},
			})
			if err != nil {
				fmt.Printf("failed to create service account %s: %s", hostnamePort, err.Error())
				os.Exit(1)
			}
			if createResponse.Status == partnerserviceaccount.CreateServiceAccountResponse_STATUS_SUCCESS {
				fmt.Printf("%s", createResponse.ClientId)
			} else {
				fmt.Printf("bad response: %+v", createResponse)
				os.Exit(1)
			}
		default:
			fmt.Printf("unsupported partner command %s", os.Args[2])
			usage()
		}
	case "sp":
		switch os.Args[2] {
		case "ping":
			pingCmd.Parse(os.Args[3:])
			client := createSPClient(hostnamePort)

			response, err := client.Ping(context.Background(), &pingv1.PingRequest{})
			if err != nil {
				fmt.Printf("failed to ping %s: %s", hostnamePort, err.Error())
				os.Exit(1)
			}
			fmt.Printf("%+v", response)
		case "auth":
			authCmd.Parse(os.Args[3:])
			client := createSPClient(hostnamePort)

			authResponse, err := client.Authenticate(context.Background(), &authenticatev1.AuthenticateRequest{
				Realm:    realmv1.Realm_REALM_ADMIN,
				Username: username,
				Password: password,
			})
			if err != nil {
				fmt.Printf("failed to ping %s: %s", hostnamePort, err.Error())
				os.Exit(1)
			}
			if authResponse.Status == authenticatev1.AuthenticateResponse_STATUS_SUCCESS {
				fmt.Printf("%s", authResponse.Token)
			} else {
				fmt.Printf("bad response: %+v", authResponse)
				os.Exit(1)
			}
		case "create-service-account":
			createServiceAccountCmd.Parse(os.Args[3:])
			publicKeyFileContent, err := ioutil.ReadFile(publicKeyPEM)
			if err != nil {
				fmt.Printf("failed to read file %s: %s", publicKeyPEM, err.Error())
				os.Exit(1)
			}
			client := createSPClient(hostnamePort)
			authContext := withBearerJwtToContext(context.Background(), authJwt)
			roleArr := strings.Split(roles, ",")
			createResponse, err := client.CreateServiceAccount(authContext, &spserviceaccount.CreateServiceAccountRequest{
				Roles: roleArr,
				AuthenticationProperties: &spserviceaccount.CreateServiceAccountRequest_JwtInfo{
					JwtInfo: &spserviceaccount.JwtPrivateKey{
						VerificationPublicKeyPem: string(publicKeyFileContent),
					},
				},
			})
			if createResponse.Status == spserviceaccount.CreateServiceAccountResponse_STATUS_SUCCESS {
				fmt.Printf("%s", createResponse.ClientId)
			} else {
				fmt.Printf("bad response: %+v", createResponse)
				os.Exit(1)
			}
		default:
			fmt.Printf("unsupported sp command %s", os.Args[2])
			usage()
		}
	default:
		fmt.Printf("unsupported server type %s", os.Args[1])
		usage()
	}
}

func usage() {
	flag.PrintDefaults()
	fmt.Printf(usageDescription)
	fmt.Printf("usage:\nten80 partner (ping|auth|create-service-account)\n")
	fmt.Printf("ten80 sp (ping|auth|create-service-account)\n")
	fmt.Printf("ten80 generate-jwt -clientid [clientid] -rsa_private_key_file [pubkey.pem]")
	os.Exit(1)
}
