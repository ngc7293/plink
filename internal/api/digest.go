package api

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DigestResponse struct {
	Realm string
	Nonce string
	Stale bool
}

func ParseDigestResponse(s string) (*DigestResponse, error) {
	response := &DigestResponse{}

	if !strings.HasPrefix(s, "Digest ") {
		return nil, errors.New("not a digest response")
	}

	fragments := strings.Split(s[7:], ",")

	for _, fragment := range fragments {
		tokens := strings.Split(fragment, "=")

		if len(tokens) != 2 {
			return nil, errors.New("unexpected fragment in Digest response")
		}

		key := strings.Trim(tokens[0], " ")
		value := strings.Trim(tokens[1], " ")

		switch key {
		case "realm":
			realm, err := strconv.Unquote(value)

			if err != nil {
				return nil, fmt.Errorf("unexpected value '%s' for key 'realm'", value)
			}

			response.Realm = realm
		case "nonce":
			nonce, err := strconv.Unquote(value)

			if err != nil {
				return nil, fmt.Errorf("unexpected value '%s' for key 'nonce'", value)
			}

			response.Nonce = nonce
		case "stale":
			stale, err := strconv.ParseBool(value)

			if err != nil {
				return nil, fmt.Errorf("unexpected value '%s' for key 'stale'", value)
			}

			response.Stale = stale
		}
	}

	return response, nil
}

func Md5Hash(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func BuildDigestHeader(method string, uri string, realm string, nonce string, username string, password string) (string, error) {
	ha1 := Md5Hash(fmt.Sprintf("%s:%s:%s", username, realm, password))
	ha2 := Md5Hash(fmt.Sprintf("%s:%s", method, uri))

	response := Md5Hash(fmt.Sprintf("%s:%s:%s", ha1, nonce, ha2))
	return fmt.Sprintf("Digest username=\"%s\", realm=\"%s\", nonce=\"%s\", uri=\"%s\", response=\"%s\"", username, realm, nonce, uri, response), nil
}
