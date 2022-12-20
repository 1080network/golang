package shared

import (
	"context"
	"regexp"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/require"
)

var bearerRegex = regexp.MustCompile(`([Bb]earer\s+)(\S+)`)

func Test_JWT_NoClaims(t *testing.T) {
	tests  := []struct{
		name  string
		roles []string
		options []JWTPerRpcRequestOption
	}{
		{"default usage", []string{"RolePartnerHumanAdmin"}, []JWTPerRpcRequestOption{}},
		{"with multiple Roles", []string{"RolePartnerHumanAdmin", "SecondTestRole"}, []JWTPerRpcRequestOption{}},
		{"with other key", []string{"SomeTestRole"}, []JWTPerRpcRequestOption{WithKey([]byte("4N+Yi3CfDX3ZWnNDxpO24FLwrqgRbmRd7424rUnBWw8="))}},
		{"with secure", []string{"another role"}, []JWTPerRpcRequestOption{WithSecure()}},
		{"with claims", []string{"fakerole"}, []JWTPerRpcRequestOption{WithClaims(map[string]string{"tenant":"test"})}},
		{"with several options", []string{"superroleone", "superroletwo"}, []JWTPerRpcRequestOption{
			WithClaims(map[string]string{"tenant":"test", "otherclaim": "claim2"}),
		    WithKey([]byte("pd6JGGZ1DHtmOE1LYJTd0/LEQ+OWHdUgNrEc/5h1oEA="))}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			perRpcCred := BuildJWTPerRpcCredential(test.roles, test.options...)
			require.NotNil(t, perRpcCred)
			testCtx := context.Background()
			metaMap, err := perRpcCred.GetRequestMetadata(testCtx)
			require.NoError(t, err)
			require.Contains(t, metaMap, "authentication")
			header := metaMap["authentication"]
			require.Contains(t, header, "Bearer")
			items := bearerRegex.FindStringSubmatch(header)
			require.Len(t, items,  3)
			tokenStr := items[2]
			verifyToken(t, []byte(tokenStr), perRpcCred.roles, perRpcCred.claims, perRpcCred.jwtKey)
		})
	}
}


func verifyToken(t *testing.T, token []byte, roles []string, claims map[string]string, key []byte){
	parsedJwt, err := jwt.Parse(token, jwt.WithKey(jwa.HS256, key))
	require.NoError(t, err)
	tokenClaims := parsedJwt.PrivateClaims()
	require.Contains(t, tokenClaims, "roles")
	rawRoles := tokenClaims["roles"].([]interface{})
	require.NotEmpty(t, rawRoles)
	tokenRoles := make([]string, len(rawRoles))
	for i, v := range rawRoles {
    	tokenRoles[i] = v.(string)
	}
	require.NotEmpty(t, tokenRoles)
	require.ElementsMatch(t, tokenRoles, roles)
	for key, value :=  range claims{
		require.Contains(t, tokenClaims, key)
		claim := tokenClaims[key]
		require.Equal(t, value, claim.(string))
	}
}