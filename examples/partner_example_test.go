package example

import (
	"context"
	"github.com/1080network/golang/partner"
	partnersvc "github.com/1080network/golang/partner/proto/mica/partner/servicev1"
	"github.com/1080network/golang/partner/proto/micashared/common/pingv1"
	"github.com/1080network/golang/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initializeClient() (partnersvc.PartnerToMicaServiceClient, error) {
	var certificateAuthorities []byte
	var clientSecretKey []byte
	var clientCertificate []byte
	conn, err := shared.SecureClientDial("api.hron3n.members.mica.io:443", certificateAuthorities, clientSecretKey, clientCertificate)
	if err != nil {
		return nil, err
	}
	partnerClient := partner.BuildPartnerToMicaClient(conn)
	return partnerClient, nil
}

func Test_ping(t *testing.T) {
	client, err := initializeClient()
	assert.NoError(t, err)
	response, err := client.Ping(context.TODO(), &pingv1.PingRequest{})
	assert.NoError(t, err)
	assert.Equal(t, pingv1.PingResponse_STATUS_SUCCESS, response.Status)
}
