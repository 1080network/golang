package example

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/administrationv1"
	"github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/instrumentv1"
	"github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/servicev1"
	"github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/userv1"
	"github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/uuekv1"
	"github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/valuev1"
	"github.com/1080network/golang/serviceprovider/proto/micashared/common/enums/approvaltypev1"
	"github.com/1080network/golang/serviceprovider/proto/micashared/common/enums/currencyv1"
	"github.com/1080network/golang/serviceprovider/proto/micashared/common/enums/instrumenttypev1"
	"github.com/1080network/golang/serviceprovider/proto/micashared/common/pingv1"
	commonv1 "github.com/1080network/golang/serviceprovider/proto/micashared/common/v1"
	"github.com/1080network/golang/shared"
	"github.com/cockroachdb/apd/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"testing"
	"time"
)

var startTime *timestamppb.Timestamp

type ServiceProvider struct {
	servicev1.UnimplementedServiceProviderFromMicaServiceServer
}

func (sp ServiceProvider) ObtainValue(ctx context.Context, request *valuev1.ObtainValueRequest) (*valuev1.ObtainValueResponse, error) {
	response := &valuev1.ObtainValueResponse{
		//simulate a fake Id
		TransactionRef: uuid.NewString(),
	}
	if request.ApprovalType == approvaltypev1.ApprovalType_APPROVAL_TYPE_FULL {
		response.Status = valuev1.ObtainValueResponse_STATUS_FULL_APPROVAL
		response.ApprovedAmount = request.Value.RequestedAmount
	} else if request.ApprovalType == approvaltypev1.ApprovalType_APPROVAL_TYPE_PARTIAL {
		requested, _, err := apd.NewFromString(request.Value.RequestedAmount)
		if err != nil {
			response.Status = valuev1.ObtainValueResponse_STATUS_ERROR
			return response, nil
		}
		result := apd.New(0, 0)
		apd.BaseContext.Sub(result, requested, apd.New(1, -2))
		response.ApprovedAmount = result.String()
		response.Status = valuev1.ObtainValueResponse_STATUS_PARTIAL_APPROVAL
	}
	return response, nil
}

func (sp ServiceProvider) Ping(ctx context.Context, request *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return &pingv1.PingResponse{
		Status:          pingv1.PingResponse_STATUS_SUCCESS,
		ServerTime:      timestamppb.New(time.Now()),
		ServerStartTime: startTime,
		BuildVersion:    "some build version",
		BuildSha:        "some buid sha",
		BuildTime:       "some build time",
		ServiceType:     "From Mica service",
	}, nil
}

func StartServer() {
	serverCertificatePem, err := loadPemBlob("SERVER_CERTIFICATE_PEM")
	if err != nil {
		panic(err)
	}
	serverKeyPem, err := loadPemBlob("CLIENT_KEY_PEM")
	if err != nil {
		panic(err)
	}
	rootCAPem, err := loadPemBlob("ROOT_CA")
	if err != nil {
		panic(err)
	}
	cert, err := tls.X509KeyPair(serverCertificatePem, serverKeyPem)
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(rootCAPem); !ok {
		panic("failed to append root CA cert for grpc Server")
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	grpcTlsCredentials := grpc.Creds(credentials.NewTLS(tlsConfig))
	server := grpc.NewServer(grpcTlsCredentials)
	serviceImpl := ServiceProvider{}
	servicev1.RegisterServiceProviderFromMicaServiceServer(server, serviceImpl)
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", 84434)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("unable to start grpc server")
	}
	startTime = timestamppb.New(time.Now())
	err = server.Serve(lis)
	if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		panic(err)
	}
}

func TestPingExternal(t *testing.T) {
	adminClient := getAdminClient()
	adminClient.PingExternal(context.TODO(), &pingv1.PingRequest{})
}

func TestProvisionUuek(t *testing.T) {
	ctx := context.TODO()
	micaClient := getServiceProviderClient()
	request := &uuekv1.ProvisionServiceProviderUUEKRequest{
		Criteria: &uuekv1.ProvisionServiceProviderUUEKRequest_ServiceProviderInstrumentRef{
			ServiceProviderInstrumentRef: "11a431ab-1168-4b35-bc80-5a3cab2a4968",
		},
		UseCriteria: &uuekv1.ProvisionServiceProviderUUEKRequest_NumberOfUses{
			NumberOfUses: 1,
		},
	}
	response, err := micaClient.ProvisionServiceProviderUUEK(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, uuekv1.ProvisionServiceProviderUUEKResponse_STATUS_SUCCESS, response.Status)
}

func TestRegisterUser(t *testing.T) {
	ctx := context.TODO()
	micaClient := getServiceProviderClient()
	request := &userv1.RegisterUserRequest{
		ServiceProviderUserRef: "4b1519fc-dbf8-4156-b401-35b8058098f4",
		UserDemographic: &commonv1.UserDemographic{
			Phone: "19993214567",
		},
	}
	response, err := micaClient.RegisterUser(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, userv1.RegisterUserResponse_STATUS_SUCCESS, response.Status)
}

func TestRegisterInstrument(t *testing.T) {
	ctx := context.TODO()
	micaClient := getServiceProviderClient()
	request := &instrumentv1.RegisterInstrumentRequest{
		ServiceProviderUserKey:       "EKmril69iaoeoQOadTiaFav2gKTJQLQ",
		ServiceProviderInstrumentRef: "22419a4b-6b82-454e-8a10-78094369d62e",
		InstrumentType:               instrumenttypev1.InstrumentType_INSTRUMENT_TYPE_CHECKING,
		Currency:                     currencyv1.Currency_CURRENCY_USD,
		Nickname:                     "1225",
	}
	response, err := micaClient.RegisterInstrument(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, instrumentv1.RegisterInstrumentResponse_STATUS_SUCCESS, response.Status)
}

func getServiceProviderClient() servicev1.ServiceProviderToMicaServiceClient {
	certificateAuthorities, err := loadPemBlob("ROOT_CA_PEM")
	if err != nil {
		panic(err)
	}
	clientSecretKey, err := loadPemBlob("CLIENT_KEY_PEM")
	if err != nil {
		panic(err)
	}
	clientCertificate, err := loadPemBlob("CLIENT_CERTIFICATE_PEM")
	if err != nil {
		panic(err)
	}
	conn, err := shared.SecureClientDial("api.hron3n.members.mica.io:443", certificateAuthorities, clientCertificate, clientSecretKey)
	if err != nil {
		panic(err)
	}
	return servicev1.NewServiceProviderToMicaServiceClient(conn)
}

func getAdminClient() administrationv1.ServiceProviderAdministrationServiceClient {
	certificateAuthorities, err := loadPemBlob("ROOT_CA_PEM")
	if err != nil {
		panic(err)
	}
	clientSecretKey, err := loadPemBlob("CLIENT_KEY_PEM")
	if err != nil {
		panic(err)
	}
	clientCertificate, err := loadPemBlob("CLIENT_CERTIFICATE_PEM")
	if err != nil {
		panic(err)
	}
	conn, err := shared.SecureClientDial("api.hron3n.members.mica.io:443", certificateAuthorities, clientCertificate, clientSecretKey)
	if err != nil {
		panic(err)
	}
	return administrationv1.NewServiceProviderAdministrationServiceClient(conn)
}
