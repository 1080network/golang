package example

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

func loadPemBlob(envVar string) ([]byte, error) {
	//expect the pem blob to be base64 encoded
	value, ok := os.LookupEnv(envVar)
	if !ok {
		return nil, errors.New(fmt.Sprintf("environment varible %s is not set", envVar))
	}
	return base64.StdEncoding.DecodeString(value)
}
