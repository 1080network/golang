package shared

import (
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"sync"
	"time"
)
import "context"


type JWTAuthCredential struct {
	secure     bool
	jwtKey     []byte
	expiration time.Time
	roles      []string
	claims     map[string]string
	authHeader string
	mutex      sync.Mutex
}

type JWTPerRpcRequestOption func (cred *JWTAuthCredential)

func WithSecure() JWTPerRpcRequestOption {
	return func (cred *JWTAuthCredential){
		cred.secure = true
	}
}

func WithKey(keyBytes []byte) JWTPerRpcRequestOption{
	return func(cred *JWTAuthCredential) {
		cred.jwtKey = keyBytes
	}
}

func WithTenant(tenant string) JWTPerRpcRequestOption{
	return func(cred *JWTAuthCredential)  {
		if cred.claims == nil{
			cred.claims = map[string]string{"tenant": tenant}
		}else{
			cred.claims["tenant"] = tenant
		}
	}
}

func WithClaims(claims map[string]string) JWTPerRpcRequestOption{
	return func(cred *JWTAuthCredential) {
		if cred.claims == nil {
			cred.claims = claims
		}else {
			for key, value := range claims{
				cred.claims[key] = value
			}
		}
	}
}

func BuildJWTPerRpcCredential(roles []string,  options ...JWTPerRpcRequestOption) *JWTAuthCredential {
	authCredential := &JWTAuthCredential{
		secure: false,
		jwtKey: []byte("idv7Nf/7476PxbMX6iVIXHCc/kMSqSjPQNgBNDU2kz4="),
		roles: roles,
	}

	for _, optionSet := range options{
		optionSet(authCredential)
	}
	return authCredential
}


func (jac *JWTAuthCredential) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	metadataContent, _ := jac.getAuthHeader()
	return map[string]string{"authentication": metadataContent}, nil
}

func (jac *JWTAuthCredential) RequireTransportSecurity() bool {
	return jac.secure
}

func (jac *JWTAuthCredential) getAuthHeader() (string, error) {
	expiry := time.Now().Add(time.Minute)
	if jac.authHeader != "" && jac.expiration.After(expiry) {
		return jac.authHeader, nil
	}
	jac.mutex.Lock()
	defer jac.mutex.Unlock()
	// in case it has been refreshed
	if jac.expiration.After(expiry) {
		return jac.authHeader, nil
	}
	err := jac.buildAuthHeader()
	if err != nil {
		return "", err
	}
	return jac.authHeader, nil
}

func (h *JWTAuthCredential) buildAuthHeader() error {
	h.expiration = time.Now().Add(1 * time.Hour)
	tokenbuilder := jwt.NewBuilder().Issuer("mica.io").IssuedAt(time.Now()).Claim("roles", h.roles).
		Expiration(h.expiration)
	for claim, value := range h.claims {
		tokenbuilder.Claim(claim, value)
	}
	token, err := tokenbuilder.Build()
	if err != nil {
		return err
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, h.jwtKey))
	if err != nil {
		return err
	}
	h.authHeader = fmt.Sprintf("Bearer %s", signed)
	return nil
}
