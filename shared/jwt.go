package shared

import (
	"encoding/base64"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sync"
	"time"
)
import "context"

func JWTClientInterceptor(roles []string, jwtSecret string) (grpc.UnaryClientInterceptor, error) {
	jwtHandler, err := buildJwtHandler(jwtSecret, roles)
	if err != nil {
		return nil, err
	}
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		authorization, err := jwtHandler.getAuthHeader()
		if err != nil {
			return err
		}
		metadata.AppendToOutgoingContext(ctx, "authorization", authorization)
		return invoker(ctx, method, req, reply, cc, opts...)
	}, nil
}

type jwthandler struct {
	jwtKey     []byte
	expiration time.Time
	roles      []string
	authHeader string
	mutex      sync.Mutex
}

func buildJwtHandler(key string, roles []string) (*jwthandler, error) {
	defaultKey := "idv7Nf/7476PxbMX6iVIXHCc/kMSqSjPQNgBNDU2kz4="
	if len(key) > 0 {
		defaultKey = key
	}
	b, err := base64.StdEncoding.DecodeString(defaultKey)
	if err != nil {
		return nil, err
	}
	handler := &jwthandler{
		jwtKey:     b,
		roles:      roles,
		authHeader: "",
		mutex:      sync.Mutex{},
	}
	err = handler.buildAuthHeader()
	if err != nil {
		return nil, err
	}
	return handler, nil
}

func (h *jwthandler) getAuthHeader() (string, error) {
	expiry := time.Now().Add(time.Minute)
	if h.expiration.After(expiry) {
		return h.authHeader, nil
	}
	h.mutex.Lock()
	defer h.mutex.Unlock()
	// in case it has been refreshed
	if h.expiration.After(expiry) {
		return h.authHeader, nil
	}
	err := h.buildAuthHeader()
	if err != nil {
		return "", err
	}
	return h.authHeader, nil
}

func (h *jwthandler) buildAuthHeader() error {
	h.expiration = time.Now().Add(1 * time.Hour)
	tok, err := jwt.NewBuilder().Issuer("mica.io").IssuedAt(time.Now()).Claim("roles", h.roles).
		Expiration(h.expiration).
		Build()
	if err != nil {
		return err
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, h.jwtKey))
	if err != nil {
		return err
	}
	h.authHeader = fmt.Sprintf("Bearer %s", signed)
	return nil
}
